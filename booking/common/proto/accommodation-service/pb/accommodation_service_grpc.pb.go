// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.4
// source: accommodation_service.proto

package accommodation

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

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*AccommodationsResponse, error)
	GetAccommodationById(ctx context.Context, in *GetAccommodationByIdRequest, opts ...grpc.CallOption) (*AccommodationResponse, error)
	GetAccommodationsByHostId(ctx context.Context, in *GetAccommodationsByHostIdRequest, opts ...grpc.CallOption) (*AccommodationsResponse, error)
	DeleteAccommodation(ctx context.Context, in *DeleteAccommodationRequest, opts ...grpc.CallOption) (*AccommodationResponse, error)
	DeleteAccommodationsByHostId(ctx context.Context, in *DeleteAccommodationsByHostIdRequest, opts ...grpc.CallOption) (*AccommodationsResponse, error)
	CreateAccommodation(ctx context.Context, in *CreateAccommodationRequest, opts ...grpc.CallOption) (*AccommodationResponse, error)
	CreateAvailability(ctx context.Context, in *CreateAvailabilityRequest, opts ...grpc.CallOption) (*AccommodationResponse, error)
	Search(ctx context.Context, in *GetAccommodationsByParamsRequest, opts ...grpc.CallOption) (*AccommodationsResponse, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*AccommodationsResponse, error) {
	out := new(AccommodationsResponse)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAccommodationById(ctx context.Context, in *GetAccommodationByIdRequest, opts ...grpc.CallOption) (*AccommodationResponse, error) {
	out := new(AccommodationResponse)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/GetAccommodationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAccommodationsByHostId(ctx context.Context, in *GetAccommodationsByHostIdRequest, opts ...grpc.CallOption) (*AccommodationsResponse, error) {
	out := new(AccommodationsResponse)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/GetAccommodationsByHostId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) DeleteAccommodation(ctx context.Context, in *DeleteAccommodationRequest, opts ...grpc.CallOption) (*AccommodationResponse, error) {
	out := new(AccommodationResponse)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/DeleteAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) DeleteAccommodationsByHostId(ctx context.Context, in *DeleteAccommodationsByHostIdRequest, opts ...grpc.CallOption) (*AccommodationsResponse, error) {
	out := new(AccommodationsResponse)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/DeleteAccommodationsByHostId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) CreateAccommodation(ctx context.Context, in *CreateAccommodationRequest, opts ...grpc.CallOption) (*AccommodationResponse, error) {
	out := new(AccommodationResponse)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/CreateAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) CreateAvailability(ctx context.Context, in *CreateAvailabilityRequest, opts ...grpc.CallOption) (*AccommodationResponse, error) {
	out := new(AccommodationResponse)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/CreateAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) Search(ctx context.Context, in *GetAccommodationsByParamsRequest, opts ...grpc.CallOption) (*AccommodationsResponse, error) {
	out := new(AccommodationsResponse)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	GetAll(context.Context, *GetAllRequest) (*AccommodationsResponse, error)
	GetAccommodationById(context.Context, *GetAccommodationByIdRequest) (*AccommodationResponse, error)
	GetAccommodationsByHostId(context.Context, *GetAccommodationsByHostIdRequest) (*AccommodationsResponse, error)
	DeleteAccommodation(context.Context, *DeleteAccommodationRequest) (*AccommodationResponse, error)
	DeleteAccommodationsByHostId(context.Context, *DeleteAccommodationsByHostIdRequest) (*AccommodationsResponse, error)
	CreateAccommodation(context.Context, *CreateAccommodationRequest) (*AccommodationResponse, error)
	CreateAvailability(context.Context, *CreateAvailabilityRequest) (*AccommodationResponse, error)
	Search(context.Context, *GetAccommodationsByParamsRequest) (*AccommodationsResponse, error)
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (UnimplementedAccommodationServiceServer) GetAll(context.Context, *GetAllRequest) (*AccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAccommodationById(context.Context, *GetAccommodationByIdRequest) (*AccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationById not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAccommodationsByHostId(context.Context, *GetAccommodationsByHostIdRequest) (*AccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationsByHostId not implemented")
}
func (UnimplementedAccommodationServiceServer) DeleteAccommodation(context.Context, *DeleteAccommodationRequest) (*AccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) DeleteAccommodationsByHostId(context.Context, *DeleteAccommodationsByHostIdRequest) (*AccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccommodationsByHostId not implemented")
}
func (UnimplementedAccommodationServiceServer) CreateAccommodation(context.Context, *CreateAccommodationRequest) (*AccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) CreateAvailability(context.Context, *CreateAvailabilityRequest) (*AccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAvailability not implemented")
}
func (UnimplementedAccommodationServiceServer) Search(context.Context, *GetAccommodationsByParamsRequest) (*AccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedAccommodationServiceServer) mustEmbedUnimplementedAccommodationServiceServer() {}

// UnsafeAccommodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationServiceServer will
// result in compilation errors.
type UnsafeAccommodationServiceServer interface {
	mustEmbedUnimplementedAccommodationServiceServer()
}

func RegisterAccommodationServiceServer(s grpc.ServiceRegistrar, srv AccommodationServiceServer) {
	s.RegisterService(&AccommodationService_ServiceDesc, srv)
}

func _AccommodationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAccommodationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccommodationByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAccommodationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/GetAccommodationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAccommodationById(ctx, req.(*GetAccommodationByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAccommodationsByHostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccommodationsByHostIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAccommodationsByHostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/GetAccommodationsByHostId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAccommodationsByHostId(ctx, req.(*GetAccommodationsByHostIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_DeleteAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).DeleteAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/DeleteAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).DeleteAccommodation(ctx, req.(*DeleteAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_DeleteAccommodationsByHostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccommodationsByHostIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).DeleteAccommodationsByHostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/DeleteAccommodationsByHostId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).DeleteAccommodationsByHostId(ctx, req.(*DeleteAccommodationsByHostIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_CreateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/CreateAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateAccommodation(ctx, req.(*CreateAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_CreateAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/CreateAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateAvailability(ctx, req.(*CreateAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccommodationsByParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Search(ctx, req.(*GetAccommodationsByParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationService_ServiceDesc is the grpc.ServiceDesc for AccommodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accommodation.AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _AccommodationService_GetAll_Handler,
		},
		{
			MethodName: "GetAccommodationById",
			Handler:    _AccommodationService_GetAccommodationById_Handler,
		},
		{
			MethodName: "GetAccommodationsByHostId",
			Handler:    _AccommodationService_GetAccommodationsByHostId_Handler,
		},
		{
			MethodName: "DeleteAccommodation",
			Handler:    _AccommodationService_DeleteAccommodation_Handler,
		},
		{
			MethodName: "DeleteAccommodationsByHostId",
			Handler:    _AccommodationService_DeleteAccommodationsByHostId_Handler,
		},
		{
			MethodName: "CreateAccommodation",
			Handler:    _AccommodationService_CreateAccommodation_Handler,
		},
		{
			MethodName: "CreateAvailability",
			Handler:    _AccommodationService_CreateAvailability_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _AccommodationService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accommodation_service.proto",
}
