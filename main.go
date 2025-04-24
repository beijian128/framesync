package main

import (
	"flag"
	"sync"

	"github.com/beijian128/framesync/frame/appframe"
	"github.com/beijian128/framesync/frame/appframe/master"
	"github.com/beijian128/framesync/frame/framework/netcluster"
	"github.com/beijian128/framesync/frame/log"
	"github.com/beijian128/framesync/frame/util"
	"github.com/beijian128/framesync/services/gate"
	"github.com/sirupsen/logrus"
)

var (
	help          = flag.Bool("h", false, "help")
	netconfigFile = flag.String("netconfig", "netconfig.json", "netconfig file")
	appconfigFile = flag.String("config", "app.yaml", "app config")
	noMaster      = flag.Bool("noMaster", false, "ignore Master server")
	noGate        = flag.Bool("noGate", false, "ignore gate server")
)

var (
	masterName = "master"
	gateName   = "gate"

	lobbyName = "lobby"
)

func init() {
	master.DisableMasterInitGlobalLogrus = true
	appframe.DisableApplicationInitGlobalLogrus = true
}

func main() {

	flag.Parse()
	if *help {
		flag.Usage()
		return
	}

	log.InitLogrus(&log.Config{
		Name:  "chat",
		Level: 5,
		Outputs: map[string]map[string]interface{}{
			"file": map[string]interface{}{
				"path":   "./logs",
				"rotate": true,
			},
		},
	})

	netconfig, err := netcluster.ParseClusterConfigFile(*netconfigFile)
	if err != nil {
		logrus.WithError(err).Panic("netconfigFile.load", err)
		return
	}

	findOneNode := func(base string) string {
		list := []string{base, base + "1"}
		for _, target := range list {
			for nodeName, _ := range netconfig.Masters {
				if nodeName == target {
					return target
				}
			}
			for nodeName, _ := range netconfig.Slaves {
				if nodeName == target {
					return target
				}
			}
		}
		return base
	}

	wg := sync.WaitGroup{}

	// master
	if !*noMaster {
		wg.Add(1)
		util.SafeGo(func() {
			defer wg.Done()
			m, err := master.New(*netconfigFile, findOneNode(masterName))
			if err != nil {
				logrus.WithField("name", masterName).WithError(err).Panic("New master fail")
			}
			m.Run()
		})
	}

	if !*noGate {
		// gate
		wg.Add(1)
		util.SafeGo(func() {
			defer wg.Done()
			app, err := appframe.NewGateApplication(*netconfigFile, findOneNode(gateName))
			if err != nil {
				logrus.WithField("name", gateName).WithError(err).Panic("New gate app fail")
			}
			err = gate.InitGateSvr(app, *appconfigFile)
			if err != nil {
				logrus.WithField("name", gateName).WithError(err).Panic("Init gatesvr fail")
			}
			app.Run()
		})
	}

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	logrus.Info("web server start, listen on  :8080 ,http://localhost:8080/web/")
	// 	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	// 	http.ListenAndServe(":8080", nil)
	// }()

	wg.Wait()
}
