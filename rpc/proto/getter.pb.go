// Code generated by protoc-gen-go.
// source: getter.proto
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

// The request message containing the bucket and key
type GetRequest struct {
	Bucket string `protobuf:"bytes,1,opt,name=bucket" json:"bucket,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// The response message containing the item informations
type GetReply struct {
	Key       string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	CreatedAt int64  `protobuf:"varint,2,opt,name=createdAt" json:"createdAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,3,opt,name=updatedAt" json:"updatedAt,omitempty"`
	Data      []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *GetReply) Reset()                    { *m = GetReply{} }
func (m *GetReply) String() string            { return proto1.CompactTextString(m) }
func (*GetReply) ProtoMessage()               {}
func (*GetReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto1.RegisterType((*GetRequest)(nil), "proto.GetRequest")
	proto1.RegisterType((*GetReply)(nil), "proto.GetReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Getter service

type GetterClient interface {
	// Get an item
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
}

type getterClient struct {
	cc *grpc.ClientConn
}

func NewGetterClient(cc *grpc.ClientConn) GetterClient {
	return &getterClient{cc}
}

func (c *getterClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := grpc.Invoke(ctx, "/proto.Getter/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Getter service

type GetterServer interface {
	// Get an item
	Get(context.Context, *GetRequest) (*GetReply, error)
}

func RegisterGetterServer(s *grpc.Server, srv GetterServer) {
	s.RegisterService(&_Getter_serviceDesc, srv)
}

func _Getter_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetterServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Getter/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetterServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Getter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Getter",
	HandlerType: (*GetterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Getter_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto1.RegisterFile("getter.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0x4f, 0x0b, 0x82, 0x40,
	0x10, 0xc5, 0xdb, 0xd6, 0x24, 0x07, 0xa1, 0x9a, 0x43, 0x48, 0x74, 0x10, 0x4f, 0x42, 0xe0, 0xa1,
	0xa8, 0x7b, 0x27, 0xef, 0xfb, 0x0d, 0xfc, 0x33, 0x74, 0x50, 0x70, 0xb3, 0xf1, 0xe0, 0xb7, 0x0f,
	0x47, 0xd3, 0x4e, 0xfb, 0xf6, 0xf7, 0xe3, 0xc1, 0x1b, 0xf0, 0x5f, 0xc4, 0x4c, 0x6d, 0x62, 0xdb,
	0x86, 0x1b, 0xdc, 0xc8, 0x13, 0x3d, 0x00, 0x52, 0x62, 0x43, 0xef, 0x8e, 0x3e, 0x8c, 0x47, 0x70,
	0xf3, 0xae, 0xa8, 0x88, 0x03, 0x15, 0xaa, 0xd8, 0x33, 0xd3, 0x0f, 0xf7, 0xa0, 0x2b, 0xea, 0x83,
	0xb5, 0xc0, 0x21, 0x46, 0x35, 0x6c, 0xa5, 0x67, 0xeb, 0xfe, 0x67, 0xd5, 0x6c, 0xf1, 0x0c, 0x5e,
	0xd1, 0x52, 0xc6, 0x54, 0x3e, 0x59, 0x5a, 0xda, 0x2c, 0x60, 0xb0, 0x9d, 0x2d, 0x27, 0xab, 0x47,
	0x3b, 0x03, 0x44, 0x70, 0xca, 0x8c, 0xb3, 0xc0, 0x09, 0x55, 0xec, 0x1b, 0xc9, 0xd7, 0x3b, 0xb8,
	0xa9, 0x8c, 0xc7, 0x0b, 0xe8, 0x94, 0x18, 0x0f, 0xe3, 0x15, 0xc9, 0xb2, 0xfd, 0xb4, 0xfb, 0x47,
	0xb6, 0xee, 0xa3, 0x55, 0xee, 0x0a, 0xb9, 0x7d, 0x03, 0x00, 0x00, 0xff, 0xff, 0x36, 0x1f, 0xfb,
	0x7e, 0xfa, 0x00, 0x00, 0x00,
}
