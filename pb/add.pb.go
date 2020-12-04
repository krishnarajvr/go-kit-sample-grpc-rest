// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The sum request contains two parameters.
type SumRequest struct {
	A                    int64    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_077cd88a1973142f, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SumRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

// The sum response contains the result of the calculation.
type SumReply struct {
	V                    int64    `protobuf:"varint,1,opt,name=v,proto3" json:"v,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumReply) Reset()         { *m = SumReply{} }
func (m *SumReply) String() string { return proto.CompactTextString(m) }
func (*SumReply) ProtoMessage()    {}
func (*SumReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_077cd88a1973142f, []int{1}
}

func (m *SumReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumReply.Unmarshal(m, b)
}
func (m *SumReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumReply.Marshal(b, m, deterministic)
}
func (m *SumReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumReply.Merge(m, src)
}
func (m *SumReply) XXX_Size() int {
	return xxx_messageInfo_SumReply.Size(m)
}
func (m *SumReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SumReply.DiscardUnknown(m)
}

var xxx_messageInfo_SumReply proto.InternalMessageInfo

func (m *SumReply) GetV() int64 {
	if m != nil {
		return m.V
	}
	return 0
}

// The Concat request contains two parameters.
type ConcatRequest struct {
	A                    string   `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    string   `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConcatRequest) Reset()         { *m = ConcatRequest{} }
func (m *ConcatRequest) String() string { return proto.CompactTextString(m) }
func (*ConcatRequest) ProtoMessage()    {}
func (*ConcatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_077cd88a1973142f, []int{2}
}

func (m *ConcatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcatRequest.Unmarshal(m, b)
}
func (m *ConcatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcatRequest.Marshal(b, m, deterministic)
}
func (m *ConcatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcatRequest.Merge(m, src)
}
func (m *ConcatRequest) XXX_Size() int {
	return xxx_messageInfo_ConcatRequest.Size(m)
}
func (m *ConcatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConcatRequest proto.InternalMessageInfo

func (m *ConcatRequest) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *ConcatRequest) GetB() string {
	if m != nil {
		return m.B
	}
	return ""
}

// The Concat response contains the result of the concatenation.
type ConcatReply struct {
	V                    string   `protobuf:"bytes,1,opt,name=v,proto3" json:"v,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConcatReply) Reset()         { *m = ConcatReply{} }
func (m *ConcatReply) String() string { return proto.CompactTextString(m) }
func (*ConcatReply) ProtoMessage()    {}
func (*ConcatReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_077cd88a1973142f, []int{3}
}

func (m *ConcatReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcatReply.Unmarshal(m, b)
}
func (m *ConcatReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcatReply.Marshal(b, m, deterministic)
}
func (m *ConcatReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcatReply.Merge(m, src)
}
func (m *ConcatReply) XXX_Size() int {
	return xxx_messageInfo_ConcatReply.Size(m)
}
func (m *ConcatReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcatReply.DiscardUnknown(m)
}

var xxx_messageInfo_ConcatReply proto.InternalMessageInfo

func (m *ConcatReply) GetV() string {
	if m != nil {
		return m.V
	}
	return ""
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "pb.SumRequest")
	proto.RegisterType((*SumReply)(nil), "pb.SumReply")
	proto.RegisterType((*ConcatRequest)(nil), "pb.ConcatRequest")
	proto.RegisterType((*ConcatReply)(nil), "pb.ConcatReply")
}

func init() {
	proto.RegisterFile("add.proto", fileDescriptor_077cd88a1973142f)
}

var fileDescriptor_077cd88a1973142f = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x4c, 0x49, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe0, 0xe2, 0x0a, 0x2e, 0xcd,
	0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe1, 0x62, 0x4c, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0e, 0x62, 0x4c, 0x04, 0xf1, 0x92, 0x24, 0x98, 0x20, 0xbc, 0x24, 0x25, 0x09, 0x2e,
	0x0e, 0xb0, 0xca, 0x82, 0x9c, 0x4a, 0x90, 0x4c, 0x19, 0x4c, 0x5d, 0x99, 0x92, 0x36, 0x17, 0xaf,
	0x73, 0x7e, 0x5e, 0x72, 0x62, 0x09, 0x86, 0x31, 0x9c, 0x28, 0xc6, 0x70, 0x82, 0x8c, 0x91, 0xe6,
	0xe2, 0x86, 0x29, 0x46, 0x31, 0x09, 0x28, 0x59, 0x66, 0x14, 0xc3, 0xc5, 0xec, 0x98, 0x92, 0x22,
	0xa4, 0xca, 0xc5, 0x0c, 0xb4, 0x4a, 0x88, 0x4f, 0xaf, 0x20, 0x49, 0x0f, 0xe1, 0x3a, 0x29, 0x1e,
	0x38, 0x1f, 0xa8, 0x53, 0x89, 0x41, 0x48, 0x8f, 0x8b, 0x0d, 0x62, 0x94, 0x90, 0x20, 0x48, 0x06,
	0xc5, 0x0d, 0x52, 0xfc, 0xc8, 0x42, 0x60, 0xf5, 0x49, 0x6c, 0x60, 0x6f, 0x1b, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xb4, 0xc9, 0xe7, 0x58, 0x03, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AddClient is the client API for Add service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddClient interface {
	// Sums two integers.
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error)
	// Concatenates two strings
	Concat(ctx context.Context, in *ConcatRequest, opts ...grpc.CallOption) (*ConcatReply, error)
}

type addClient struct {
	cc grpc.ClientConnInterface
}

func NewAddClient(cc grpc.ClientConnInterface) AddClient {
	return &addClient{cc}
}

func (c *addClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error) {
	out := new(SumReply)
	err := c.cc.Invoke(ctx, "/pb.Add/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addClient) Concat(ctx context.Context, in *ConcatRequest, opts ...grpc.CallOption) (*ConcatReply, error) {
	out := new(ConcatReply)
	err := c.cc.Invoke(ctx, "/pb.Add/Concat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServer is the server API for Add service.
type AddServer interface {
	// Sums two integers.
	Sum(context.Context, *SumRequest) (*SumReply, error)
	// Concatenates two strings
	Concat(context.Context, *ConcatRequest) (*ConcatReply, error)
}

// UnimplementedAddServer can be embedded to have forward compatible implementations.
type UnimplementedAddServer struct {
}

func (*UnimplementedAddServer) Sum(ctx context.Context, req *SumRequest) (*SumReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedAddServer) Concat(ctx context.Context, req *ConcatRequest) (*ConcatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Concat not implemented")
}

func RegisterAddServer(s *grpc.Server, srv AddServer) {
	s.RegisterService(&_Add_serviceDesc, srv)
}

func _Add_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Add/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Add_Concat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConcatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServer).Concat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Add/Concat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServer).Concat(ctx, req.(*ConcatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Add_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Add",
	HandlerType: (*AddServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Add_Sum_Handler,
		},
		{
			MethodName: "Concat",
			Handler:    _Add_Concat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "add.proto",
}