// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0
// source: instance.proto

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

const (
	InstanceService_CreateTrainingImage_FullMethodName   = "/proto.InstanceService/CreateTrainingImage"
	InstanceService_CreateDeploymentImage_FullMethodName = "/proto.InstanceService/CreateDeploymentImage"
	InstanceService_DeleteImage_FullMethodName           = "/proto.InstanceService/DeleteImage"
	InstanceService_CreateInstance_FullMethodName        = "/proto.InstanceService/CreateInstance"
	InstanceService_GetInstanceInfo_FullMethodName       = "/proto.InstanceService/GetInstanceInfo"
	InstanceService_OperateInstance_FullMethodName       = "/proto.InstanceService/OperateInstance"
)

// InstanceServiceClient is the client API for InstanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstanceServiceClient interface {
	CreateTrainingImage(ctx context.Context, in *CreateTrainingImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error)
	CreateDeploymentImage(ctx context.Context, in *CreateDeploymentImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error)
	DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*DeleteImageResponse, error)
	CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error)
	GetInstanceInfo(ctx context.Context, in *InstanceInfoRequest, opts ...grpc.CallOption) (*InstanceInfoResponse, error)
	OperateInstance(ctx context.Context, in *OperateInstanceRequest, opts ...grpc.CallOption) (*OperateInstanceResponse, error)
}

type instanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstanceServiceClient(cc grpc.ClientConnInterface) InstanceServiceClient {
	return &instanceServiceClient{cc}
}

func (c *instanceServiceClient) CreateTrainingImage(ctx context.Context, in *CreateTrainingImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error) {
	out := new(CreateImageResponse)
	err := c.cc.Invoke(ctx, InstanceService_CreateTrainingImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) CreateDeploymentImage(ctx context.Context, in *CreateDeploymentImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error) {
	out := new(CreateImageResponse)
	err := c.cc.Invoke(ctx, InstanceService_CreateDeploymentImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*DeleteImageResponse, error) {
	out := new(DeleteImageResponse)
	err := c.cc.Invoke(ctx, InstanceService_DeleteImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error) {
	out := new(CreateInstanceResponse)
	err := c.cc.Invoke(ctx, InstanceService_CreateInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) GetInstanceInfo(ctx context.Context, in *InstanceInfoRequest, opts ...grpc.CallOption) (*InstanceInfoResponse, error) {
	out := new(InstanceInfoResponse)
	err := c.cc.Invoke(ctx, InstanceService_GetInstanceInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) OperateInstance(ctx context.Context, in *OperateInstanceRequest, opts ...grpc.CallOption) (*OperateInstanceResponse, error) {
	out := new(OperateInstanceResponse)
	err := c.cc.Invoke(ctx, InstanceService_OperateInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstanceServiceServer is the server API for InstanceService service.
// All implementations must embed UnimplementedInstanceServiceServer
// for forward compatibility
type InstanceServiceServer interface {
	CreateTrainingImage(context.Context, *CreateTrainingImageRequest) (*CreateImageResponse, error)
	CreateDeploymentImage(context.Context, *CreateDeploymentImageRequest) (*CreateImageResponse, error)
	DeleteImage(context.Context, *DeleteImageRequest) (*DeleteImageResponse, error)
	CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error)
	GetInstanceInfo(context.Context, *InstanceInfoRequest) (*InstanceInfoResponse, error)
	OperateInstance(context.Context, *OperateInstanceRequest) (*OperateInstanceResponse, error)
	mustEmbedUnimplementedInstanceServiceServer()
}

// UnimplementedInstanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInstanceServiceServer struct {
}

func (UnimplementedInstanceServiceServer) CreateTrainingImage(context.Context, *CreateTrainingImageRequest) (*CreateImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrainingImage not implemented")
}
func (UnimplementedInstanceServiceServer) CreateDeploymentImage(context.Context, *CreateDeploymentImageRequest) (*CreateImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeploymentImage not implemented")
}
func (UnimplementedInstanceServiceServer) DeleteImage(context.Context, *DeleteImageRequest) (*DeleteImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImage not implemented")
}
func (UnimplementedInstanceServiceServer) CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInstance not implemented")
}
func (UnimplementedInstanceServiceServer) GetInstanceInfo(context.Context, *InstanceInfoRequest) (*InstanceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstanceInfo not implemented")
}
func (UnimplementedInstanceServiceServer) OperateInstance(context.Context, *OperateInstanceRequest) (*OperateInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OperateInstance not implemented")
}
func (UnimplementedInstanceServiceServer) mustEmbedUnimplementedInstanceServiceServer() {}

// UnsafeInstanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstanceServiceServer will
// result in compilation errors.
type UnsafeInstanceServiceServer interface {
	mustEmbedUnimplementedInstanceServiceServer()
}

func RegisterInstanceServiceServer(s grpc.ServiceRegistrar, srv InstanceServiceServer) {
	s.RegisterService(&InstanceService_ServiceDesc, srv)
}

func _InstanceService_CreateTrainingImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTrainingImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).CreateTrainingImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InstanceService_CreateTrainingImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).CreateTrainingImage(ctx, req.(*CreateTrainingImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_CreateDeploymentImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeploymentImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).CreateDeploymentImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InstanceService_CreateDeploymentImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).CreateDeploymentImage(ctx, req.(*CreateDeploymentImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_DeleteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).DeleteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InstanceService_DeleteImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).DeleteImage(ctx, req.(*DeleteImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_CreateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).CreateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InstanceService_CreateInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).CreateInstance(ctx, req.(*CreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_GetInstanceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstanceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).GetInstanceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InstanceService_GetInstanceInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).GetInstanceInfo(ctx, req.(*InstanceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_OperateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).OperateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InstanceService_OperateInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).OperateInstance(ctx, req.(*OperateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InstanceService_ServiceDesc is the grpc.ServiceDesc for InstanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InstanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.InstanceService",
	HandlerType: (*InstanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTrainingImage",
			Handler:    _InstanceService_CreateTrainingImage_Handler,
		},
		{
			MethodName: "CreateDeploymentImage",
			Handler:    _InstanceService_CreateDeploymentImage_Handler,
		},
		{
			MethodName: "DeleteImage",
			Handler:    _InstanceService_DeleteImage_Handler,
		},
		{
			MethodName: "CreateInstance",
			Handler:    _InstanceService_CreateInstance_Handler,
		},
		{
			MethodName: "GetInstanceInfo",
			Handler:    _InstanceService_GetInstanceInfo_Handler,
		},
		{
			MethodName: "OperateInstance",
			Handler:    _InstanceService_OperateInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "instance.proto",
}
