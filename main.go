package main

import (
	"flag"
	"github.com/beijian128/framesync/services/gate"
	"net/http"
	"sync"
	"time"

	"github.com/beijian128/framesync/frame/appframe"
	"github.com/beijian128/framesync/frame/appframe/master"
	"github.com/beijian128/framesync/frame/log"
	"github.com/beijian128/framesync/frame/util"
	"github.com/sirupsen/logrus"
)

var (
	help        = flag.Bool("h", false, "help")
	clusterFile = flag.String("集群配置", "cluster_file.json", "集群配置/服务发现")
)

var (
	masterName = "master"
	gateName   = "gate"
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

	_, err := log.InitLogrus(&log.Config{
		Name:  "chat",
		Level: 5,
		Outputs: map[string]map[string]interface{}{
			"file": {
				"path":   "./logs",
				"rotate": true,
			},
		},
	})
	if err != nil {
		return
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	util.SafeGo(func() {
		defer wg.Done()
		m, err := master.New(*clusterFile, masterName)
		if err != nil {
			logrus.WithField("name", masterName).WithError(err).Panic("New master fail")
		}
		m.Run()
	})

	wg.Add(1)
	util.SafeGo(func() {
		defer wg.Done()
		app, err := appframe.NewGateApplication(*clusterFile, gateName)
		if err != nil {
			logrus.WithField("name", gateName).WithError(err).Panic("New gate app fail")
		}
		err = gate.Init(app)
		if err != nil {
			logrus.WithField("name", gateName).WithError(err).Panic("Init gatesvr fail")
		}
		app.Run()
	})

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.AfterFunc(time.Second*1, func() {
			logrus.Info("浏览器打开 http://localhost:8080/web/")
		})
		http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
		http.ListenAndServe(":8080", nil)
	}()

	wg.Wait()
}
