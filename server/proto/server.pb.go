// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	server.proto

It has these top-level messages:
	ExecuteRequest
	RunRequest
	ServerEmpty
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ExecuteRequest struct {
}

func (m *ExecuteRequest) Reset()                    { *m = ExecuteRequest{} }
func (m *ExecuteRequest) String() string            { return proto1.CompactTextString(m) }
func (*ExecuteRequest) ProtoMessage()               {}
func (*ExecuteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RunRequest struct {
}

func (m *RunRequest) Reset()                    { *m = RunRequest{} }
func (m *RunRequest) String() string            { return proto1.CompactTextString(m) }
func (*RunRequest) ProtoMessage()               {}
func (*RunRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ServerEmpty struct {
}

func (m *ServerEmpty) Reset()                    { *m = ServerEmpty{} }
func (m *ServerEmpty) String() string            { return proto1.CompactTextString(m) }
func (*ServerEmpty) ProtoMessage()               {}
func (*ServerEmpty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto1.RegisterType((*ExecuteRequest)(nil), "proto.ExecuteRequest")
	proto1.RegisterType((*RunRequest)(nil), "proto.RunRequest")
	proto1.RegisterType((*ServerEmpty)(nil), "proto.ServerEmpty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BenchServer service

type BenchServerClient interface {
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ServerEmpty, error)
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*ServerEmpty, error)
}

type benchServerClient struct {
	cc *grpc.ClientConn
}

func NewBenchServerClient(cc *grpc.ClientConn) BenchServerClient {
	return &benchServerClient{cc}
}

func (c *benchServerClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ServerEmpty, error) {
	out := new(ServerEmpty)
	err := grpc.Invoke(ctx, "/proto.BenchServer/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *benchServerClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*ServerEmpty, error) {
	out := new(ServerEmpty)
	err := grpc.Invoke(ctx, "/proto.BenchServer/Run", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BenchServer service

type BenchServerServer interface {
	Execute(context.Context, *ExecuteRequest) (*ServerEmpty, error)
	Run(context.Context, *RunRequest) (*ServerEmpty, error)
}

func RegisterBenchServerServer(s *grpc.Server, srv BenchServerServer) {
	s.RegisterService(&_BenchServer_serviceDesc, srv)
}

func _BenchServer_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenchServerServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BenchServer/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenchServerServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BenchServer_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenchServerServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BenchServer/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenchServerServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BenchServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BenchServer",
	HandlerType: (*BenchServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _BenchServer_Execute_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _BenchServer_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

func init() { proto1.RegisterFile("server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x02, 0x5c, 0x7c,
	0xae, 0x15, 0xa9, 0xc9, 0xa5, 0x25, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x4a, 0x3c,
	0x5c, 0x5c, 0x41, 0xa5, 0x79, 0x30, 0x1e, 0x2f, 0x17, 0x77, 0x30, 0x58, 0x9b, 0x6b, 0x6e, 0x41,
	0x49, 0xa5, 0x51, 0x21, 0x17, 0xb7, 0x53, 0x6a, 0x5e, 0x72, 0x06, 0x44, 0x4c, 0xc8, 0x84, 0x8b,
	0x1d, 0xaa, 0x5b, 0x48, 0x14, 0x62, 0xae, 0x1e, 0xaa, 0x69, 0x52, 0x42, 0x50, 0x61, 0x24, 0x43,
	0x84, 0x74, 0xb8, 0x98, 0x83, 0x4a, 0xf3, 0x84, 0x04, 0xa1, 0x52, 0x08, 0xdb, 0xb0, 0xa9, 0x4e,
	0x62, 0x03, 0x0b, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xca, 0x9e, 0xa2, 0xbf, 0x00,
	0x00, 0x00,
}
