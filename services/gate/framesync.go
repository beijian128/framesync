package gate

import (
	"github.com/beijian128/framesync/frame/appframe"
	"github.com/beijian128/framesync/proto/cmsg"
	"golang.org/x/net/context"
	"math/rand"
	"time"
)

var FrameSyncInstance = NewFrameSync()

type FrameSync struct {
	curFrame map[appframe.GateSession]*cmsg.PlayerInput // 当前帧的所有玩家操作指令 。暂不考虑断线重连和回放，所以不记录历史帧
	ticker   *time.Ticker
}

func NewFrameSync() *FrameSync {
	return &FrameSync{
		curFrame: make(map[appframe.GateSession]*cmsg.PlayerInput),
	}
}

func (fs *FrameSync) OnInputCommand(session appframe.GateSession, msg *cmsg.InputCommand) {
	if fs.curFrame[session] == nil {
		fs.curFrame[session] = &cmsg.PlayerInput{}
	}
	fs.curFrame[session].Commands = append(fs.curFrame[session].Commands, msg)
}

func (fs *FrameSync) OnAddSession(ss *session) {

	ss.userid = uint64(rand.Uint64())
	ss.y = rand.Int31n(100) + 1
	ss.x = rand.Int31n(100) + 1
	//ss.color =

	SessionMgrInstance.execByEverySession(func(s *session) {
		s.SendMsg(&cmsg.PlayerEnter{
			PlayerId: ss.userid,
			X:        ss.x,
			Y:        ss.y,
			Color:    ss.color,
		})
	})
}

func (fs *FrameSync) StartSync(ctx context.Context) {
	fs.ticker = time.NewTicker(16 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-fs.ticker.C:
				AppInstance.Post(func() {
					SessionMgrInstance.execByEverySession(func(s *session) {
						for _, input := range fs.curFrame {
							s.SendMsg(input)
						}
					})
					clear(fs.curFrame)
				})
			}
		}
	}()
}

func (fs *FrameSync) RemoveSession(sess *session) {
	delete(fs.curFrame, sess.GateSession)
	SessionMgrInstance.execByEverySession(func(s *session) {
		if s.ID() == sess.ID() {
			return
		}
		s.SendMsg(&cmsg.PlayerLeave{PlayerId: sess.userid})
	})
}
