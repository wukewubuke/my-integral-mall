// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goods.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//protoc --go_out=plugins=grpc:. *.proto
type AddIntegralRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Integral             int64    `protobuf:"varint,2,opt,name=integral,proto3" json:"integral,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddIntegralRequest) Reset()         { *m = AddIntegralRequest{} }
func (m *AddIntegralRequest) String() string { return proto.CompactTextString(m) }
func (*AddIntegralRequest) ProtoMessage()    {}
func (*AddIntegralRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a30593c5487368b0, []int{0}
}

func (m *AddIntegralRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddIntegralRequest.Unmarshal(m, b)
}
func (m *AddIntegralRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddIntegralRequest.Marshal(b, m, deterministic)
}
func (m *AddIntegralRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddIntegralRequest.Merge(m, src)
}
func (m *AddIntegralRequest) XXX_Size() int {
	return xxx_messageInfo_AddIntegralRequest.Size(m)
}
func (m *AddIntegralRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddIntegralRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddIntegralRequest proto.InternalMessageInfo

func (m *AddIntegralRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AddIntegralRequest) GetIntegral() int64 {
	if m != nil {
		return m.Integral
	}
	return 0
}

type IntegralResponse struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Integral             int64    `protobuf:"varint,2,opt,name=integral,proto3" json:"integral,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegralResponse) Reset()         { *m = IntegralResponse{} }
func (m *IntegralResponse) String() string { return proto.CompactTextString(m) }
func (*IntegralResponse) ProtoMessage()    {}
func (*IntegralResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a30593c5487368b0, []int{1}
}

func (m *IntegralResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegralResponse.Unmarshal(m, b)
}
func (m *IntegralResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegralResponse.Marshal(b, m, deterministic)
}
func (m *IntegralResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegralResponse.Merge(m, src)
}
func (m *IntegralResponse) XXX_Size() int {
	return xxx_messageInfo_IntegralResponse.Size(m)
}
func (m *IntegralResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegralResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IntegralResponse proto.InternalMessageInfo

func (m *IntegralResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IntegralResponse) GetIntegral() int64 {
	if m != nil {
		return m.Integral
	}
	return 0
}

type ConsumerIntegralRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ConsumerIntegral     int64    `protobuf:"varint,2,opt,name=consumerIntegral,proto3" json:"consumerIntegral,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumerIntegralRequest) Reset()         { *m = ConsumerIntegralRequest{} }
func (m *ConsumerIntegralRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumerIntegralRequest) ProtoMessage()    {}
func (*ConsumerIntegralRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a30593c5487368b0, []int{2}
}

func (m *ConsumerIntegralRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumerIntegralRequest.Unmarshal(m, b)
}
func (m *ConsumerIntegralRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumerIntegralRequest.Marshal(b, m, deterministic)
}
func (m *ConsumerIntegralRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerIntegralRequest.Merge(m, src)
}
func (m *ConsumerIntegralRequest) XXX_Size() int {
	return xxx_messageInfo_ConsumerIntegralRequest.Size(m)
}
func (m *ConsumerIntegralRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerIntegralRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerIntegralRequest proto.InternalMessageInfo

func (m *ConsumerIntegralRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ConsumerIntegralRequest) GetConsumerIntegral() int64 {
	if m != nil {
		return m.ConsumerIntegral
	}
	return 0
}

type FindOneByUserIdRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindOneByUserIdRequest) Reset()         { *m = FindOneByUserIdRequest{} }
func (m *FindOneByUserIdRequest) String() string { return proto.CompactTextString(m) }
func (*FindOneByUserIdRequest) ProtoMessage()    {}
func (*FindOneByUserIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a30593c5487368b0, []int{3}
}

func (m *FindOneByUserIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindOneByUserIdRequest.Unmarshal(m, b)
}
func (m *FindOneByUserIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindOneByUserIdRequest.Marshal(b, m, deterministic)
}
func (m *FindOneByUserIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindOneByUserIdRequest.Merge(m, src)
}
func (m *FindOneByUserIdRequest) XXX_Size() int {
	return xxx_messageInfo_FindOneByUserIdRequest.Size(m)
}
func (m *FindOneByUserIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindOneByUserIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindOneByUserIdRequest proto.InternalMessageInfo

func (m *FindOneByUserIdRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*AddIntegralRequest)(nil), "protos.AddIntegralRequest")
	proto.RegisterType((*IntegralResponse)(nil), "protos.IntegralResponse")
	proto.RegisterType((*ConsumerIntegralRequest)(nil), "protos.ConsumerIntegralRequest")
	proto.RegisterType((*FindOneByUserIdRequest)(nil), "protos.FindOneByUserIdRequest")
}

func init() { proto.RegisterFile("goods.proto", fileDescriptor_a30593c5487368b0) }

var fileDescriptor_a30593c5487368b0 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcf, 0xcf, 0x4f,
	0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0x1e, 0x5c, 0x42,
	0x8e, 0x29, 0x29, 0x9e, 0x79, 0x25, 0xa9, 0xe9, 0x45, 0x89, 0x39, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0xa5, 0xc5, 0xa9, 0x45, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xcc, 0x41, 0x50, 0x9e, 0x90, 0x14, 0x17, 0x47, 0x26, 0x54, 0xa9, 0x04, 0x13, 0x58, 0x06,
	0xce, 0x57, 0x72, 0xe3, 0x12, 0x40, 0x18, 0x53, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x4a, 0x96, 0x39,
	0xb1, 0x5c, 0xe2, 0xce, 0xf9, 0x79, 0xc5, 0xa5, 0xb9, 0xa9, 0x45, 0xc4, 0x3a, 0x4b, 0x8b, 0x4b,
	0x20, 0x19, 0x4d, 0x0b, 0xd4, 0x58, 0x0c, 0x71, 0x25, 0x03, 0x2e, 0x31, 0xb7, 0xcc, 0xbc, 0x14,
	0xff, 0xbc, 0x54, 0xa7, 0xca, 0x50, 0xb0, 0x76, 0x02, 0xa6, 0x1b, 0xbd, 0x63, 0xe4, 0xe2, 0x86,
	0xbb, 0xa4, 0x20, 0x59, 0xc8, 0x99, 0x8b, 0x1b, 0x29, 0xc8, 0x84, 0xa4, 0x20, 0x21, 0x5a, 0xac,
	0x87, 0x19, 0x8e, 0x52, 0x12, 0x30, 0x39, 0x8c, 0x90, 0xf1, 0xe5, 0x12, 0x40, 0xf7, 0xa5, 0x90,
	0x3c, 0x4c, 0x35, 0x0e, 0xff, 0xe3, 0x31, 0xce, 0x9b, 0x8b, 0x1f, 0xcd, 0x57, 0x42, 0x72, 0x30,
	0xc5, 0xd8, 0xbd, 0x8b, 0xdb, 0xb0, 0x24, 0x48, 0xda, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x78, 0x03, 0x7c, 0x71, 0x31, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IntegralRpcClient is the client API for IntegralRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IntegralRpcClient interface {
	AddIntegral(ctx context.Context, in *AddIntegralRequest, opts ...grpc.CallOption) (*IntegralResponse, error)
	ConsumerIntegral(ctx context.Context, in *ConsumerIntegralRequest, opts ...grpc.CallOption) (*IntegralResponse, error)
	FindOneByUserId(ctx context.Context, in *FindOneByUserIdRequest, opts ...grpc.CallOption) (*IntegralResponse, error)
}

type integralRpcClient struct {
	cc *grpc.ClientConn
}

func NewIntegralRpcClient(cc *grpc.ClientConn) IntegralRpcClient {
	return &integralRpcClient{cc}
}

func (c *integralRpcClient) AddIntegral(ctx context.Context, in *AddIntegralRequest, opts ...grpc.CallOption) (*IntegralResponse, error) {
	out := new(IntegralResponse)
	err := c.cc.Invoke(ctx, "/protos.IntegralRpc/AddIntegral", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralRpcClient) ConsumerIntegral(ctx context.Context, in *ConsumerIntegralRequest, opts ...grpc.CallOption) (*IntegralResponse, error) {
	out := new(IntegralResponse)
	err := c.cc.Invoke(ctx, "/protos.IntegralRpc/ConsumerIntegral", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralRpcClient) FindOneByUserId(ctx context.Context, in *FindOneByUserIdRequest, opts ...grpc.CallOption) (*IntegralResponse, error) {
	out := new(IntegralResponse)
	err := c.cc.Invoke(ctx, "/protos.IntegralRpc/FindOneByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntegralRpcServer is the server API for IntegralRpc service.
type IntegralRpcServer interface {
	AddIntegral(context.Context, *AddIntegralRequest) (*IntegralResponse, error)
	ConsumerIntegral(context.Context, *ConsumerIntegralRequest) (*IntegralResponse, error)
	FindOneByUserId(context.Context, *FindOneByUserIdRequest) (*IntegralResponse, error)
}

func RegisterIntegralRpcServer(s *grpc.Server, srv IntegralRpcServer) {
	s.RegisterService(&_IntegralRpc_serviceDesc, srv)
}

func _IntegralRpc_AddIntegral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddIntegralRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegralRpcServer).AddIntegral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.IntegralRpc/AddIntegral",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegralRpcServer).AddIntegral(ctx, req.(*AddIntegralRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegralRpc_ConsumerIntegral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumerIntegralRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegralRpcServer).ConsumerIntegral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.IntegralRpc/ConsumerIntegral",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegralRpcServer).ConsumerIntegral(ctx, req.(*ConsumerIntegralRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegralRpc_FindOneByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegralRpcServer).FindOneByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.IntegralRpc/FindOneByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegralRpcServer).FindOneByUserId(ctx, req.(*FindOneByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IntegralRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.IntegralRpc",
	HandlerType: (*IntegralRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddIntegral",
			Handler:    _IntegralRpc_AddIntegral_Handler,
		},
		{
			MethodName: "ConsumerIntegral",
			Handler:    _IntegralRpc_ConsumerIntegral_Handler,
		},
		{
			MethodName: "FindOneByUserId",
			Handler:    _IntegralRpc_FindOneByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goods.proto",
}
