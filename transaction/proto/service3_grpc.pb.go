// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceThreeClient is the client API for ServiceThree service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceThreeClient interface {
	UpdateEmail(ctx context.Context, in *UpdateEmailRequest, opts ...grpc.CallOption) (*UpdateEmailResponse, error)
	RollbackThree(ctx context.Context, in *RollbackThreeRequest, opts ...grpc.CallOption) (*RollbackThreeResponse, error)
}

type serviceThreeClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceThreeClient(cc grpc.ClientConnInterface) ServiceThreeClient {
	return &serviceThreeClient{cc}
}

func (c *serviceThreeClient) UpdateEmail(ctx context.Context, in *UpdateEmailRequest, opts ...grpc.CallOption) (*UpdateEmailResponse, error) {
	out := new(UpdateEmailResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceThree/UpdateEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceThreeClient) RollbackThree(ctx context.Context, in *RollbackThreeRequest, opts ...grpc.CallOption) (*RollbackThreeResponse, error) {
	out := new(RollbackThreeResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceThree/RollbackThree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceThreeServer is the server API for ServiceThree service.
// All implementations must embed UnimplementedServiceThreeServer
// for forward compatibility
type ServiceThreeServer interface {
	UpdateEmail(context.Context, *UpdateEmailRequest) (*UpdateEmailResponse, error)
	RollbackThree(context.Context, *RollbackThreeRequest) (*RollbackThreeResponse, error)
	mustEmbedUnimplementedServiceThreeServer()
}

// UnimplementedServiceThreeServer must be embedded to have forward compatible implementations.
type UnimplementedServiceThreeServer struct {
}

func (UnimplementedServiceThreeServer) UpdateEmail(context.Context, *UpdateEmailRequest) (*UpdateEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmail not implemented")
}
func (UnimplementedServiceThreeServer) RollbackThree(context.Context, *RollbackThreeRequest) (*RollbackThreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackThree not implemented")
}
func (UnimplementedServiceThreeServer) mustEmbedUnimplementedServiceThreeServer() {}

// UnsafeServiceThreeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceThreeServer will
// result in compilation errors.
type UnsafeServiceThreeServer interface {
	mustEmbedUnimplementedServiceThreeServer()
}

func RegisterServiceThreeServer(s grpc.ServiceRegistrar, srv ServiceThreeServer) {
	s.RegisterService(&ServiceThree_ServiceDesc, srv)
}

func _ServiceThree_UpdateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceThreeServer).UpdateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceThree/UpdateEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceThreeServer).UpdateEmail(ctx, req.(*UpdateEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceThree_RollbackThree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackThreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceThreeServer).RollbackThree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceThree/RollbackThree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceThreeServer).RollbackThree(ctx, req.(*RollbackThreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceThree_ServiceDesc is the grpc.ServiceDesc for ServiceThree service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceThree_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServiceThree",
	HandlerType: (*ServiceThreeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateEmail",
			Handler:    _ServiceThree_UpdateEmail_Handler,
		},
		{
			MethodName: "RollbackThree",
			Handler:    _ServiceThree_RollbackThree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service3.proto",
}