package gate

import (
	_ "net/http/pprof"

	"github.com/beijian128/framesync/frame/appframe"
	"golang.org/x/net/context"
)

var SessionMgrInstance *sessionManager

var AppInstance *appframe.GateApplication

func Init(app *appframe.GateApplication) error {
	AppInstance = app

	SessionMgrInstance = initSessionManager(app)

	appframe.ListenGateSessionMsgSugar(app, FrameSyncInstance.OnInputCommand)

	ctx, cancel := context.WithCancel(context.Background())
	app.OnRunHandler(func() {
		FrameSyncInstance.StartSync(ctx)
	})
	app.OnFiniHandler(func() {
		cancel()
	})
	return nil
}
