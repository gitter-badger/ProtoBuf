// Code generated by protoc-gen-go.
// source: proto/test/test.proto
// DO NOT EDIT!

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	proto/test/test.proto

It has these top-level messages:
	ExtraStuff
	PingRequest
	PingResponse
*/
package test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type PingRequest_FailureType int32

const (
	PingRequest_NONE PingRequest_FailureType = 0
	PingRequest_CODE PingRequest_FailureType = 1
	PingRequest_DROP PingRequest_FailureType = 2
)

var PingRequest_FailureType_name = map[int32]string{
	0: "NONE",
	1: "CODE",
	2: "DROP",
}
var PingRequest_FailureType_value = map[string]int32{
	"NONE": 0,
	"CODE": 1,
	"DROP": 2,
}

func (x PingRequest_FailureType) String() string {
	return proto.EnumName(PingRequest_FailureType_name, int32(x))
}
func (PingRequest_FailureType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type ExtraStuff struct {
	Addresses map[int32]string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to Title:
	//	*ExtraStuff_FirstName
	//	*ExtraStuff_IdNumber
	Title       isExtraStuff_Title `protobuf_oneof:"title"`
	CardNumbers []uint32           `protobuf:"varint,4,rep,packed,name=card_numbers,json=cardNumbers" json:"card_numbers,omitempty"`
}

func (m *ExtraStuff) Reset()                    { *m = ExtraStuff{} }
func (m *ExtraStuff) String() string            { return proto.CompactTextString(m) }
func (*ExtraStuff) ProtoMessage()               {}
func (*ExtraStuff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isExtraStuff_Title interface {
	isExtraStuff_Title()
}

type ExtraStuff_FirstName struct {
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,oneof"`
}
type ExtraStuff_IdNumber struct {
	IdNumber int32 `protobuf:"varint,3,opt,name=id_number,json=idNumber,oneof"`
}

func (*ExtraStuff_FirstName) isExtraStuff_Title() {}
func (*ExtraStuff_IdNumber) isExtraStuff_Title()  {}

func (m *ExtraStuff) GetTitle() isExtraStuff_Title {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *ExtraStuff) GetAddresses() map[int32]string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ExtraStuff) GetFirstName() string {
	if x, ok := m.GetTitle().(*ExtraStuff_FirstName); ok {
		return x.FirstName
	}
	return ""
}

func (m *ExtraStuff) GetIdNumber() int32 {
	if x, ok := m.GetTitle().(*ExtraStuff_IdNumber); ok {
		return x.IdNumber
	}
	return 0
}

func (m *ExtraStuff) GetCardNumbers() []uint32 {
	if m != nil {
		return m.CardNumbers
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ExtraStuff) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ExtraStuff_OneofMarshaler, _ExtraStuff_OneofUnmarshaler, _ExtraStuff_OneofSizer, []interface{}{
		(*ExtraStuff_FirstName)(nil),
		(*ExtraStuff_IdNumber)(nil),
	}
}

func _ExtraStuff_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ExtraStuff)
	// title
	switch x := m.Title.(type) {
	case *ExtraStuff_FirstName:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.FirstName)
	case *ExtraStuff_IdNumber:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IdNumber))
	case nil:
	default:
		return fmt.Errorf("ExtraStuff.Title has unexpected type %T", x)
	}
	return nil
}

func _ExtraStuff_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ExtraStuff)
	switch tag {
	case 2: // title.first_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Title = &ExtraStuff_FirstName{x}
		return true, err
	case 3: // title.id_number
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Title = &ExtraStuff_IdNumber{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _ExtraStuff_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ExtraStuff)
	// title
	switch x := m.Title.(type) {
	case *ExtraStuff_FirstName:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.FirstName)))
		n += len(x.FirstName)
	case *ExtraStuff_IdNumber:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IdNumber))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PingRequest struct {
	Value             string                  `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	ResponseCount     int32                   `protobuf:"varint,2,opt,name=response_count,json=responseCount" json:"response_count,omitempty"`
	ErrorCodeReturned uint32                  `protobuf:"varint,3,opt,name=error_code_returned,json=errorCodeReturned" json:"error_code_returned,omitempty"`
	FailureType       PingRequest_FailureType `protobuf:"varint,4,opt,name=failure_type,json=failureType,enum=test.PingRequest_FailureType" json:"failure_type,omitempty"`
	CheckMetadata     bool                    `protobuf:"varint,5,opt,name=check_metadata,json=checkMetadata" json:"check_metadata,omitempty"`
	SendHeaders       bool                    `protobuf:"varint,6,opt,name=send_headers,json=sendHeaders" json:"send_headers,omitempty"`
	SendTrailers      bool                    `protobuf:"varint,7,opt,name=send_trailers,json=sendTrailers" json:"send_trailers,omitempty"`
	MessageLatencyMs  int32                   `protobuf:"varint,8,opt,name=message_latency_ms,json=messageLatencyMs" json:"message_latency_ms,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingRequest) GetResponseCount() int32 {
	if m != nil {
		return m.ResponseCount
	}
	return 0
}

func (m *PingRequest) GetErrorCodeReturned() uint32 {
	if m != nil {
		return m.ErrorCodeReturned
	}
	return 0
}

func (m *PingRequest) GetFailureType() PingRequest_FailureType {
	if m != nil {
		return m.FailureType
	}
	return PingRequest_NONE
}

func (m *PingRequest) GetCheckMetadata() bool {
	if m != nil {
		return m.CheckMetadata
	}
	return false
}

func (m *PingRequest) GetSendHeaders() bool {
	if m != nil {
		return m.SendHeaders
	}
	return false
}

func (m *PingRequest) GetSendTrailers() bool {
	if m != nil {
		return m.SendTrailers
	}
	return false
}

func (m *PingRequest) GetMessageLatencyMs() int32 {
	if m != nil {
		return m.MessageLatencyMs
	}
	return 0
}

type PingResponse struct {
	Value   string `protobuf:"bytes,1,opt,name=Value" json:"Value,omitempty"`
	Counter int32  `protobuf:"varint,2,opt,name=counter" json:"counter,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PingResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingResponse) GetCounter() int32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func init() {
	proto.RegisterType((*ExtraStuff)(nil), "test.ExtraStuff")
	proto.RegisterType((*PingRequest)(nil), "test.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "test.PingResponse")
	proto.RegisterEnum("test.PingRequest_FailureType", PingRequest_FailureType_name, PingRequest_FailureType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	PingEmpty(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) PingEmpty(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/test.TestService/PingEmpty", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/test.TestService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/test.TestService/PingError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/test.TestService/PingList", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_PingListClient interface {
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingListClient struct {
	grpc.ClientStream
}

func (x *testServicePingListClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	PingEmpty(context.Context, *google_protobuf.Empty) (*PingResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	PingError(context.Context, *PingRequest) (*google_protobuf.Empty, error)
	PingList(*PingRequest, TestService_PingListServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_PingEmpty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingEmpty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/PingEmpty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingEmpty(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/PingError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingError(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).PingList(m, &testServicePingListServer{stream})
}

type TestService_PingListServer interface {
	Send(*PingResponse) error
	grpc.ServerStream
}

type testServicePingListServer struct {
	grpc.ServerStream
}

func (x *testServicePingListServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingEmpty",
			Handler:    _TestService_PingEmpty_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _TestService_Ping_Handler,
		},
		{
			MethodName: "PingError",
			Handler:    _TestService_PingError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingList",
			Handler:       _TestService_PingList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/test/test.proto",
}

// Client API for FailService service

type FailServiceClient interface {
	NonExistant(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type failServiceClient struct {
	cc *grpc.ClientConn
}

func NewFailServiceClient(cc *grpc.ClientConn) FailServiceClient {
	return &failServiceClient{cc}
}

func (c *failServiceClient) NonExistant(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/test.FailService/NonExistant", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FailService service

type FailServiceServer interface {
	NonExistant(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterFailServiceServer(s *grpc.Server, srv FailServiceServer) {
	s.RegisterService(&_FailService_serviceDesc, srv)
}

func _FailService_NonExistant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FailServiceServer).NonExistant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.FailService/NonExistant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FailServiceServer).NonExistant(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.FailService",
	HandlerType: (*FailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NonExistant",
			Handler:    _FailService_NonExistant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/test/test.proto",
}

func init() { proto.RegisterFile("proto/test/test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x6b, 0x13, 0x41,
	0x10, 0xcf, 0x36, 0x49, 0x9b, 0xcc, 0x35, 0x25, 0x5d, 0x3f, 0x38, 0x2a, 0xa5, 0x31, 0x22, 0x04,
	0x94, 0x8b, 0x44, 0x94, 0x2a, 0x2a, 0x6a, 0x7b, 0xd2, 0x87, 0x36, 0x2d, 0xdb, 0xe2, 0xeb, 0xb1,
	0xbd, 0x9b, 0xa4, 0x47, 0xef, 0x23, 0xee, 0xee, 0x95, 0xde, 0x3f, 0xe0, 0x1f, 0xeb, 0xbb, 0xef,
	0xb2, 0xbb, 0x77, 0x24, 0xa2, 0x85, 0xbe, 0x2c, 0x33, 0xbf, 0xf9, 0xcd, 0xc7, 0xfe, 0x66, 0xe0,
	0xd1, 0x42, 0xe4, 0x2a, 0x1f, 0x2b, 0x94, 0xca, 0x3c, 0x9e, 0xf1, 0x69, 0x4b, 0xdb, 0x3b, 0x4f,
	0xe6, 0x79, 0x3e, 0x4f, 0x70, 0x6c, 0xb0, 0xcb, 0x62, 0x36, 0xc6, 0x74, 0xa1, 0x4a, 0x4b, 0x19,
	0xfe, 0x26, 0x00, 0xfe, 0xad, 0x12, 0xfc, 0x5c, 0x15, 0xb3, 0x19, 0xfd, 0x08, 0x5d, 0x1e, 0x45,
	0x02, 0xa5, 0x44, 0xe9, 0x92, 0x41, 0x73, 0xe4, 0x4c, 0xf6, 0x3c, 0x53, 0x71, 0x49, 0xf2, 0xbe,
	0xd4, 0x0c, 0x3f, 0x53, 0xa2, 0x64, 0xcb, 0x0c, 0xba, 0x07, 0x30, 0x8b, 0x85, 0x54, 0x41, 0xc6,
	0x53, 0x74, 0xd7, 0x06, 0x64, 0xd4, 0x3d, 0x6a, 0xb0, 0xae, 0xc1, 0xa6, 0x3c, 0x45, 0xba, 0x0b,
	0xdd, 0x38, 0x0a, 0xb2, 0x22, 0xbd, 0x44, 0xe1, 0x36, 0x07, 0x64, 0xd4, 0x3e, 0x6a, 0xb0, 0x4e,
	0x1c, 0x4d, 0x0d, 0x42, 0x9f, 0xc2, 0x66, 0xc8, 0x45, 0x4d, 0x90, 0x6e, 0x6b, 0xd0, 0x1c, 0xf5,
	0x98, 0xa3, 0x31, 0xcb, 0x90, 0x3b, 0x1f, 0x60, 0xeb, 0xef, 0xfe, 0xb4, 0x0f, 0xcd, 0x6b, 0x2c,
	0x5d, 0xa2, 0xab, 0x31, 0x6d, 0xd2, 0x87, 0xd0, 0xbe, 0xe1, 0x49, 0x51, 0x4d, 0xc0, 0xac, 0xf3,
	0x7e, 0x6d, 0x9f, 0x7c, 0xdd, 0x80, 0xb6, 0x8a, 0x55, 0x82, 0xc3, 0x9f, 0x4d, 0x70, 0xce, 0xe2,
	0x6c, 0xce, 0xf0, 0x47, 0x81, 0x52, 0x2d, 0x53, 0xc8, 0x4a, 0x0a, 0x7d, 0x0e, 0x5b, 0x02, 0xe5,
	0x22, 0xcf, 0x24, 0x06, 0x61, 0x5e, 0x64, 0xca, 0x54, 0x6c, 0xb3, 0x5e, 0x8d, 0x1e, 0x68, 0x90,
	0x7a, 0xf0, 0x00, 0x85, 0xc8, 0x45, 0x10, 0xe6, 0x11, 0x06, 0x02, 0x55, 0x21, 0x32, 0x8c, 0xcc,
	0xff, 0x7a, 0x6c, 0xdb, 0x84, 0x0e, 0xf2, 0x08, 0x59, 0x15, 0xa0, 0x9f, 0x61, 0x73, 0xc6, 0xe3,
	0xa4, 0x10, 0x18, 0xa8, 0x72, 0x81, 0x6e, 0x6b, 0x40, 0x46, 0x5b, 0x93, 0x5d, 0x2b, 0xf4, 0xca,
	0x54, 0xde, 0x37, 0xcb, 0xba, 0x28, 0x17, 0xc8, 0x9c, 0xd9, 0xd2, 0xd1, 0x83, 0x85, 0x57, 0x18,
	0x5e, 0x07, 0x29, 0x2a, 0x1e, 0x71, 0xc5, 0xdd, 0xf6, 0x80, 0x8c, 0x3a, 0xac, 0x67, 0xd0, 0x93,
	0x0a, 0xd4, 0x7a, 0x4a, 0xcc, 0xa2, 0xe0, 0x0a, 0x79, 0xa4, 0xf5, 0x5c, 0x37, 0x24, 0x47, 0x63,
	0x47, 0x16, 0xa2, 0xcf, 0xa0, 0x67, 0x28, 0x4a, 0xf0, 0x38, 0xd1, 0x9c, 0x0d, 0xc3, 0x31, 0x79,
	0x17, 0x15, 0x46, 0x5f, 0x02, 0x4d, 0x51, 0x4a, 0x3e, 0xc7, 0x20, 0xe1, 0x0a, 0xb3, 0xb0, 0x0c,
	0x52, 0xe9, 0x76, 0x8c, 0x16, 0xfd, 0x2a, 0x72, 0x6c, 0x03, 0x27, 0x72, 0xf8, 0x02, 0x9c, 0x95,
	0xc1, 0x69, 0x07, 0x5a, 0xd3, 0xd3, 0xa9, 0xdf, 0x6f, 0x68, 0xeb, 0xe0, 0xf4, 0xd0, 0xef, 0x13,
	0x6d, 0x1d, 0xb2, 0xd3, 0xb3, 0xfe, 0xda, 0xf0, 0x13, 0x6c, 0xda, 0x1f, 0x5b, 0x41, 0xf5, 0x22,
	0xbe, 0xaf, 0x2e, 0xc2, 0x38, 0xd4, 0x85, 0x0d, 0xa3, 0x3f, 0x8a, 0x6a, 0x03, 0xb5, 0x3b, 0xf9,
	0x45, 0xc0, 0xb9, 0x40, 0xa9, 0xce, 0x51, 0xdc, 0xc4, 0x21, 0xd2, 0x77, 0xd0, 0xd5, 0xf5, 0x7c,
	0x7d, 0xe3, 0xf4, 0xb1, 0x67, 0x6f, 0xdf, 0xab, 0x6f, 0xdf, 0x33, 0xf8, 0x0e, 0x5d, 0x95, 0xda,
	0x36, 0x1e, 0x36, 0xe8, 0x18, 0x5a, 0x1a, 0xa1, 0xdb, 0xff, 0x2c, 0xe2, 0x8e, 0x84, 0xfd, 0xaa,
	0x97, 0x5e, 0xf0, 0xff, 0xb2, 0xee, 0x68, 0x3f, 0x6c, 0xd0, 0x37, 0xd0, 0xd1, 0xc4, 0xe3, 0x58,
	0xaa, 0x7b, 0xb7, 0x7b, 0x45, 0x26, 0xbe, 0x55, 0xb6, 0xfe, 0xeb, 0x5b, 0x70, 0xa6, 0x79, 0xe6,
	0xdf, 0xc6, 0x52, 0xf1, 0xec, 0xfe, 0x85, 0x2e, 0xd7, 0xcd, 0x3c, 0xaf, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x77, 0x09, 0x4e, 0xb9, 0x37, 0x04, 0x00, 0x00,
}
