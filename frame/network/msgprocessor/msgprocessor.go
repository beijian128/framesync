package msgprocessor

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"sync"

	"github.com/beijian128/framesync/frame/network/connection"
	"github.com/beijian128/framesync/frame/util"
	"google.golang.org/protobuf/proto"
)

type ConnectHandler func(connection.Connection)

type CloseHandler func(connection.Connection)

type BytesHandler func(conn connection.Connection, ext any, msgId uint32, bytes []byte)

type MsgHandler func(conn connection.Connection, ext any, msgId uint32, bytes []byte, msg any)

type MsgProcessor interface {
	OnDecodeExt(extData []byte) (ext any, err error)
	OnEncodeExt(ext any) (extData []byte, err error)
	OnConnect(conn connection.Connection)
	OnClose(conn connection.Connection)
	OnMessage(conn connection.Connection, ext any, msgId uint32, msgData []byte) error
}

type GetMsgHandler interface {
	GetMsgHandler(typ reflect.Type) (MsgHandler, bool)
}

// ------------消息处理-------------------------
//
// 关联id与msg结构
var (
	msgMutex   sync.Mutex
	msgID2Typ  map[uint32]reflect.Type
	msgType2ID map[reflect.Type]uint32
	msgID2Name map[uint32]string
)

func init() {
	msgID2Typ = make(map[uint32]reflect.Type)
	msgType2ID = make(map[reflect.Type]uint32)
	msgID2Name = make(map[uint32]string)
}

// RegisterMessage ...
func RegisterMessage(msg any) (uint32, error) {
	msgName := proto.MessageName(msg.(proto.Message))
	msgType := reflect.TypeOf(msg)
	return RegisterMessageNameType(string(msgName), msgType)

}

func RegisterMessageNameType(msgName string, msgType reflect.Type) (uint32, error) {
	id := util.StringHash(msgName)

	if msgType == nil || msgType.Kind() != reflect.Ptr || msgType.Elem() == nil {
		return id, errors.New("register message, message pointer required")
	}

	msgMutex.Lock()
	defer msgMutex.Unlock()

	if usedMsgType, bHave := msgID2Typ[id]; bHave {
		if usedMsgType != msgType {
			return id, fmt.Errorf("register message, message id:%v-type:%v, typein:%v is already registered", id, usedMsgType, msgType)
		}

		return id, nil
	}

	msgID2Typ[id] = msgType
	msgType2ID[msgType] = id
	msgID2Name[id] = msgName

	logrus.WithFields(logrus.Fields{
		"msgId":   id,
		"msgtype": msgType,
		"msgName": msgName,
	}).Debug("RegisterMessage")

	return id, nil
}

// MessageType ...
func MessageType(id uint32) (reflect.Type, bool) {
	msgMutex.Lock()
	defer msgMutex.Unlock()

	if msgType, bHave := msgID2Typ[id]; bHave {
		return msgType, true
	}

	return nil, false
}

// MessageID ...
func MessageID(msgType reflect.Type) (uint32, bool) {
	msgMutex.Lock()
	defer msgMutex.Unlock()

	if msgID, bHave := msgType2ID[msgType]; bHave {
		return msgID, true
	}

	return 0, false
}

func MessageName(id uint32) string {
	msgMutex.Lock()
	defer msgMutex.Unlock()

	if name, bHave := msgID2Name[id]; bHave {
		return name
	}

	return fmt.Sprintf("msgid_%d", id)
}

// OnUnmarshal ... on read gorutine
func OnUnmarshal(id uint32, data []byte) (any, error) {
	//fmt.Println("recv data ", data)
	if msgType, bHave := MessageType(id); bHave {
		msg := reflect.New(msgType.Elem()).Interface()
		//err := proto.UnmarshalMerge(data, msg.(proto.Message))
		err := json.Unmarshal(data, msg.(proto.Message))

		return msg, err
	}

	return nil, fmt.Errorf("message %v not registered", id)
}

// OnMarshal .. on write gorutine
func OnMarshal(msg any) (uint32, []byte, error) {
	//
	msgType := reflect.TypeOf(msg)

	if msgID, bHave := MessageID(msgType); bHave {
		// data
		//data, err := proto.Marshal(msg.(proto.Message))
		data, err := json.Marshal(msg.(proto.Message))

		return msgID, data, err
	}

	if msgID, e := RegisterMessage(msg); e == nil {
		// data
		//data, err := proto.Marshal(msg.(proto.Message))
		data, err := json.Marshal(msg.(proto.Message))
		return msgID, data, err
	}

	return 0, nil, fmt.Errorf("OnMarshal, message %v auto register failed", msgType)
}
