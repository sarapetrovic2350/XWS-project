// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.4
// source: rating_service.proto

package rating

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

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	GetAllRatingsHost(ctx context.Context, in *GetAllRatingsHostRequest, opts ...grpc.CallOption) (*GetAllRatingsHostResponse, error)
	GetAllRatingsAccommodation(ctx context.Context, in *GetAllRatingsAccommodationRequest, opts ...grpc.CallOption) (*GetAllRatingsAccommodationResponse, error)
	CreateRatingForHost(ctx context.Context, in *CreateRatingForHostRequest, opts ...grpc.CallOption) (*CreateRatingForHostResponse, error)
	CreateRatingForAccommodation(ctx context.Context, in *CreateRatingForAccommodationRequest, opts ...grpc.CallOption) (*CreateRatingForAccommodationResponse, error)
	DeleteRatingForHost(ctx context.Context, in *DeleteRatingForHostRequest, opts ...grpc.CallOption) (*DeleteRatingForHostResponse, error)
	UpdateRatingForHost(ctx context.Context, in *UpdateRatingForHostRequest, opts ...grpc.CallOption) (*UpdateRatingForHostResponse, error)
	DeleteRatingForAccommodation(ctx context.Context, in *DeleteRatingForAccommodationRequest, opts ...grpc.CallOption) (*DeleteRatingForAccommodationResponse, error)
	UpdateRatingForAccommodation(ctx context.Context, in *UpdateRatingForAccommodationRequest, opts ...grpc.CallOption) (*UpdateRatingForAccommodationResponse, error)
	GetAvgRatingForHost(ctx context.Context, in *GetAvgRatingForHostRequest, opts ...grpc.CallOption) (*GetAvgRatingForHostResponse, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) GetAllRatingsHost(ctx context.Context, in *GetAllRatingsHostRequest, opts ...grpc.CallOption) (*GetAllRatingsHostResponse, error) {
	out := new(GetAllRatingsHostResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetAllRatingsHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetAllRatingsAccommodation(ctx context.Context, in *GetAllRatingsAccommodationRequest, opts ...grpc.CallOption) (*GetAllRatingsAccommodationResponse, error) {
	out := new(GetAllRatingsAccommodationResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetAllRatingsAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) CreateRatingForHost(ctx context.Context, in *CreateRatingForHostRequest, opts ...grpc.CallOption) (*CreateRatingForHostResponse, error) {
	out := new(CreateRatingForHostResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/CreateRatingForHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) CreateRatingForAccommodation(ctx context.Context, in *CreateRatingForAccommodationRequest, opts ...grpc.CallOption) (*CreateRatingForAccommodationResponse, error) {
	out := new(CreateRatingForAccommodationResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/CreateRatingForAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) DeleteRatingForHost(ctx context.Context, in *DeleteRatingForHostRequest, opts ...grpc.CallOption) (*DeleteRatingForHostResponse, error) {
	out := new(DeleteRatingForHostResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/DeleteRatingForHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) UpdateRatingForHost(ctx context.Context, in *UpdateRatingForHostRequest, opts ...grpc.CallOption) (*UpdateRatingForHostResponse, error) {
	out := new(UpdateRatingForHostResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/UpdateRatingForHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) DeleteRatingForAccommodation(ctx context.Context, in *DeleteRatingForAccommodationRequest, opts ...grpc.CallOption) (*DeleteRatingForAccommodationResponse, error) {
	out := new(DeleteRatingForAccommodationResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/DeleteRatingForAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) UpdateRatingForAccommodation(ctx context.Context, in *UpdateRatingForAccommodationRequest, opts ...grpc.CallOption) (*UpdateRatingForAccommodationResponse, error) {
	out := new(UpdateRatingForAccommodationResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/UpdateRatingForAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetAvgRatingForHost(ctx context.Context, in *GetAvgRatingForHostRequest, opts ...grpc.CallOption) (*GetAvgRatingForHostResponse, error) {
	out := new(GetAvgRatingForHostResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetAvgRatingForHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	GetAllRatingsHost(context.Context, *GetAllRatingsHostRequest) (*GetAllRatingsHostResponse, error)
	GetAllRatingsAccommodation(context.Context, *GetAllRatingsAccommodationRequest) (*GetAllRatingsAccommodationResponse, error)
	CreateRatingForHost(context.Context, *CreateRatingForHostRequest) (*CreateRatingForHostResponse, error)
	CreateRatingForAccommodation(context.Context, *CreateRatingForAccommodationRequest) (*CreateRatingForAccommodationResponse, error)
	DeleteRatingForHost(context.Context, *DeleteRatingForHostRequest) (*DeleteRatingForHostResponse, error)
	UpdateRatingForHost(context.Context, *UpdateRatingForHostRequest) (*UpdateRatingForHostResponse, error)
	DeleteRatingForAccommodation(context.Context, *DeleteRatingForAccommodationRequest) (*DeleteRatingForAccommodationResponse, error)
	UpdateRatingForAccommodation(context.Context, *UpdateRatingForAccommodationRequest) (*UpdateRatingForAccommodationResponse, error)
	GetAvgRatingForHost(context.Context, *GetAvgRatingForHostRequest) (*GetAvgRatingForHostResponse, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) GetAllRatingsHost(context.Context, *GetAllRatingsHostRequest) (*GetAllRatingsHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRatingsHost not implemented")
}
func (UnimplementedRatingServiceServer) GetAllRatingsAccommodation(context.Context, *GetAllRatingsAccommodationRequest) (*GetAllRatingsAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRatingsAccommodation not implemented")
}
func (UnimplementedRatingServiceServer) CreateRatingForHost(context.Context, *CreateRatingForHostRequest) (*CreateRatingForHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRatingForHost not implemented")
}
func (UnimplementedRatingServiceServer) CreateRatingForAccommodation(context.Context, *CreateRatingForAccommodationRequest) (*CreateRatingForAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRatingForAccommodation not implemented")
}
func (UnimplementedRatingServiceServer) DeleteRatingForHost(context.Context, *DeleteRatingForHostRequest) (*DeleteRatingForHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRatingForHost not implemented")
}
func (UnimplementedRatingServiceServer) UpdateRatingForHost(context.Context, *UpdateRatingForHostRequest) (*UpdateRatingForHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRatingForHost not implemented")
}
func (UnimplementedRatingServiceServer) DeleteRatingForAccommodation(context.Context, *DeleteRatingForAccommodationRequest) (*DeleteRatingForAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRatingForAccommodation not implemented")
}
func (UnimplementedRatingServiceServer) UpdateRatingForAccommodation(context.Context, *UpdateRatingForAccommodationRequest) (*UpdateRatingForAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRatingForAccommodation not implemented")
}
func (UnimplementedRatingServiceServer) GetAvgRatingForHost(context.Context, *GetAvgRatingForHostRequest) (*GetAvgRatingForHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvgRatingForHost not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_GetAllRatingsHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRatingsHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAllRatingsHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetAllRatingsHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAllRatingsHost(ctx, req.(*GetAllRatingsHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetAllRatingsAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRatingsAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAllRatingsAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetAllRatingsAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAllRatingsAccommodation(ctx, req.(*GetAllRatingsAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_CreateRatingForHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRatingForHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).CreateRatingForHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/CreateRatingForHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).CreateRatingForHost(ctx, req.(*CreateRatingForHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_CreateRatingForAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRatingForAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).CreateRatingForAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/CreateRatingForAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).CreateRatingForAccommodation(ctx, req.(*CreateRatingForAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_DeleteRatingForHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRatingForHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).DeleteRatingForHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/DeleteRatingForHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).DeleteRatingForHost(ctx, req.(*DeleteRatingForHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_UpdateRatingForHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRatingForHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).UpdateRatingForHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/UpdateRatingForHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).UpdateRatingForHost(ctx, req.(*UpdateRatingForHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_DeleteRatingForAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRatingForAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).DeleteRatingForAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/DeleteRatingForAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).DeleteRatingForAccommodation(ctx, req.(*DeleteRatingForAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_UpdateRatingForAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRatingForAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).UpdateRatingForAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/UpdateRatingForAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).UpdateRatingForAccommodation(ctx, req.(*UpdateRatingForAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetAvgRatingForHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvgRatingForHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAvgRatingForHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetAvgRatingForHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAvgRatingForHost(ctx, req.(*GetAvgRatingForHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rating.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllRatingsHost",
			Handler:    _RatingService_GetAllRatingsHost_Handler,
		},
		{
			MethodName: "GetAllRatingsAccommodation",
			Handler:    _RatingService_GetAllRatingsAccommodation_Handler,
		},
		{
			MethodName: "CreateRatingForHost",
			Handler:    _RatingService_CreateRatingForHost_Handler,
		},
		{
			MethodName: "CreateRatingForAccommodation",
			Handler:    _RatingService_CreateRatingForAccommodation_Handler,
		},
		{
			MethodName: "DeleteRatingForHost",
			Handler:    _RatingService_DeleteRatingForHost_Handler,
		},
		{
			MethodName: "UpdateRatingForHost",
			Handler:    _RatingService_UpdateRatingForHost_Handler,
		},
		{
			MethodName: "DeleteRatingForAccommodation",
			Handler:    _RatingService_DeleteRatingForAccommodation_Handler,
		},
		{
			MethodName: "UpdateRatingForAccommodation",
			Handler:    _RatingService_UpdateRatingForAccommodation_Handler,
		},
		{
			MethodName: "GetAvgRatingForHost",
			Handler:    _RatingService_GetAvgRatingForHost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rating_service.proto",
}
