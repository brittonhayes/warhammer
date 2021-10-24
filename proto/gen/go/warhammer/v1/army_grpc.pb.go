// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package warhammerv1

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

// ArmyServiceClient is the client API for ArmyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArmyServiceClient interface {
	CreateUnit(ctx context.Context, in *CreateUnitRequest, opts ...grpc.CallOption) (*CreateUnitResponse, error)
}

type armyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArmyServiceClient(cc grpc.ClientConnInterface) ArmyServiceClient {
	return &armyServiceClient{cc}
}

func (c *armyServiceClient) CreateUnit(ctx context.Context, in *CreateUnitRequest, opts ...grpc.CallOption) (*CreateUnitResponse, error) {
	out := new(CreateUnitResponse)
	err := c.cc.Invoke(ctx, "/warhammer.v1.ArmyService/CreateUnit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArmyServiceServer is the server API for ArmyService service.
// All implementations should embed UnimplementedArmyServiceServer
// for forward compatibility
type ArmyServiceServer interface {
	CreateUnit(context.Context, *CreateUnitRequest) (*CreateUnitResponse, error)
}

// UnimplementedArmyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedArmyServiceServer struct {
}

func (UnimplementedArmyServiceServer) CreateUnit(context.Context, *CreateUnitRequest) (*CreateUnitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUnit not implemented")
}

// UnsafeArmyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArmyServiceServer will
// result in compilation errors.
type UnsafeArmyServiceServer interface {
	mustEmbedUnimplementedArmyServiceServer()
}

func RegisterArmyServiceServer(s grpc.ServiceRegistrar, srv ArmyServiceServer) {
	s.RegisterService(&ArmyService_ServiceDesc, srv)
}

func _ArmyService_CreateUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUnitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArmyServiceServer).CreateUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warhammer.v1.ArmyService/CreateUnit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArmyServiceServer).CreateUnit(ctx, req.(*CreateUnitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArmyService_ServiceDesc is the grpc.ServiceDesc for ArmyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArmyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "warhammer.v1.ArmyService",
	HandlerType: (*ArmyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUnit",
			Handler:    _ArmyService_CreateUnit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warhammer/v1/army.proto",
}
