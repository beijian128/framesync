package gate

import (
	_ "net/http/pprof"

	"github.com/beijian128/framesync/config"
	"github.com/beijian128/framesync/frame/appframe"
	appframeslb "github.com/beijian128/framesync/frame/appframe/slb"
	"github.com/beijian128/framesync/frame/framework/netcluster"
	"github.com/beijian128/framesync/services"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

var SessionMgrInstance *sessionManager

var AppInstance *appframe.GateApplication
var AppCfg *config.AppConfig

// InitGateSvr 初始化 gatesvr.
func InitGateSvr(app *appframe.GateApplication, cfgFile string) error {

	cfg, err := config.LoadConfig(cfgFile)
	if err != nil {
		return err
	}
	AppCfg = cfg
	AppInstance = app

	app.RegisterService(services.ServiceType_Lobby, appframeslb.WithLoadBalanceSingleton(app, services.ServiceType_Lobby))

	SessionMgrInstance = initSessionManager(app)

	initGateMsgRoute(app)
	initGateMsgHandler(app)

	app.OnExitHandler(Close)

	ctx, cancel := context.WithCancel(context.Background())
	app.OnRunHandler(func() {
		FrameSyncInstance.StartSync(ctx)
	})
	app.OnFiniHandler(func() {
		cancel()
	})
	return nil
}

func Close() {

}

func onLobbyServerDisconnect(svrID uint32, event netcluster.SvrEvent) {
	switch event {
	case netcluster.SvrEventQuit, netcluster.SvrEventDisconnect:
		logrus.WithFields(logrus.Fields{
			"svrID": svrID,
			"event": event,
		}).Error("lobby server disconnect")

		SessionMgrInstance.execByEverySession(func(s *session) {
			s.Close()
		})
	}
}
