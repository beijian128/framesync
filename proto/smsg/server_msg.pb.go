// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: smsg/server_msg.proto

package smsg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// gate->lobby
type NoticeSessionClosed struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Userid        uint64                 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Session       uint64                 `protobuf:"varint,2,opt,name=session,proto3" json:"session,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NoticeSessionClosed) Reset() {
	*x = NoticeSessionClosed{}
	mi := &file_smsg_server_msg_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoticeSessionClosed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeSessionClosed) ProtoMessage() {}

func (x *NoticeSessionClosed) ProtoReflect() protoreflect.Message {
	mi := &file_smsg_server_msg_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeSessionClosed.ProtoReflect.Descriptor instead.
func (*NoticeSessionClosed) Descriptor() ([]byte, []int) {
	return file_smsg_server_msg_proto_rawDescGZIP(), []int{0}
}

func (x *NoticeSessionClosed) GetUserid() uint64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *NoticeSessionClosed) GetSession() uint64 {
	if x != nil {
		return x.Session
	}
	return 0
}

// ErrCode 通用错误定义
type ErrCode struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 用于响应消息(框架使用的字段, 使用者不用手动设置).
	Seqid int64 `protobuf:"varint,1,opt,name=seqid,proto3" json:"seqid,omitempty"`
	// 错误码.
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	// 错误详细消息.
	Msg           string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ErrCode) Reset() {
	*x = ErrCode{}
	mi := &file_smsg_server_msg_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrCode) ProtoMessage() {}

func (x *ErrCode) ProtoReflect() protoreflect.Message {
	mi := &file_smsg_server_msg_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrCode.ProtoReflect.Descriptor instead.
func (*ErrCode) Descriptor() ([]byte, []int) {
	return file_smsg_server_msg_proto_rawDescGZIP(), []int{1}
}

func (x *ErrCode) GetSeqid() int64 {
	if x != nil {
		return x.Seqid
	}
	return 0
}

func (x *ErrCode) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrCode) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_smsg_server_msg_proto protoreflect.FileDescriptor

const file_smsg_server_msg_proto_rawDesc = "" +
	"\n" +
	"\x15smsg/server_msg.proto\x12\x04smsg\"G\n" +
	"\x13NoticeSessionClosed\x12\x16\n" +
	"\x06userid\x18\x01 \x01(\x04R\x06userid\x12\x18\n" +
	"\asession\x18\x02 \x01(\x04R\asession\"E\n" +
	"\aErrCode\x12\x14\n" +
	"\x05seqid\x18\x01 \x01(\x03R\x05seqid\x12\x12\n" +
	"\x04code\x18\x02 \x01(\x05R\x04code\x12\x10\n" +
	"\x03msg\x18\x03 \x01(\tR\x03msgB\aZ\x05/smsgb\x06proto3"

var (
	file_smsg_server_msg_proto_rawDescOnce sync.Once
	file_smsg_server_msg_proto_rawDescData []byte
)

func file_smsg_server_msg_proto_rawDescGZIP() []byte {
	file_smsg_server_msg_proto_rawDescOnce.Do(func() {
		file_smsg_server_msg_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_smsg_server_msg_proto_rawDesc), len(file_smsg_server_msg_proto_rawDesc)))
	})
	return file_smsg_server_msg_proto_rawDescData
}

var file_smsg_server_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_smsg_server_msg_proto_goTypes = []any{
	(*NoticeSessionClosed)(nil), // 0: smsg.NoticeSessionClosed
	(*ErrCode)(nil),             // 1: smsg.ErrCode
}
var file_smsg_server_msg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_smsg_server_msg_proto_init() }
func file_smsg_server_msg_proto_init() {
	if File_smsg_server_msg_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_smsg_server_msg_proto_rawDesc), len(file_smsg_server_msg_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_smsg_server_msg_proto_goTypes,
		DependencyIndexes: file_smsg_server_msg_proto_depIdxs,
		MessageInfos:      file_smsg_server_msg_proto_msgTypes,
	}.Build()
	File_smsg_server_msg_proto = out.File
	file_smsg_server_msg_proto_goTypes = nil
	file_smsg_server_msg_proto_depIdxs = nil
}
