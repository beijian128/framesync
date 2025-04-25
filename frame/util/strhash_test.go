package util

import (
	_ "github.com/beijian128/framesync/proto/cmsg"
	_ "github.com/beijian128/framesync/proto/smsg"
	"google.golang.org/protobuf/reflect/protoregistry"
	"testing"
)

func TestStringHash(t *testing.T) {

	file := "cmsg/client_msg.proto"
	fd, err := protoregistry.GlobalFiles.FindFileByPath(file)
	if err != nil {
		return
	}

	mds := fd.Messages()
	for i := mds.Len() - 1; i >= 0; i-- {
		x := mds.Get(i)
		fullName := x.FullName()
		_, err := protoregistry.GlobalTypes.FindMessageByName(fullName)
		if err != nil {
			return
		}
		msgID := StringHash(string(fullName))
		t.Log(fullName, msgID)
	}

}

func TestStringHashProtoName3(t *testing.T) {

	file := "smsg/gate_msg.proto"
	fd, err := protoregistry.GlobalFiles.FindFileByPath(file)
	if err != nil {
		return
	}

	mds := fd.Messages()
	for i := mds.Len() - 1; i >= 0; i-- {
		x := mds.Get(i)
		fullName := x.FullName()
		_, err := protoregistry.GlobalTypes.FindMessageByName(fullName)
		if err != nil {
			return
		}
		msgID := StringHash(string(fullName))
		if msgID == 3161802341 {
			t.Log(fullName, msgID)
		}

	}

}
