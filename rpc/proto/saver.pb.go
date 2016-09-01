// Code generated by protoc-gen-go.
// source: saver.proto
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

// The request message containing the user's data.
type SaveRequest struct {
	Bucket string `protobuf:"bytes,1,opt,name=bucket" json:"bucket,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SaveRequest) Reset()                    { *m = SaveRequest{} }
func (m *SaveRequest) String() string            { return proto1.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()               {}
func (*SaveRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// The response message containing the status
type SaveReply struct {
	Status int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *SaveReply) Reset()                    { *m = SaveReply{} }
func (m *SaveReply) String() string            { return proto1.CompactTextString(m) }
func (*SaveReply) ProtoMessage()               {}
func (*SaveReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto1.RegisterType((*SaveRequest)(nil), "proto.SaveRequest")
	proto1.RegisterType((*SaveReply)(nil), "proto.SaveReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Saver service

type SaverClient interface {
	// Saves user data
	Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveReply, error)
}

type saverClient struct {
	cc *grpc.ClientConn
}

func NewSaverClient(cc *grpc.ClientConn) SaverClient {
	return &saverClient{cc}
}

func (c *saverClient) Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveReply, error) {
	out := new(SaveReply)
	err := grpc.Invoke(ctx, "/proto.Saver/Save", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Saver service

type SaverServer interface {
	// Saves user data
	Save(context.Context, *SaveRequest) (*SaveReply, error)
}

func RegisterSaverServer(s *grpc.Server, srv SaverServer) {
	s.RegisterService(&_Saver_serviceDesc, srv)
}

func _Saver_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaverServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Saver/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaverServer).Save(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Saver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Saver",
	HandlerType: (*SaverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _Saver_Save_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto1.RegisterFile("saver.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4e, 0x2c, 0x4b,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xde, 0x5c, 0xdc, 0xc1,
	0x89, 0x65, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0x49, 0xa5,
	0xc9, 0xd9, 0xa9, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x90, 0x00, 0x17,
	0x73, 0x76, 0x6a, 0xa5, 0x04, 0x13, 0x58, 0x10, 0xc4, 0x14, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c,
	0x49, 0x94, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x95, 0x94, 0xb9, 0x38, 0x21, 0x86,
	0x15, 0xe4, 0x54, 0x82, 0x8c, 0x2a, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x06, 0x1b, 0xc5, 0x1a, 0x04,
	0xe5, 0x19, 0x99, 0x73, 0xb1, 0x82, 0x14, 0x15, 0x09, 0xe9, 0x71, 0xb1, 0x80, 0x18, 0x42, 0x42,
	0x10, 0x17, 0xe9, 0x21, 0xb9, 0x43, 0x4a, 0x00, 0x45, 0xac, 0x20, 0xa7, 0x52, 0x89, 0x21, 0x89,
	0x0d, 0x2c, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x61, 0x82, 0x87, 0xc7, 0x00, 0x00,
	0x00,
}
