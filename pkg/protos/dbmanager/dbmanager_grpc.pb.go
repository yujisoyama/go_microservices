// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dbmanager

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DbManagerClient is the client API for DbManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DbManagerClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type dbManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewDbManagerClient(cc grpc.ClientConnInterface) DbManagerClient {
	return &dbManagerClient{cc}
}

func (c *dbManagerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/dbmanager.DbManager/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbManagerClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/dbmanager.DbManager/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DbManagerServer is the server API for DbManager service.
// All implementations must embed UnimplementedDbManagerServer
// for forward compatibility
type DbManagerServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	mustEmbedUnimplementedDbManagerServer()
}

// UnimplementedDbManagerServer must be embedded to have forward compatible implementations.
type UnimplementedDbManagerServer struct {
}

func (UnimplementedDbManagerServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDbManagerServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedDbManagerServer) mustEmbedUnimplementedDbManagerServer() {}

// UnsafeDbManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DbManagerServer will
// result in compilation errors.
type UnsafeDbManagerServer interface {
	mustEmbedUnimplementedDbManagerServer()
}

func RegisterDbManagerServer(s *grpc.Server, srv DbManagerServer) {
	s.RegisterService(&_DbManager_serviceDesc, srv)
}

func _DbManager_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbManagerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbmanager.DbManager/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbManagerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbManager_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbManagerServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbmanager.DbManager/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbManagerServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DbManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dbmanager.DbManager",
	HandlerType: (*DbManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _DbManager_Ping_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _DbManager_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protos/dbmanager/dbmanager.proto",
}
