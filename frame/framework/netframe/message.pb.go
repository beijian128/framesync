// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: netframe/message.proto

package netframe

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

// extend message
type Server_Extend struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	ServerId      uint32                    `protobuf:"varint,1,opt,name=ServerId,proto3" json:"ServerId,omitempty"`
	SessionId     uint64                    `protobuf:"varint,2,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	SeqId         int64                     `protobuf:"varint,3,opt,name=SeqId,proto3" json:"SeqId,omitempty"` // s/s 内部服务间RPC 序列号
	UserId        uint64                    `protobuf:"varint,4,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ExtParam      int64                     `protobuf:"varint,5,opt,name=ExtParam,proto3" json:"ExtParam,omitempty"` //扩展参数
	RouterMeta    *Server_Extend_RouterMeta `protobuf:"bytes,6,opt,name=routerMeta,proto3" json:"routerMeta,omitempty"`
	Account       string                    `protobuf:"bytes,7,opt,name=account,proto3" json:"account,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_Extend) Reset() {
	*x = Server_Extend{}
	mi := &file_netframe_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_Extend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_Extend) ProtoMessage() {}

func (x *Server_Extend) ProtoReflect() protoreflect.Message {
	mi := &file_netframe_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_Extend.ProtoReflect.Descriptor instead.
func (*Server_Extend) Descriptor() ([]byte, []int) {
	return file_netframe_message_proto_rawDescGZIP(), []int{0}
}

func (x *Server_Extend) GetServerId() uint32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *Server_Extend) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *Server_Extend) GetSeqId() int64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *Server_Extend) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Server_Extend) GetExtParam() int64 {
	if x != nil {
		return x.ExtParam
	}
	return 0
}

func (x *Server_Extend) GetRouterMeta() *Server_Extend_RouterMeta {
	if x != nil {
		return x.RouterMeta
	}
	return nil
}

func (x *Server_Extend) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// heart
type Server_ReqHeartBeat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TimeStamp     int64                  `protobuf:"varint,1,opt,name=TimeStamp,proto3" json:"TimeStamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_ReqHeartBeat) Reset() {
	*x = Server_ReqHeartBeat{}
	mi := &file_netframe_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_ReqHeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_ReqHeartBeat) ProtoMessage() {}

func (x *Server_ReqHeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_netframe_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_ReqHeartBeat.ProtoReflect.Descriptor instead.
func (*Server_ReqHeartBeat) Descriptor() ([]byte, []int) {
	return file_netframe_message_proto_rawDescGZIP(), []int{1}
}

func (x *Server_ReqHeartBeat) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

type Server_RespHeartBeat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TimeStamp     int64                  `protobuf:"varint,1,opt,name=TimeStamp,proto3" json:"TimeStamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_RespHeartBeat) Reset() {
	*x = Server_RespHeartBeat{}
	mi := &file_netframe_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_RespHeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_RespHeartBeat) ProtoMessage() {}

func (x *Server_RespHeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_netframe_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_RespHeartBeat.ProtoReflect.Descriptor instead.
func (*Server_RespHeartBeat) Descriptor() ([]byte, []int) {
	return file_netframe_message_proto_rawDescGZIP(), []int{2}
}

func (x *Server_RespHeartBeat) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

// client->server 注册
type Server_ReqRegister struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ClientType      uint32                 `protobuf:"varint,1,opt,name=ClientType,proto3" json:"ClientType,omitempty"`
	ClientId        uint32                 `protobuf:"varint,2,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	ClientStartTime int64                  `protobuf:"varint,3,opt,name=ClientStartTime,proto3" json:"ClientStartTime,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Server_ReqRegister) Reset() {
	*x = Server_ReqRegister{}
	mi := &file_netframe_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_ReqRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_ReqRegister) ProtoMessage() {}

func (x *Server_ReqRegister) ProtoReflect() protoreflect.Message {
	mi := &file_netframe_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_ReqRegister.ProtoReflect.Descriptor instead.
func (*Server_ReqRegister) Descriptor() ([]byte, []int) {
	return file_netframe_message_proto_rawDescGZIP(), []int{3}
}

func (x *Server_ReqRegister) GetClientType() uint32 {
	if x != nil {
		return x.ClientType
	}
	return 0
}

func (x *Server_ReqRegister) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *Server_ReqRegister) GetClientStartTime() int64 {
	if x != nil {
		return x.ClientStartTime
	}
	return 0
}

// server->client
type Server_RespRegister struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ServerType      uint32                 `protobuf:"varint,1,opt,name=ServerType,proto3" json:"ServerType,omitempty"`
	ServerId        uint32                 `protobuf:"varint,2,opt,name=ServerId,proto3" json:"ServerId,omitempty"`
	ServerStartTime int64                  `protobuf:"varint,3,opt,name=ServerStartTime,proto3" json:"ServerStartTime,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Server_RespRegister) Reset() {
	*x = Server_RespRegister{}
	mi := &file_netframe_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_RespRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_RespRegister) ProtoMessage() {}

func (x *Server_RespRegister) ProtoReflect() protoreflect.Message {
	mi := &file_netframe_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_RespRegister.ProtoReflect.Descriptor instead.
func (*Server_RespRegister) Descriptor() ([]byte, []int) {
	return file_netframe_message_proto_rawDescGZIP(), []int{4}
}

func (x *Server_RespRegister) GetServerType() uint32 {
	if x != nil {
		return x.ServerType
	}
	return 0
}

func (x *Server_RespRegister) GetServerId() uint32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *Server_RespRegister) GetServerStartTime() int64 {
	if x != nil {
		return x.ServerStartTime
	}
	return 0
}

// client->server 取消注册
type Server_ReportUnRegister struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ServerStartTime int64                  `protobuf:"varint,1,opt,name=ServerStartTime,proto3" json:"ServerStartTime,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Server_ReportUnRegister) Reset() {
	*x = Server_ReportUnRegister{}
	mi := &file_netframe_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_ReportUnRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_ReportUnRegister) ProtoMessage() {}

func (x *Server_ReportUnRegister) ProtoReflect() protoreflect.Message {
	mi := &file_netframe_message_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_ReportUnRegister.ProtoReflect.Descriptor instead.
func (*Server_ReportUnRegister) Descriptor() ([]byte, []int) {
	return file_netframe_message_proto_rawDescGZIP(), []int{5}
}

func (x *Server_ReportUnRegister) GetServerStartTime() int64 {
	if x != nil {
		return x.ServerStartTime
	}
	return 0
}

type Client_ReqHeartBeat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TimeStamp     int64                  `protobuf:"varint,1,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Client_ReqHeartBeat) Reset() {
	*x = Client_ReqHeartBeat{}
	mi := &file_netframe_message_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Client_ReqHeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client_ReqHeartBeat) ProtoMessage() {}

func (x *Client_ReqHeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_netframe_message_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client_ReqHeartBeat.ProtoReflect.Descriptor instead.
func (*Client_ReqHeartBeat) Descriptor() ([]byte, []int) {
	return file_netframe_message_proto_rawDescGZIP(), []int{6}
}

func (x *Client_ReqHeartBeat) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

type Client_RespHeartBeat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TimeStamp     int64                  `protobuf:"varint,1,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Client_RespHeartBeat) Reset() {
	*x = Client_RespHeartBeat{}
	mi := &file_netframe_message_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Client_RespHeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client_RespHeartBeat) ProtoMessage() {}

func (x *Client_RespHeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_netframe_message_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client_RespHeartBeat.ProtoReflect.Descriptor instead.
func (*Client_RespHeartBeat) Descriptor() ([]byte, []int) {
	return file_netframe_message_proto_rawDescGZIP(), []int{7}
}

func (x *Client_RespHeartBeat) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

type Server_Extend_RouterMeta struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReqFrom       uint32                 `protobuf:"varint,1,opt,name=reqFrom,proto3" json:"reqFrom,omitempty"` // req 的发起者 svrID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_Extend_RouterMeta) Reset() {
	*x = Server_Extend_RouterMeta{}
	mi := &file_netframe_message_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_Extend_RouterMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_Extend_RouterMeta) ProtoMessage() {}

func (x *Server_Extend_RouterMeta) ProtoReflect() protoreflect.Message {
	mi := &file_netframe_message_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_Extend_RouterMeta.ProtoReflect.Descriptor instead.
func (*Server_Extend_RouterMeta) Descriptor() ([]byte, []int) {
	return file_netframe_message_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Server_Extend_RouterMeta) GetReqFrom() uint32 {
	if x != nil {
		return x.ReqFrom
	}
	return 0
}

var File_netframe_message_proto protoreflect.FileDescriptor

const file_netframe_message_proto_rawDesc = "" +
	"\n" +
	"\x16netframe/message.proto\x12\bnetframe\"\x99\x02\n" +
	"\rServer_Extend\x12\x1a\n" +
	"\bServerId\x18\x01 \x01(\rR\bServerId\x12\x1c\n" +
	"\tSessionId\x18\x02 \x01(\x04R\tSessionId\x12\x14\n" +
	"\x05SeqId\x18\x03 \x01(\x03R\x05SeqId\x12\x16\n" +
	"\x06UserId\x18\x04 \x01(\x04R\x06UserId\x12\x1a\n" +
	"\bExtParam\x18\x05 \x01(\x03R\bExtParam\x12B\n" +
	"\n" +
	"routerMeta\x18\x06 \x01(\v2\".netframe.Server_Extend.RouterMetaR\n" +
	"routerMeta\x12\x18\n" +
	"\aaccount\x18\a \x01(\tR\aaccount\x1a&\n" +
	"\n" +
	"RouterMeta\x12\x18\n" +
	"\areqFrom\x18\x01 \x01(\rR\areqFrom\"3\n" +
	"\x13Server_ReqHeartBeat\x12\x1c\n" +
	"\tTimeStamp\x18\x01 \x01(\x03R\tTimeStamp\"4\n" +
	"\x14Server_RespHeartBeat\x12\x1c\n" +
	"\tTimeStamp\x18\x01 \x01(\x03R\tTimeStamp\"z\n" +
	"\x12Server_ReqRegister\x12\x1e\n" +
	"\n" +
	"ClientType\x18\x01 \x01(\rR\n" +
	"ClientType\x12\x1a\n" +
	"\bClientId\x18\x02 \x01(\rR\bClientId\x12(\n" +
	"\x0fClientStartTime\x18\x03 \x01(\x03R\x0fClientStartTime\"{\n" +
	"\x13Server_RespRegister\x12\x1e\n" +
	"\n" +
	"ServerType\x18\x01 \x01(\rR\n" +
	"ServerType\x12\x1a\n" +
	"\bServerId\x18\x02 \x01(\rR\bServerId\x12(\n" +
	"\x0fServerStartTime\x18\x03 \x01(\x03R\x0fServerStartTime\"C\n" +
	"\x17Server_ReportUnRegister\x12(\n" +
	"\x0fServerStartTime\x18\x01 \x01(\x03R\x0fServerStartTime\"3\n" +
	"\x13Client_ReqHeartBeat\x12\x1c\n" +
	"\ttimeStamp\x18\x01 \x01(\x03R\ttimeStamp\"4\n" +
	"\x14Client_RespHeartBeat\x12\x1c\n" +
	"\ttimeStamp\x18\x01 \x01(\x03R\ttimeStampB\vZ\t/netframeb\x06proto3"

var (
	file_netframe_message_proto_rawDescOnce sync.Once
	file_netframe_message_proto_rawDescData []byte
)

func file_netframe_message_proto_rawDescGZIP() []byte {
	file_netframe_message_proto_rawDescOnce.Do(func() {
		file_netframe_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_netframe_message_proto_rawDesc), len(file_netframe_message_proto_rawDesc)))
	})
	return file_netframe_message_proto_rawDescData
}

var file_netframe_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_netframe_message_proto_goTypes = []any{
	(*Server_Extend)(nil),            // 0: netframe.Server_Extend
	(*Server_ReqHeartBeat)(nil),      // 1: netframe.Server_ReqHeartBeat
	(*Server_RespHeartBeat)(nil),     // 2: netframe.Server_RespHeartBeat
	(*Server_ReqRegister)(nil),       // 3: netframe.Server_ReqRegister
	(*Server_RespRegister)(nil),      // 4: netframe.Server_RespRegister
	(*Server_ReportUnRegister)(nil),  // 5: netframe.Server_ReportUnRegister
	(*Client_ReqHeartBeat)(nil),      // 6: netframe.Client_ReqHeartBeat
	(*Client_RespHeartBeat)(nil),     // 7: netframe.Client_RespHeartBeat
	(*Server_Extend_RouterMeta)(nil), // 8: netframe.Server_Extend.RouterMeta
}
var file_netframe_message_proto_depIdxs = []int32{
	8, // 0: netframe.Server_Extend.routerMeta:type_name -> netframe.Server_Extend.RouterMeta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_netframe_message_proto_init() }
func file_netframe_message_proto_init() {
	if File_netframe_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_netframe_message_proto_rawDesc), len(file_netframe_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_netframe_message_proto_goTypes,
		DependencyIndexes: file_netframe_message_proto_depIdxs,
		MessageInfos:      file_netframe_message_proto_msgTypes,
	}.Build()
	File_netframe_message_proto = out.File
	file_netframe_message_proto_goTypes = nil
	file_netframe_message_proto_depIdxs = nil
}
