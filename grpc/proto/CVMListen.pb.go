// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CVMListen.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}

var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}

func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24800a83a703ac1f, []int{0}
}

type InfoMessage struct {
	OperationType        int32    `protobuf:"varint,1,opt,name=OperationType,proto3" json:"OperationType,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoMessage) Reset()         { *m = InfoMessage{} }
func (m *InfoMessage) String() string { return proto.CompactTextString(m) }
func (*InfoMessage) ProtoMessage()    {}
func (*InfoMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_24800a83a703ac1f, []int{0}
}

func (m *InfoMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoMessage.Unmarshal(m, b)
}
func (m *InfoMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoMessage.Marshal(b, m, deterministic)
}
func (m *InfoMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoMessage.Merge(m, src)
}
func (m *InfoMessage) XXX_Size() int {
	return xxx_messageInfo_InfoMessage.Size(m)
}
func (m *InfoMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InfoMessage proto.InternalMessageInfo

func (m *InfoMessage) GetOperationType() int32 {
	if m != nil {
		return m.OperationType
	}
	return 0
}

func (m *InfoMessage) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Bytes struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bytes) Reset()         { *m = Bytes{} }
func (m *Bytes) String() string { return proto.CompactTextString(m) }
func (*Bytes) ProtoMessage()    {}
func (*Bytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_24800a83a703ac1f, []int{1}
}

func (m *Bytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bytes.Unmarshal(m, b)
}
func (m *Bytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bytes.Marshal(b, m, deterministic)
}
func (m *Bytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bytes.Merge(m, src)
}
func (m *Bytes) XXX_Size() int {
	return xxx_messageInfo_Bytes.Size(m)
}
func (m *Bytes) XXX_DiscardUnknown() {
	xxx_messageInfo_Bytes.DiscardUnknown(m)
}

var xxx_messageInfo_Bytes proto.InternalMessageInfo

func (m *Bytes) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Result struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_24800a83a703ac1f, []int{2}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

type UploadStatus struct {
	Message              string           `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code                 UploadStatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=proto.UploadStatusCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadStatus) Reset()         { *m = UploadStatus{} }
func (m *UploadStatus) String() string { return proto.CompactTextString(m) }
func (*UploadStatus) ProtoMessage()    {}
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_24800a83a703ac1f, []int{3}
}

func (m *UploadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadStatus.Unmarshal(m, b)
}
func (m *UploadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadStatus.Marshal(b, m, deterministic)
}
func (m *UploadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStatus.Merge(m, src)
}
func (m *UploadStatus) XXX_Size() int {
	return xxx_messageInfo_UploadStatus.Size(m)
}
func (m *UploadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStatus proto.InternalMessageInfo

func (m *UploadStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadStatus) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

func init() {
	proto.RegisterEnum("proto.UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
	proto.RegisterType((*InfoMessage)(nil), "proto.InfoMessage")
	proto.RegisterType((*Bytes)(nil), "proto.Bytes")
	proto.RegisterType((*Result)(nil), "proto.Result")
	proto.RegisterType((*UploadStatus)(nil), "proto.UploadStatus")
}

func init() { proto.RegisterFile("CVMListen.proto", fileDescriptor_24800a83a703ac1f) }

var fileDescriptor_24800a83a703ac1f = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x61, 0x4b, 0xf3, 0x30,
	0x18, 0x5c, 0xcb, 0xda, 0xbe, 0x7b, 0xd6, 0xbd, 0x96, 0x47, 0xc1, 0x32, 0x10, 0x46, 0xf1, 0x43,
	0x51, 0x28, 0xb8, 0xfd, 0x03, 0x2b, 0x13, 0xc1, 0x31, 0x89, 0xd6, 0xef, 0x91, 0x3e, 0x6a, 0x59,
	0x4d, 0x6a, 0x93, 0x2a, 0xfb, 0xf7, 0xb2, 0x34, 0xc2, 0x14, 0x3f, 0x25, 0x77, 0x97, 0xe3, 0x2e,
	0x07, 0x07, 0xf9, 0xe3, 0xea, 0xb6, 0x52, 0x9a, 0x44, 0xd6, 0xb4, 0x52, 0x4b, 0xf4, 0xcc, 0x91,
	0x5c, 0xc3, 0xf8, 0x46, 0x3c, 0xcb, 0x15, 0x29, 0xc5, 0x5f, 0x08, 0x4f, 0x61, 0xb2, 0x6e, 0xa8,
	0xe5, 0xba, 0x92, 0xe2, 0x61, 0xdb, 0x50, 0xec, 0xcc, 0x9c, 0xd4, 0x63, 0x3f, 0x49, 0x44, 0x18,
	0x5e, 0x71, 0xcd, 0x63, 0x77, 0xe6, 0xa4, 0x23, 0x66, 0xee, 0xc9, 0x09, 0x78, 0x97, 0x5b, 0x4d,
	0x0a, 0x8f, 0xc0, 0xfb, 0xe0, 0x75, 0xd7, 0x5b, 0x43, 0xd6, 0x83, 0xe4, 0x1f, 0xf8, 0x8c, 0x54,
	0x57, 0xeb, 0xa4, 0x80, 0xb0, 0x68, 0x6a, 0xc9, 0xcb, 0x7b, 0xcd, 0x75, 0xa7, 0x30, 0x86, 0xc0,
	0xa6, 0x1b, 0xc7, 0x88, 0x7d, 0x43, 0x3c, 0x87, 0x61, 0x2e, 0x4b, 0x32, 0x31, 0xff, 0xe7, 0xc7,
	0x7d, 0xf1, 0x6c, 0xdf, 0xbc, 0x93, 0x99, 0x79, 0x74, 0xb6, 0x80, 0xe8, 0xb7, 0x82, 0x63, 0x08,
	0x0a, 0xb1, 0x11, 0xf2, 0x53, 0x44, 0x03, 0xf4, 0xc1, 0x5d, 0x6f, 0x22, 0x07, 0x01, 0xfc, 0x25,
	0xaf, 0x6a, 0x2a, 0x23, 0x77, 0xfe, 0x0e, 0x61, 0x3f, 0xca, 0xb2, 0x95, 0x6f, 0x77, 0x15, 0x5e,
	0x00, 0xf4, 0x78, 0xb7, 0x09, 0xa2, 0x4d, 0xdc, 0x1b, 0x68, 0x3a, 0xb1, 0x9c, 0xfd, 0xcc, 0x00,
	0x33, 0x08, 0xf2, 0x57, 0x2e, 0x04, 0xd5, 0x18, 0x5a, 0xcd, 0xec, 0x30, 0x3d, 0xfc, 0xa3, 0x6f,
	0xea, 0x3c, 0xf9, 0x86, 0x5d, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x2c, 0x21, 0x95, 0x91,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ListenFromPiClient is the client API for ListenFromPi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ListenFromPiClient interface {
	// PI 端推信息到CVM
	ListenInfo(ctx context.Context, in *InfoMessage, opts ...grpc.CallOption) (*Result, error)
	Channel(ctx context.Context, opts ...grpc.CallOption) (ListenFromPi_ChannelClient, error)
}

type listenFromPiClient struct {
	cc *grpc.ClientConn
}

func NewListenFromPiClient(cc *grpc.ClientConn) ListenFromPiClient {
	return &listenFromPiClient{cc}
}

func (c *listenFromPiClient) ListenInfo(ctx context.Context, in *InfoMessage, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/proto.ListenFromPi/ListenInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenFromPiClient) Channel(ctx context.Context, opts ...grpc.CallOption) (ListenFromPi_ChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ListenFromPi_serviceDesc.Streams[0], "/proto.ListenFromPi/Channel", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenFromPiChannelClient{stream}
	return x, nil
}

type ListenFromPi_ChannelClient interface {
	Send(*Bytes) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type listenFromPiChannelClient struct {
	grpc.ClientStream
}

func (x *listenFromPiChannelClient) Send(m *Bytes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenFromPiChannelClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ListenFromPiServer is the server API for ListenFromPi service.
type ListenFromPiServer interface {
	// PI 端推信息到CVM
	ListenInfo(context.Context, *InfoMessage) (*Result, error)
	Channel(ListenFromPi_ChannelServer) error
}

func RegisterListenFromPiServer(s *grpc.Server, srv ListenFromPiServer) {
	s.RegisterService(&_ListenFromPi_serviceDesc, srv)
}

func _ListenFromPi_ListenInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenFromPiServer).ListenInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ListenFromPi/ListenInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenFromPiServer).ListenInfo(ctx, req.(*InfoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenFromPi_Channel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenFromPiServer).Channel(&listenFromPiChannelServer{stream})
}

type ListenFromPi_ChannelServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Bytes, error)
	grpc.ServerStream
}

type listenFromPiChannelServer struct {
	grpc.ServerStream
}

func (x *listenFromPiChannelServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenFromPiChannelServer) Recv() (*Bytes, error) {
	m := new(Bytes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ListenFromPi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ListenFromPi",
	HandlerType: (*ListenFromPiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListenInfo",
			Handler:    _ListenFromPi_ListenInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Channel",
			Handler:       _ListenFromPi_Channel_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "CVMListen.proto",
}