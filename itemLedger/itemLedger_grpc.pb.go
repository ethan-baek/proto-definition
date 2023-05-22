// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: itemLedger.proto

package itemLedger

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
	ItemLedger_Transfer_FullMethodName = "/itemledger.ItemLedger/Transfer"
	ItemLedger_Charge_FullMethodName   = "/itemledger.ItemLedger/Charge"
	ItemLedger_Withdraw_FullMethodName = "/itemledger.ItemLedger/Withdraw"
	ItemLedger_OwnerOf_FullMethodName  = "/itemledger.ItemLedger/OwnerOf"
	ItemLedger_Balance_FullMethodName  = "/itemledger.ItemLedger/Balance"
	ItemLedger_Release_FullMethodName  = "/itemledger.ItemLedger/Release"
)

// ItemLedgerClient is the client API for ItemLedger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemLedgerClient interface {
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	Charge(ctx context.Context, in *ChargeRequest, opts ...grpc.CallOption) (*ChargeResponse, error)
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error)
	OwnerOf(ctx context.Context, in *OwnerOfRequest, opts ...grpc.CallOption) (*OwnerOfResponse, error)
	Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error)
}

type itemLedgerClient struct {
	cc grpc.ClientConnInterface
}

func NewItemLedgerClient(cc grpc.ClientConnInterface) ItemLedgerClient {
	return &itemLedgerClient{cc}
}

func (c *itemLedgerClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, ItemLedger_Transfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemLedgerClient) Charge(ctx context.Context, in *ChargeRequest, opts ...grpc.CallOption) (*ChargeResponse, error) {
	out := new(ChargeResponse)
	err := c.cc.Invoke(ctx, ItemLedger_Charge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemLedgerClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error) {
	out := new(WithdrawResponse)
	err := c.cc.Invoke(ctx, ItemLedger_Withdraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemLedgerClient) OwnerOf(ctx context.Context, in *OwnerOfRequest, opts ...grpc.CallOption) (*OwnerOfResponse, error) {
	out := new(OwnerOfResponse)
	err := c.cc.Invoke(ctx, ItemLedger_OwnerOf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemLedgerClient) Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, ItemLedger_Balance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemLedgerClient) Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error) {
	out := new(ReleaseResponse)
	err := c.cc.Invoke(ctx, ItemLedger_Release_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemLedgerServer is the server API for ItemLedger service.
// All implementations must embed UnimplementedItemLedgerServer
// for forward compatibility
type ItemLedgerServer interface {
	Transfer(context.Context, *TransferRequest) (*TransferResponse, error)
	Charge(context.Context, *ChargeRequest) (*ChargeResponse, error)
	Withdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error)
	OwnerOf(context.Context, *OwnerOfRequest) (*OwnerOfResponse, error)
	Balance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	Release(context.Context, *ReleaseRequest) (*ReleaseResponse, error)
	mustEmbedUnimplementedItemLedgerServer()
}

// UnimplementedItemLedgerServer must be embedded to have forward compatible implementations.
type UnimplementedItemLedgerServer struct {
}

func (UnimplementedItemLedgerServer) Transfer(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedItemLedgerServer) Charge(context.Context, *ChargeRequest) (*ChargeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Charge not implemented")
}
func (UnimplementedItemLedgerServer) Withdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedItemLedgerServer) OwnerOf(context.Context, *OwnerOfRequest) (*OwnerOfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OwnerOf not implemented")
}
func (UnimplementedItemLedgerServer) Balance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedItemLedgerServer) Release(context.Context, *ReleaseRequest) (*ReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}
func (UnimplementedItemLedgerServer) mustEmbedUnimplementedItemLedgerServer() {}

// UnsafeItemLedgerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemLedgerServer will
// result in compilation errors.
type UnsafeItemLedgerServer interface {
	mustEmbedUnimplementedItemLedgerServer()
}

func RegisterItemLedgerServer(s grpc.ServiceRegistrar, srv ItemLedgerServer) {
	s.RegisterService(&ItemLedger_ServiceDesc, srv)
}

func _ItemLedger_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemLedgerServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemLedger_Transfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemLedgerServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemLedger_Charge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChargeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemLedgerServer).Charge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemLedger_Charge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemLedgerServer).Charge(ctx, req.(*ChargeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemLedger_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemLedgerServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemLedger_Withdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemLedgerServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemLedger_OwnerOf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OwnerOfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemLedgerServer).OwnerOf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemLedger_OwnerOf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemLedgerServer).OwnerOf(ctx, req.(*OwnerOfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemLedger_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemLedgerServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemLedger_Balance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemLedgerServer).Balance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemLedger_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemLedgerServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemLedger_Release_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemLedgerServer).Release(ctx, req.(*ReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ItemLedger_ServiceDesc is the grpc.ServiceDesc for ItemLedger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemLedger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "itemledger.ItemLedger",
	HandlerType: (*ItemLedgerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _ItemLedger_Transfer_Handler,
		},
		{
			MethodName: "Charge",
			Handler:    _ItemLedger_Charge_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _ItemLedger_Withdraw_Handler,
		},
		{
			MethodName: "OwnerOf",
			Handler:    _ItemLedger_OwnerOf_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _ItemLedger_Balance_Handler,
		},
		{
			MethodName: "Release",
			Handler:    _ItemLedger_Release_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "itemLedger.proto",
}