// Code generated by protoc-gen-go.
// source: deleter.proto
// DO NOT EDIT!

package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The request message containing the bucket and key
type DeleteRequest struct {
	Bucket string `protobuf:"bytes,1,opt,name=bucket" json:"bucket,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// The response message containing the status
type DeleteReply struct {
	Status int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *DeleteReply) Reset()                    { *m = DeleteReply{} }
func (m *DeleteReply) String() string            { return proto.CompactTextString(m) }
func (*DeleteReply) ProtoMessage()               {}
func (*DeleteReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func init() {
	proto.RegisterType((*DeleteRequest)(nil), "internal.DeleteRequest")
	proto.RegisterType((*DeleteReply)(nil), "internal.DeleteReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Deleter service

type DeleterClient interface {
	// Get an item
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error)
}

type deleterClient struct {
	cc *grpc.ClientConn
}

func NewDeleterClient(cc *grpc.ClientConn) DeleterClient {
	return &deleterClient{cc}
}

func (c *deleterClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error) {
	out := new(DeleteReply)
	err := grpc.Invoke(ctx, "/internal.Deleter/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Deleter service

type DeleterServer interface {
	// Get an item
	Delete(context.Context, *DeleteRequest) (*DeleteReply, error)
}

func RegisterDeleterServer(s *grpc.Server, srv DeleterServer) {
	s.RegisterService(&_Deleter_serviceDesc, srv)
}

func _Deleter_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeleterServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Deleter/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeleterServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Deleter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internal.Deleter",
	HandlerType: (*DeleterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _Deleter_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() { proto.RegisterFile("deleter.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0x51, 0xb2, 0xe4, 0xe2, 0x75, 0x01, 0x4b, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16,
	0x97, 0x08, 0x89, 0x71, 0xb1, 0x25, 0x95, 0x26, 0x67, 0xa7, 0x96, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x41, 0x79, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x4c, 0x60, 0x41, 0x10,
	0x53, 0x49, 0x95, 0x8b, 0x1b, 0xa6, 0xb5, 0x20, 0xa7, 0x12, 0xa4, 0xb1, 0xb8, 0x24, 0xb1, 0xa4,
	0xb4, 0x18, 0xac, 0x91, 0x35, 0x08, 0xca, 0x33, 0x72, 0xe5, 0x62, 0x87, 0x28, 0x2b, 0x12, 0xb2,
	0xe2, 0x62, 0x83, 0x30, 0x85, 0xc4, 0xf5, 0x60, 0x2e, 0xd0, 0x43, 0xb1, 0x5e, 0x4a, 0x14, 0x53,
	0xa2, 0x20, 0xa7, 0x52, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x72, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x62, 0x81, 0xab, 0x71, 0xca, 0x00, 0x00, 0x00,
}