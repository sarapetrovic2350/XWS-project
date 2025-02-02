// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.4
// source: reservation_service.proto

package reservation

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

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetActiveReservationsByGuestId(ctx context.Context, in *GetActiveReservationsRequest, opts ...grpc.CallOption) (*GetActiveReservationsResponse, error)
	GetReservationsByUserId(ctx context.Context, in *GetReservationsByUserIdRequest, opts ...grpc.CallOption) (*GetReservationsByUserIdResponse, error)
	GetReservationsByAccommodationId(ctx context.Context, in *GetReservationsByAccommodationRequest, opts ...grpc.CallOption) (*GetReservationsByAccommodationResponse, error)
	GetActiveReservationsByHostId(ctx context.Context, in *GetActiveReservationsRequest, opts ...grpc.CallOption) (*GetActiveReservationsResponse, error)
	GetPendingReservationsForHost(ctx context.Context, in *GetPendingReservationsForHostRequest, opts ...grpc.CallOption) (*GetPendingReservationsForHostResponse, error)
	GetReservationsForHost(ctx context.Context, in *GetReservationsForHostRequest, opts ...grpc.CallOption) (*GetReservationsForHostResponse, error)
	GetNumberOfPastReservationsByHostId(ctx context.Context, in *GetNumberOfPastReservationsByHostRequest, opts ...grpc.CallOption) (*GetNumberOfPastReservationsByHostResponse, error)
	GetDurationOfPastReservationsByHostId(ctx context.Context, in *GetDurationOfPastReservationsByHostIdRequest, opts ...grpc.CallOption) (*GetDurationOfPastReservationsByHostIdResponse, error)
	GetAcceptanceRateByHostId(ctx context.Context, in *GetAcceptanceRateByHostIdRequest, opts ...grpc.CallOption) (*GetAcceptanceRateByHostIdResponse, error)
	RejectPendingReservationByHost(ctx context.Context, in *RejectPendingReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error)
	AcceptPendingReservationByHost(ctx context.Context, in *AcceptPendingReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error)
	DeletePendingReservationByGuest(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error)
	CancelReservationByGuest(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*CancelReservationResponse, error)
	CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetActiveReservationsByGuestId(ctx context.Context, in *GetActiveReservationsRequest, opts ...grpc.CallOption) (*GetActiveReservationsResponse, error) {
	out := new(GetActiveReservationsResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetActiveReservationsByGuestId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetReservationsByUserId(ctx context.Context, in *GetReservationsByUserIdRequest, opts ...grpc.CallOption) (*GetReservationsByUserIdResponse, error) {
	out := new(GetReservationsByUserIdResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetReservationsByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetReservationsByAccommodationId(ctx context.Context, in *GetReservationsByAccommodationRequest, opts ...grpc.CallOption) (*GetReservationsByAccommodationResponse, error) {
	out := new(GetReservationsByAccommodationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetReservationsByAccommodationId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetActiveReservationsByHostId(ctx context.Context, in *GetActiveReservationsRequest, opts ...grpc.CallOption) (*GetActiveReservationsResponse, error) {
	out := new(GetActiveReservationsResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetActiveReservationsByHostId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetPendingReservationsForHost(ctx context.Context, in *GetPendingReservationsForHostRequest, opts ...grpc.CallOption) (*GetPendingReservationsForHostResponse, error) {
	out := new(GetPendingReservationsForHostResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetPendingReservationsForHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetReservationsForHost(ctx context.Context, in *GetReservationsForHostRequest, opts ...grpc.CallOption) (*GetReservationsForHostResponse, error) {
	out := new(GetReservationsForHostResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetReservationsForHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetNumberOfPastReservationsByHostId(ctx context.Context, in *GetNumberOfPastReservationsByHostRequest, opts ...grpc.CallOption) (*GetNumberOfPastReservationsByHostResponse, error) {
	out := new(GetNumberOfPastReservationsByHostResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetNumberOfPastReservationsByHostId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetDurationOfPastReservationsByHostId(ctx context.Context, in *GetDurationOfPastReservationsByHostIdRequest, opts ...grpc.CallOption) (*GetDurationOfPastReservationsByHostIdResponse, error) {
	out := new(GetDurationOfPastReservationsByHostIdResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetDurationOfPastReservationsByHostId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAcceptanceRateByHostId(ctx context.Context, in *GetAcceptanceRateByHostIdRequest, opts ...grpc.CallOption) (*GetAcceptanceRateByHostIdResponse, error) {
	out := new(GetAcceptanceRateByHostIdResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetAcceptanceRateByHostId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) RejectPendingReservationByHost(ctx context.Context, in *RejectPendingReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error) {
	out := new(ReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/RejectPendingReservationByHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) AcceptPendingReservationByHost(ctx context.Context, in *AcceptPendingReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error) {
	out := new(ReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/AcceptPendingReservationByHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) DeletePendingReservationByGuest(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error) {
	out := new(DeleteReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/DeletePendingReservationByGuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CancelReservationByGuest(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*CancelReservationResponse, error) {
	out := new(CancelReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/CancelReservationByGuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error) {
	out := new(CreateReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/CreateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetActiveReservationsByGuestId(context.Context, *GetActiveReservationsRequest) (*GetActiveReservationsResponse, error)
	GetReservationsByUserId(context.Context, *GetReservationsByUserIdRequest) (*GetReservationsByUserIdResponse, error)
	GetReservationsByAccommodationId(context.Context, *GetReservationsByAccommodationRequest) (*GetReservationsByAccommodationResponse, error)
	GetActiveReservationsByHostId(context.Context, *GetActiveReservationsRequest) (*GetActiveReservationsResponse, error)
	GetPendingReservationsForHost(context.Context, *GetPendingReservationsForHostRequest) (*GetPendingReservationsForHostResponse, error)
	GetReservationsForHost(context.Context, *GetReservationsForHostRequest) (*GetReservationsForHostResponse, error)
	GetNumberOfPastReservationsByHostId(context.Context, *GetNumberOfPastReservationsByHostRequest) (*GetNumberOfPastReservationsByHostResponse, error)
	GetDurationOfPastReservationsByHostId(context.Context, *GetDurationOfPastReservationsByHostIdRequest) (*GetDurationOfPastReservationsByHostIdResponse, error)
	GetAcceptanceRateByHostId(context.Context, *GetAcceptanceRateByHostIdRequest) (*GetAcceptanceRateByHostIdResponse, error)
	RejectPendingReservationByHost(context.Context, *RejectPendingReservationRequest) (*ReservationResponse, error)
	AcceptPendingReservationByHost(context.Context, *AcceptPendingReservationRequest) (*ReservationResponse, error)
	DeletePendingReservationByGuest(context.Context, *DeleteReservationRequest) (*DeleteReservationResponse, error)
	CancelReservationByGuest(context.Context, *CancelReservationRequest) (*CancelReservationResponse, error)
	CreateReservation(context.Context, *CreateReservationRequest) (*CreateReservationResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedReservationServiceServer) GetActiveReservationsByGuestId(context.Context, *GetActiveReservationsRequest) (*GetActiveReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveReservationsByGuestId not implemented")
}
func (UnimplementedReservationServiceServer) GetReservationsByUserId(context.Context, *GetReservationsByUserIdRequest) (*GetReservationsByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationsByUserId not implemented")
}
func (UnimplementedReservationServiceServer) GetReservationsByAccommodationId(context.Context, *GetReservationsByAccommodationRequest) (*GetReservationsByAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationsByAccommodationId not implemented")
}
func (UnimplementedReservationServiceServer) GetActiveReservationsByHostId(context.Context, *GetActiveReservationsRequest) (*GetActiveReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveReservationsByHostId not implemented")
}
func (UnimplementedReservationServiceServer) GetPendingReservationsForHost(context.Context, *GetPendingReservationsForHostRequest) (*GetPendingReservationsForHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPendingReservationsForHost not implemented")
}
func (UnimplementedReservationServiceServer) GetReservationsForHost(context.Context, *GetReservationsForHostRequest) (*GetReservationsForHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationsForHost not implemented")
}
func (UnimplementedReservationServiceServer) GetNumberOfPastReservationsByHostId(context.Context, *GetNumberOfPastReservationsByHostRequest) (*GetNumberOfPastReservationsByHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberOfPastReservationsByHostId not implemented")
}
func (UnimplementedReservationServiceServer) GetDurationOfPastReservationsByHostId(context.Context, *GetDurationOfPastReservationsByHostIdRequest) (*GetDurationOfPastReservationsByHostIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDurationOfPastReservationsByHostId not implemented")
}
func (UnimplementedReservationServiceServer) GetAcceptanceRateByHostId(context.Context, *GetAcceptanceRateByHostIdRequest) (*GetAcceptanceRateByHostIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAcceptanceRateByHostId not implemented")
}
func (UnimplementedReservationServiceServer) RejectPendingReservationByHost(context.Context, *RejectPendingReservationRequest) (*ReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectPendingReservationByHost not implemented")
}
func (UnimplementedReservationServiceServer) AcceptPendingReservationByHost(context.Context, *AcceptPendingReservationRequest) (*ReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptPendingReservationByHost not implemented")
}
func (UnimplementedReservationServiceServer) DeletePendingReservationByGuest(context.Context, *DeleteReservationRequest) (*DeleteReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePendingReservationByGuest not implemented")
}
func (UnimplementedReservationServiceServer) CancelReservationByGuest(context.Context, *CancelReservationRequest) (*CancelReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelReservationByGuest not implemented")
}
func (UnimplementedReservationServiceServer) CreateReservation(context.Context, *CreateReservationRequest) (*CreateReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReservation not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetActiveReservationsByGuestId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveReservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetActiveReservationsByGuestId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetActiveReservationsByGuestId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetActiveReservationsByGuestId(ctx, req.(*GetActiveReservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetReservationsByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationsByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetReservationsByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetReservationsByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetReservationsByUserId(ctx, req.(*GetReservationsByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetReservationsByAccommodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationsByAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetReservationsByAccommodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetReservationsByAccommodationId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetReservationsByAccommodationId(ctx, req.(*GetReservationsByAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetActiveReservationsByHostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveReservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetActiveReservationsByHostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetActiveReservationsByHostId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetActiveReservationsByHostId(ctx, req.(*GetActiveReservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetPendingReservationsForHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPendingReservationsForHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetPendingReservationsForHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetPendingReservationsForHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetPendingReservationsForHost(ctx, req.(*GetPendingReservationsForHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetReservationsForHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationsForHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetReservationsForHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetReservationsForHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetReservationsForHost(ctx, req.(*GetReservationsForHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetNumberOfPastReservationsByHostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNumberOfPastReservationsByHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetNumberOfPastReservationsByHostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetNumberOfPastReservationsByHostId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetNumberOfPastReservationsByHostId(ctx, req.(*GetNumberOfPastReservationsByHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetDurationOfPastReservationsByHostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDurationOfPastReservationsByHostIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetDurationOfPastReservationsByHostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetDurationOfPastReservationsByHostId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetDurationOfPastReservationsByHostId(ctx, req.(*GetDurationOfPastReservationsByHostIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAcceptanceRateByHostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAcceptanceRateByHostIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAcceptanceRateByHostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetAcceptanceRateByHostId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAcceptanceRateByHostId(ctx, req.(*GetAcceptanceRateByHostIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_RejectPendingReservationByHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectPendingReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).RejectPendingReservationByHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/RejectPendingReservationByHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).RejectPendingReservationByHost(ctx, req.(*RejectPendingReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_AcceptPendingReservationByHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptPendingReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).AcceptPendingReservationByHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/AcceptPendingReservationByHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).AcceptPendingReservationByHost(ctx, req.(*AcceptPendingReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_DeletePendingReservationByGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).DeletePendingReservationByGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/DeletePendingReservationByGuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).DeletePendingReservationByGuest(ctx, req.(*DeleteReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CancelReservationByGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CancelReservationByGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/CancelReservationByGuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CancelReservationByGuest(ctx, req.(*CancelReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CreateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CreateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/CreateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CreateReservation(ctx, req.(*CreateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _ReservationService_GetAll_Handler,
		},
		{
			MethodName: "GetActiveReservationsByGuestId",
			Handler:    _ReservationService_GetActiveReservationsByGuestId_Handler,
		},
		{
			MethodName: "GetReservationsByUserId",
			Handler:    _ReservationService_GetReservationsByUserId_Handler,
		},
		{
			MethodName: "GetReservationsByAccommodationId",
			Handler:    _ReservationService_GetReservationsByAccommodationId_Handler,
		},
		{
			MethodName: "GetActiveReservationsByHostId",
			Handler:    _ReservationService_GetActiveReservationsByHostId_Handler,
		},
		{
			MethodName: "GetPendingReservationsForHost",
			Handler:    _ReservationService_GetPendingReservationsForHost_Handler,
		},
		{
			MethodName: "GetReservationsForHost",
			Handler:    _ReservationService_GetReservationsForHost_Handler,
		},
		{
			MethodName: "GetNumberOfPastReservationsByHostId",
			Handler:    _ReservationService_GetNumberOfPastReservationsByHostId_Handler,
		},
		{
			MethodName: "GetDurationOfPastReservationsByHostId",
			Handler:    _ReservationService_GetDurationOfPastReservationsByHostId_Handler,
		},
		{
			MethodName: "GetAcceptanceRateByHostId",
			Handler:    _ReservationService_GetAcceptanceRateByHostId_Handler,
		},
		{
			MethodName: "RejectPendingReservationByHost",
			Handler:    _ReservationService_RejectPendingReservationByHost_Handler,
		},
		{
			MethodName: "AcceptPendingReservationByHost",
			Handler:    _ReservationService_AcceptPendingReservationByHost_Handler,
		},
		{
			MethodName: "DeletePendingReservationByGuest",
			Handler:    _ReservationService_DeletePendingReservationByGuest_Handler,
		},
		{
			MethodName: "CancelReservationByGuest",
			Handler:    _ReservationService_CancelReservationByGuest_Handler,
		},
		{
			MethodName: "CreateReservation",
			Handler:    _ReservationService_CreateReservation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation_service.proto",
}
