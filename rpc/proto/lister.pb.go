// Code generated by protoc-gen-go.
// source: lister.proto
// DO NOT EDIT!

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

// The request message containing the bucket
type ListRequest struct {
	Bucket string `protobuf:"bytes,1,opt,name=bucket" json:"bucket,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// The response message containing the status
type ListReply struct {
	Items [][]byte `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto1.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func init() {
	proto1.RegisterType((*ListRequest)(nil), "proto.ListRequest")
	proto1.RegisterType((*ListReply)(nil), "proto.ListReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Lister service

type ListerClient interface {
	// List the bucket content
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
}

type listerClient struct {
	cc *grpc.ClientConn
}

func NewListerClient(cc *grpc.ClientConn) ListerClient {
	return &listerClient{cc}
}

func (c *listerClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/proto.Lister/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Lister service

type ListerServer interface {
	// List the bucket content
	List(context.Context, *ListRequest) (*ListReply, error)
}

func RegisterListerServer(s *grpc.Server, srv ListerServer) {
	s.RegisterService(&_Lister_serviceDesc, srv)
}

func _Lister_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Lister/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListerServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lister_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Lister",
	HandlerType: (*ListerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Lister_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() { proto1.RegisterFile("lister.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xc9, 0x2c, 0x2e,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xaa, 0x5c, 0xdc,
	0x3e, 0x99, 0xc5, 0x25, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0x49,
	0xa5, 0xc9, 0xd9, 0xa9, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x92, 0x22,
	0x17, 0x27, 0x44, 0x59, 0x41, 0x4e, 0xa5, 0x90, 0x08, 0x17, 0x6b, 0x66, 0x49, 0x6a, 0x6e, 0xb1,
	0x04, 0xa3, 0x02, 0xb3, 0x06, 0x4f, 0x10, 0x84, 0x63, 0x64, 0xc1, 0xc5, 0xe6, 0x03, 0xb6, 0x40,
	0x48, 0x8f, 0x8b, 0x05, 0xc4, 0x12, 0x12, 0x82, 0x58, 0xa5, 0x87, 0x64, 0x81, 0x94, 0x00, 0x8a,
	0x58, 0x41, 0x4e, 0xa5, 0x12, 0x43, 0x12, 0x1b, 0x58, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xb9, 0x64, 0x8a, 0xb1, 0xa1, 0x00, 0x00, 0x00,
}
