// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: connect.proto

package network

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConnectRequestType int32

const (
	ConnectRequest_tcp ConnectRequestType = 0
	ConnectRequest_udp ConnectRequestType = 1
)

// Enum value maps for ConnectRequestType.
var (
	ConnectRequestType_name = map[int32]string{
		0: "tcp",
		1: "udp",
	}
	ConnectRequestType_value = map[string]int32{
		"tcp": 0,
		"udp": 1,
	}
)

func (x ConnectRequestType) Enum() *ConnectRequestType {
	p := new(ConnectRequestType)
	*p = x
	return p
}

func (x ConnectRequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectRequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_connect_proto_enumTypes[0].Descriptor()
}

func (ConnectRequestType) Type() protoreflect.EnumType {
	return &file_connect_proto_enumTypes[0]
}

func (x ConnectRequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectRequestType.Descriptor instead.
func (ConnectRequestType) EnumDescriptor() ([]byte, []int) {
	return file_connect_proto_rawDescGZIP(), []int{0, 0}
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XType ConnectRequestType `protobuf:"varint,2,opt,name=_type,json=Type,proto3,enum=network.ConnectRequestType" json:"_type,omitempty"`
	Addr  string             `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	Port  uint32             `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_connect_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConnectRequest) GetXType() ConnectRequestType {
	if x != nil {
		return x.XType
	}
	return ConnectRequest_tcp
}

func (x *ConnectRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *ConnectRequest) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_connect_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ConnectResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_connect_proto protoreflect.FileDescriptor

var file_connect_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x9b, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x32, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x18, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x74, 0x63, 0x70, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x75, 0x64, 0x70, 0x10, 0x01, 0x22, 0x34, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x3b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_connect_proto_rawDescOnce sync.Once
	file_connect_proto_rawDescData = file_connect_proto_rawDesc
)

func file_connect_proto_rawDescGZIP() []byte {
	file_connect_proto_rawDescOnce.Do(func() {
		file_connect_proto_rawDescData = protoimpl.X.CompressGZIP(file_connect_proto_rawDescData)
	})
	return file_connect_proto_rawDescData
}

var file_connect_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_connect_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_connect_proto_goTypes = []interface{}{
	(ConnectRequestType)(0), // 0: network.connect_request.type
	(*ConnectRequest)(nil),  // 1: network.connect_request
	(*ConnectResponse)(nil), // 2: network.connect_response
}
var file_connect_proto_depIdxs = []int32{
	0, // 0: network.connect_request._type:type_name -> network.connect_request.type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_connect_proto_init() }
func file_connect_proto_init() {
	if File_connect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_connect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_connect_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connect_proto_goTypes,
		DependencyIndexes: file_connect_proto_depIdxs,
		EnumInfos:         file_connect_proto_enumTypes,
		MessageInfos:      file_connect_proto_msgTypes,
	}.Build()
	File_connect_proto = out.File
	file_connect_proto_rawDesc = nil
	file_connect_proto_goTypes = nil
	file_connect_proto_depIdxs = nil
}
