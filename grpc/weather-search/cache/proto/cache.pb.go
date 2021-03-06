// Code generated by protoc-gen-go.
// source: proto/cache.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/cache.proto

It has these top-level messages:
	LookupRequest
	LookupResponse
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

type LookupRequest struct {
	Key int64 `protobuf:"varint,1,opt,name=key" json:"key,omitempty"`
}

func (m *LookupRequest) Reset()                    { *m = LookupRequest{} }
func (m *LookupRequest) String() string            { return proto1.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()               {}
func (*LookupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LookupRequest) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

type LookupResponse struct {
	Val []byte `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (m *LookupResponse) Reset()                    { *m = LookupResponse{} }
func (m *LookupResponse) String() string            { return proto1.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()               {}
func (*LookupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LookupResponse) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto1.RegisterType((*LookupRequest)(nil), "proto.LookupRequest")
	proto1.RegisterType((*LookupResponse)(nil), "proto.LookupResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cache service

type CacheClient interface {
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error)
}

type cacheClient struct {
	cc *grpc.ClientConn
}

func NewCacheClient(cc *grpc.ClientConn) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := grpc.Invoke(ctx, "/proto.Cache/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cache service

type CacheServer interface {
	Lookup(context.Context, *LookupRequest) (*LookupResponse, error)
}

func RegisterCacheServer(s *grpc.Server, srv CacheServer) {
	s.RegisterService(&_Cache_serviceDesc, srv)
}

func _Cache_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cache/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lookup",
			Handler:    _Cache_Lookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cache.proto",
}

func init() { proto1.RegisterFile("proto/cache.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4e, 0x4c, 0xce, 0x48, 0xd5, 0x03, 0xb3, 0x85, 0x58, 0xc1, 0x94, 0x92, 0x22,
	0x17, 0xaf, 0x4f, 0x7e, 0x7e, 0x76, 0x69, 0x41, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90,
	0x00, 0x17, 0x73, 0x76, 0x6a, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x88, 0xa9, 0xa4,
	0xc4, 0xc5, 0x07, 0x53, 0x52, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x0a, 0x52, 0x53, 0x96, 0x98, 0x03,
	0x56, 0xc3, 0x13, 0x04, 0x62, 0x1a, 0x39, 0x70, 0xb1, 0x3a, 0x83, 0x0c, 0x17, 0x32, 0xe7, 0x62,
	0x83, 0x28, 0x16, 0x12, 0x81, 0x58, 0xa4, 0x87, 0x62, 0xbc, 0x94, 0x28, 0x9a, 0x28, 0xc4, 0x44,
	0x25, 0x86, 0x24, 0x36, 0xb0, 0xb8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x45, 0x33, 0x34, 0xd8,
	0xab, 0x00, 0x00, 0x00,
}
