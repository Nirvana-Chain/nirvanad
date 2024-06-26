// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: nirvanawalletd.proto

package pb

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
	Nirvanawalletd_GetBalance_FullMethodName                 = "/nirvanawalletd.nirvanawalletd/GetBalance"
	Nirvanawalletd_GetExternalSpendableUTXOs_FullMethodName  = "/nirvanawalletd.nirvanawalletd/GetExternalSpendableUTXOs"
	Nirvanawalletd_CreateUnsignedTransactions_FullMethodName = "/nirvanawalletd.nirvanawalletd/CreateUnsignedTransactions"
	Nirvanawalletd_ShowAddresses_FullMethodName              = "/nirvanawalletd.nirvanawalletd/ShowAddresses"
	Nirvanawalletd_NewAddress_FullMethodName                 = "/nirvanawalletd.nirvanawalletd/NewAddress"
	Nirvanawalletd_Shutdown_FullMethodName                   = "/nirvanawalletd.nirvanawalletd/Shutdown"
	Nirvanawalletd_Broadcast_FullMethodName                  = "/nirvanawalletd.nirvanawalletd/Broadcast"
	Nirvanawalletd_Send_FullMethodName                       = "/nirvanawalletd.nirvanawalletd/Send"
	Nirvanawalletd_Sign_FullMethodName                       = "/nirvanawalletd.nirvanawalletd/Sign"
)

// NirvanawalletdClient is the client API for Nirvanawalletd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NirvanawalletdClient interface {
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	GetExternalSpendableUTXOs(ctx context.Context, in *GetExternalSpendableUTXOsRequest, opts ...grpc.CallOption) (*GetExternalSpendableUTXOsResponse, error)
	CreateUnsignedTransactions(ctx context.Context, in *CreateUnsignedTransactionsRequest, opts ...grpc.CallOption) (*CreateUnsignedTransactionsResponse, error)
	ShowAddresses(ctx context.Context, in *ShowAddressesRequest, opts ...grpc.CallOption) (*ShowAddressesResponse, error)
	NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
	Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
	// Since SendRequest contains a password - this command should only be used on a trusted or secure connection
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	// Since SignRequest contains a password - this command should only be used on a trusted or secure connection
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
}

type nirvanawalletdClient struct {
	cc grpc.ClientConnInterface
}

func NewNirvanawalletdClient(cc grpc.ClientConnInterface) NirvanawalletdClient {
	return &nirvanawalletdClient{cc}
}

func (c *nirvanawalletdClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, Nirvanawalletd_GetBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nirvanawalletdClient) GetExternalSpendableUTXOs(ctx context.Context, in *GetExternalSpendableUTXOsRequest, opts ...grpc.CallOption) (*GetExternalSpendableUTXOsResponse, error) {
	out := new(GetExternalSpendableUTXOsResponse)
	err := c.cc.Invoke(ctx, Nirvanawalletd_GetExternalSpendableUTXOs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nirvanawalletdClient) CreateUnsignedTransactions(ctx context.Context, in *CreateUnsignedTransactionsRequest, opts ...grpc.CallOption) (*CreateUnsignedTransactionsResponse, error) {
	out := new(CreateUnsignedTransactionsResponse)
	err := c.cc.Invoke(ctx, Nirvanawalletd_CreateUnsignedTransactions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nirvanawalletdClient) ShowAddresses(ctx context.Context, in *ShowAddressesRequest, opts ...grpc.CallOption) (*ShowAddressesResponse, error) {
	out := new(ShowAddressesResponse)
	err := c.cc.Invoke(ctx, Nirvanawalletd_ShowAddresses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nirvanawalletdClient) NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error) {
	out := new(NewAddressResponse)
	err := c.cc.Invoke(ctx, Nirvanawalletd_NewAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nirvanawalletdClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, Nirvanawalletd_Shutdown_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nirvanawalletdClient) Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, Nirvanawalletd_Broadcast_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nirvanawalletdClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, Nirvanawalletd_Send_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nirvanawalletdClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, Nirvanawalletd_Sign_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NirvanawalletdServer is the server API for Nirvanawalletd service.
// All implementations must embed UnimplementedNirvanawalletdServer
// for forward compatibility
type NirvanawalletdServer interface {
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	GetExternalSpendableUTXOs(context.Context, *GetExternalSpendableUTXOsRequest) (*GetExternalSpendableUTXOsResponse, error)
	CreateUnsignedTransactions(context.Context, *CreateUnsignedTransactionsRequest) (*CreateUnsignedTransactionsResponse, error)
	ShowAddresses(context.Context, *ShowAddressesRequest) (*ShowAddressesResponse, error)
	NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error)
	// Since SendRequest contains a password - this command should only be used on a trusted or secure connection
	Send(context.Context, *SendRequest) (*SendResponse, error)
	// Since SignRequest contains a password - this command should only be used on a trusted or secure connection
	Sign(context.Context, *SignRequest) (*SignResponse, error)
	mustEmbedUnimplementedNirvanawalletdServer()
}

// UnimplementedNirvanawalletdServer must be embedded to have forward compatible implementations.
type UnimplementedNirvanawalletdServer struct {
}

func (UnimplementedNirvanawalletdServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedNirvanawalletdServer) GetExternalSpendableUTXOs(context.Context, *GetExternalSpendableUTXOsRequest) (*GetExternalSpendableUTXOsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExternalSpendableUTXOs not implemented")
}
func (UnimplementedNirvanawalletdServer) CreateUnsignedTransactions(context.Context, *CreateUnsignedTransactionsRequest) (*CreateUnsignedTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUnsignedTransactions not implemented")
}
func (UnimplementedNirvanawalletdServer) ShowAddresses(context.Context, *ShowAddressesRequest) (*ShowAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAddresses not implemented")
}
func (UnimplementedNirvanawalletdServer) NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAddress not implemented")
}
func (UnimplementedNirvanawalletdServer) Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedNirvanawalletdServer) Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedNirvanawalletdServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedNirvanawalletdServer) Sign(context.Context, *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedNirvanawalletdServer) mustEmbedUnimplementedNirvanawalletdServer() {}

// UnsafeNirvanawalletdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NirvanawalletdServer will
// result in compilation errors.
type UnsafeNirvanawalletdServer interface {
	mustEmbedUnimplementedNirvanawalletdServer()
}

func RegisterNirvanawalletdServer(s grpc.ServiceRegistrar, srv NirvanawalletdServer) {
	s.RegisterService(&Nirvanawalletd_ServiceDesc, srv)
}

func _Nirvanawalletd_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NirvanawalletdServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nirvanawalletd_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NirvanawalletdServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nirvanawalletd_GetExternalSpendableUTXOs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExternalSpendableUTXOsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NirvanawalletdServer).GetExternalSpendableUTXOs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nirvanawalletd_GetExternalSpendableUTXOs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NirvanawalletdServer).GetExternalSpendableUTXOs(ctx, req.(*GetExternalSpendableUTXOsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nirvanawalletd_CreateUnsignedTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUnsignedTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NirvanawalletdServer).CreateUnsignedTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nirvanawalletd_CreateUnsignedTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NirvanawalletdServer).CreateUnsignedTransactions(ctx, req.(*CreateUnsignedTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nirvanawalletd_ShowAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NirvanawalletdServer).ShowAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nirvanawalletd_ShowAddresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NirvanawalletdServer).ShowAddresses(ctx, req.(*ShowAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nirvanawalletd_NewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NirvanawalletdServer).NewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nirvanawalletd_NewAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NirvanawalletdServer).NewAddress(ctx, req.(*NewAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nirvanawalletd_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NirvanawalletdServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nirvanawalletd_Shutdown_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NirvanawalletdServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nirvanawalletd_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NirvanawalletdServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nirvanawalletd_Broadcast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NirvanawalletdServer).Broadcast(ctx, req.(*BroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nirvanawalletd_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NirvanawalletdServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nirvanawalletd_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NirvanawalletdServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nirvanawalletd_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NirvanawalletdServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nirvanawalletd_Sign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NirvanawalletdServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Nirvanawalletd_ServiceDesc is the grpc.ServiceDesc for Nirvanawalletd service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nirvanawalletd_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nirvanawalletd.nirvanawalletd",
	HandlerType: (*NirvanawalletdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _Nirvanawalletd_GetBalance_Handler,
		},
		{
			MethodName: "GetExternalSpendableUTXOs",
			Handler:    _Nirvanawalletd_GetExternalSpendableUTXOs_Handler,
		},
		{
			MethodName: "CreateUnsignedTransactions",
			Handler:    _Nirvanawalletd_CreateUnsignedTransactions_Handler,
		},
		{
			MethodName: "ShowAddresses",
			Handler:    _Nirvanawalletd_ShowAddresses_Handler,
		},
		{
			MethodName: "NewAddress",
			Handler:    _Nirvanawalletd_NewAddress_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _Nirvanawalletd_Shutdown_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _Nirvanawalletd_Broadcast_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Nirvanawalletd_Send_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _Nirvanawalletd_Sign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nirvanawalletd.proto",
}
