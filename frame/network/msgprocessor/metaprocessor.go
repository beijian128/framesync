package msgprocessor

import (
	"fmt"
	"github.com/beijian128/framesync/frame/framework/worker"
	"reflect"

	"google.golang.org/protobuf/proto"

	"github.com/beijian128/framesync/frame/network/connection"
)

// MetaProcessor ...
type MetaProcessor struct {
	MsgProcessor
	ConnectHandler ConnectHandler
	CloseHandler   CloseHandler
	MsgHandlers    GetMsgHandler
	BytesHandler   BytesHandler

	//MonitorHandler   func(conn connection.Connection,ext any,msgId uint32,msgData []byte)

	extType    reflect.Type
	callbackIO worker.IWorker
}

// NewMetaProcessor ...
func NewMetaProcessor(ext any, worker worker.IWorker) *MetaProcessor {
	if worker == nil {
		panic("init MetaProcessor ioservice is nil")
	}
	extType := reflect.TypeOf(ext)
	if extType != nil && (extType.Kind() != reflect.Ptr || extType.Elem() == nil) {
		panic("init MetaProcessor message ext type err, required pointer")
	}

	p := new(MetaProcessor)
	p.extType = extType
	p.callbackIO = worker
	return p
}

// OnDecodeExt decode扩展数据
func (p *MetaProcessor) OnDecodeExt(extData []byte) (ext any, err error) {
	if extData != nil && p.extType != nil {
		ext = reflect.New(p.extType.Elem()).Interface()
		err = proto.Unmarshal(extData, ext.(proto.Message))
		return ext, err
	}
	return extData, err
}

// OnEncodeExt encode扩展数据
func (p *MetaProcessor) OnEncodeExt(ext any) (extData []byte, err error) {
	if ext == nil {
		return nil, nil
	}
	if ext, ok := ext.([]byte); ok {
		return ext, nil
	}
	if ext, ok := ext.(proto.Message); ok {
		return proto.Marshal(ext)
	}
	return nil, fmt.Errorf("OnEncodeExt ext %s not a proto.Message type", reflect.TypeOf(ext))
}

// OnConnect ...
func (p *MetaProcessor) OnConnect(conn connection.Connection) {
	if p.ConnectHandler != nil {
		p.callbackIO.Post(func() {
			p.ConnectHandler(conn)
		})
	}
}

// OnClose ...
func (p *MetaProcessor) OnClose(conn connection.Connection) {
	if p.CloseHandler != nil {
		p.callbackIO.Post(func() {
			p.CloseHandler(conn)
		})
	}
}

// OnMessage 消息回调
func (p *MetaProcessor) OnMessage(conn connection.Connection, ext any, msgId uint32, msgData []byte) error {
	handler, ok := p.findMsgHandler(msgId)
	if !ok {
		if p.BytesHandler != nil {
			p.callbackIO.Post(func() {
				p.BytesHandler(conn, ext, msgId, msgData)
			})
		}
		return nil
	}

	msg, err := OnUnmarshal(msgId, msgData)
	if err != nil {
		return err
	}

	p.callbackIO.Post(func() {
		handler(conn, ext, msgId, msgData, msg)
	})

	return nil
}

func (p *MetaProcessor) findMsgHandler(msgId uint32) (MsgHandler, bool) {
	if p.MsgHandlers == nil {
		return nil, false
	}

	typ, ok := MessageType(msgId)
	if !ok {
		return nil, false
	}

	return p.MsgHandlers.GetMsgHandler(typ)
}
