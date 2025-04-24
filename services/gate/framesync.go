package gate

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/beijian128/framesync/frame/appframe"
	"github.com/beijian128/framesync/proto/cmsg"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

var FrameSyncInstance = NewFrameSync()

type FrameSync struct {
	curFrame map[appframe.GateSession]*cmsg.PlayerInput // 当前帧的所有玩家操作指令 。暂不考虑断线重连和回放，所以不记录历史帧
	ticker   *time.Ticker
	userCnt  uint64
}

func NewFrameSync() *FrameSync {
	return &FrameSync{
		curFrame: make(map[appframe.GateSession]*cmsg.PlayerInput),
	}
}

func (fs *FrameSync) OnInputCommand(session appframe.GateSession, msg *cmsg.InputCommand) {
	s, _ := SessionMgrInstance.getSession(session.ID())
	if fs.curFrame[session] == nil {
		fs.curFrame[session] = &cmsg.PlayerInput{
			PlayerId: s.userid,
			Commands: []*cmsg.InputCommand{},
		}
	}
	fs.curFrame[session].Commands = append(fs.curFrame[session].Commands, msg)
	speed := int32(3)
	if msg.Down {
		s.y += int32(speed)
	}
	if msg.Up {
		s.y -= speed
	}
	if msg.Left {
		s.x -= speed
	}
	if msg.Right {
		s.x += speed
	}
}

func (fs *FrameSync) OnAddSession(ss *session) {
	fs.userCnt++
	ss.userid = fs.userCnt
	ss.y = rand.Int31n(100) + 1
	ss.x = rand.Int31n(100) + 1
	ss.color = generateRandomColor()

	SessionMgrInstance.execByEverySession(func(s *session) {
		SessionMgrInstance.execByEverySession(func(s2 *session) {
			s.SendMsg(&cmsg.PlayerEnter{
				PlayerId: s2.userid,
				X:        s2.x,
				Y:        s2.y,
				Color:    s2.color,
			})
		})

	})
}

func (fs *FrameSync) StartSync(ctx context.Context) {
	fs.ticker = time.NewTicker(16 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fs.ticker.Stop()
				logrus.Debug("FrameSync exit")
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

func generateRandomColor() string {
	// 初始化随机数种子（Go需要显式设置种子）
	rand.Seed(time.Now().UnixNano())

	// 生成 0 到 16777215 之间的随机数（即 0x000000 到 0xFFFFFF）
	colorValue := rand.Intn(16777216) // 注意：Intn 的上限是 16777216（不包含）

	// 格式化为 6 位十六进制字符串，不足补零
	return fmt.Sprintf("#%06x", colorValue)
}
