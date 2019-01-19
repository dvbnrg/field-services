// Code generated by protoc-gen-go. DO NOT EDIT.
// source: purchasing.proto

package server

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

func init() { proto.RegisterFile("purchasing.proto", fileDescriptor_f3d74dfbb927d1c0) }

var fileDescriptor_f3d74dfbb927d1c0 = []byte{
	// 61 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0x2d, 0x4a,
	0xce, 0x48, 0x2c, 0xce, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x4e,
	0x2d, 0x2a, 0x4b, 0x2d, 0x32, 0xe2, 0xe1, 0xe2, 0x42, 0xc8, 0x25, 0xb1, 0x81, 0x25, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xee, 0x9c, 0xb7, 0xf7, 0x30, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PurchasingClient is the client API for Purchasing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PurchasingClient interface {
}

type purchasingClient struct {
	cc *grpc.ClientConn
}

func NewPurchasingClient(cc *grpc.ClientConn) PurchasingClient {
	return &purchasingClient{cc}
}

// PurchasingServer is the server API for Purchasing service.
type PurchasingServer interface {
}

func RegisterPurchasingServer(s *grpc.Server, srv PurchasingServer) {
	s.RegisterService(&_Purchasing_serviceDesc, srv)
}

var _Purchasing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.purchasing",
	HandlerType: (*PurchasingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "purchasing.proto",
}
