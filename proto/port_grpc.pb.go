// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	Load(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Storage_LoadClient, error)
	Save(ctx context.Context, opts ...grpc.CallOption) (Storage_SaveClient, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Load(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Storage_LoadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Storage_serviceDesc.Streams[0], "/proto.Storage/Load", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageLoadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Storage_LoadClient interface {
	Recv() (*Payload, error)
	grpc.ClientStream
}

type storageLoadClient struct {
	grpc.ClientStream
}

func (x *storageLoadClient) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) Save(ctx context.Context, opts ...grpc.CallOption) (Storage_SaveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Storage_serviceDesc.Streams[1], "/proto.Storage/Save", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageSaveClient{stream}
	return x, nil
}

type Storage_SaveClient interface {
	Send(*Payload) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type storageSaveClient struct {
	grpc.ClientStream
}

func (x *storageSaveClient) Send(m *Payload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageSaveClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	Load(*empty.Empty, Storage_LoadServer) error
	Save(Storage_SaveServer) error
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) Load(*empty.Empty, Storage_LoadServer) error {
	return status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (UnimplementedStorageServer) Save(Storage_SaveServer) error {
	return status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_Load_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServer).Load(m, &storageLoadServer{stream})
}

type Storage_LoadServer interface {
	Send(*Payload) error
	grpc.ServerStream
}

type storageLoadServer struct {
	grpc.ServerStream
}

func (x *storageLoadServer) Send(m *Payload) error {
	return x.ServerStream.SendMsg(m)
}

func _Storage_Save_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServer).Save(&storageSaveServer{stream})
}

type Storage_SaveServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*Payload, error)
	grpc.ServerStream
}

type storageSaveServer struct {
	grpc.ServerStream
}

func (x *storageSaveServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageSaveServer) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Load",
			Handler:       _Storage_Load_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Save",
			Handler:       _Storage_Save_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/port.proto",
}
