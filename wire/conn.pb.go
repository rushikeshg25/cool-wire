// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: conn.proto

package wire

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type Query struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Params        []*structpb.Value      `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Query) Reset() {
	*x = Query{}
	mi := &file_conn_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{0}
}

func (x *Query) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Query) GetParams() []*structpb.Value {
	if x != nil {
		return x.Params
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Response      string                 `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_conn_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_conn_proto protoreflect.FileDescriptor

const file_conn_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"conn.proto\x12\bcoolwire\x1a\x1cgoogle/protobuf/struct.proto\"M\n" +
	"\x05Query\x12\x14\n" +
	"\x05query\x18\x01 \x01(\tR\x05query\x12.\n" +
	"\x06params\x18\x02 \x03(\v2\x16.google.protobuf.ValueR\x06params\"&\n" +
	"\bResponse\x12\x1a\n" +
	"\bresponse\x18\x01 \x01(\tR\bresponse2?\n" +
	"\vWireService\x120\n" +
	"\tSendQuery\x12\x0f.coolwire.Query\x1a\x12.coolwire.ResponseB\bZ\x06./wireb\x06proto3"

var (
	file_conn_proto_rawDescOnce sync.Once
	file_conn_proto_rawDescData []byte
)

func file_conn_proto_rawDescGZIP() []byte {
	file_conn_proto_rawDescOnce.Do(func() {
		file_conn_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_conn_proto_rawDesc), len(file_conn_proto_rawDesc)))
	})
	return file_conn_proto_rawDescData
}

var file_conn_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_conn_proto_goTypes = []any{
	(*Query)(nil),          // 0: coolwire.Query
	(*Response)(nil),       // 1: coolwire.Response
	(*structpb.Value)(nil), // 2: google.protobuf.Value
}
var file_conn_proto_depIdxs = []int32{
	2, // 0: coolwire.Query.params:type_name -> google.protobuf.Value
	0, // 1: coolwire.WireService.SendQuery:input_type -> coolwire.Query
	1, // 2: coolwire.WireService.SendQuery:output_type -> coolwire.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_conn_proto_init() }
func file_conn_proto_init() {
	if File_conn_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_conn_proto_rawDesc), len(file_conn_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_conn_proto_goTypes,
		DependencyIndexes: file_conn_proto_depIdxs,
		MessageInfos:      file_conn_proto_msgTypes,
	}.Build()
	File_conn_proto = out.File
	file_conn_proto_goTypes = nil
	file_conn_proto_depIdxs = nil
}
