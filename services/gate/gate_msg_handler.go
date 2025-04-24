package gate

import (
	"github.com/beijian128/framesync/frame/appframe"
	"github.com/beijian128/framesync/frame/framework/netframe"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func initGateMsgHandler(app *appframe.GateApplication) {

}

func initGateMsgRoute(app *appframe.GateApplication) {
	logrus.Info("------------------initGateMsgRoute ---------------------")

	//lobby
	//lobby := app.GetService(services.ServiceType_Lobby)
	//RouteSessionRowMsg(app, (*cmsg.CReqLogin)(nil), lobby)
	//RouteSessionRowMsg(app, (*cmsg.CReqSendChatMessage)(nil), lobby)

	appframe.ListenGateSessionMsgSugar(app, FrameSyncInstance.OnInputCommand)

}

func GetAnyRawMessageRouter(f func(s *session, msgId uint32, data []byte, extParam int64)) appframe.GateSessionRawMsgRouter {
	return func(sid uint64, msgId uint32, data []byte, extParam int64) {
		s, ok := SessionMgrInstance.getSession(sid)
		if ok {
			f(s, msgId, data, extParam)
		}
	}
}

func RouteSessionRowMsg(app *appframe.GateApplication, msg proto.Message, service appframe.Service) {
	app.RouteSessionRawMsg(msg, GetAnyRawMessageRouter(func(s *session, mssgid uint32, data []byte, extParam int64) {
		service.ForwardRawMsgFromSession(mssgid, data, netframe.Server_Extend{
			SessionId: s.ID(),
			UserId:    s.userid,
			ExtParam:  extParam,
		})
	}))
}
