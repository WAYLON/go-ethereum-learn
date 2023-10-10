// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// ChainCollectClient is the client API for ChainCollect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChainCollectClient interface {
	// 获取区块信息
	GetBlocks(ctx context.Context, in *GetBlocksReq, opts ...grpc.CallOption) (*GetBlocksResp, error)
	// 获取交易信息
	GetTransactions(ctx context.Context, in *GetTransactionsReq, opts ...grpc.CallOption) (*GetTransactionsResp, error)
	// 获取事件信息
	GetEvents(ctx context.Context, in *GetEventsReq, opts ...grpc.CallOption) (*GetEventsResp, error)
	// 增加集合
	AddFilterSet(ctx context.Context, in *FilterSetReq, opts ...grpc.CallOption) (*FilterSetResp, error)
	// 删除集合
	DeleteFilterSet(ctx context.Context, in *FilterSetReq, opts ...grpc.CallOption) (*FilterSetResp, error)
	// 存储abi
	StorageAbi(ctx context.Context, in *StorageAbiReq, opts ...grpc.CallOption) (*NoResp, error)
	// 解析data
	ParseData(ctx context.Context, in *ParseDataReq, opts ...grpc.CallOption) (*ParseDataResp, error)
}

type chainCollectClient struct {
	cc grpc.ClientConnInterface
}

func NewChainCollectClient(cc grpc.ClientConnInterface) ChainCollectClient {
	return &chainCollectClient{cc}
}

func (c *chainCollectClient) GetBlocks(ctx context.Context, in *GetBlocksReq, opts ...grpc.CallOption) (*GetBlocksResp, error) {
	out := new(GetBlocksResp)
	err := c.cc.Invoke(ctx, "/ChainCollect.ChainCollect/GetBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainCollectClient) GetTransactions(ctx context.Context, in *GetTransactionsReq, opts ...grpc.CallOption) (*GetTransactionsResp, error) {
	out := new(GetTransactionsResp)
	err := c.cc.Invoke(ctx, "/ChainCollect.ChainCollect/GetTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainCollectClient) GetEvents(ctx context.Context, in *GetEventsReq, opts ...grpc.CallOption) (*GetEventsResp, error) {
	out := new(GetEventsResp)
	err := c.cc.Invoke(ctx, "/ChainCollect.ChainCollect/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainCollectClient) AddFilterSet(ctx context.Context, in *FilterSetReq, opts ...grpc.CallOption) (*FilterSetResp, error) {
	out := new(FilterSetResp)
	err := c.cc.Invoke(ctx, "/ChainCollect.ChainCollect/AddFilterSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainCollectClient) DeleteFilterSet(ctx context.Context, in *FilterSetReq, opts ...grpc.CallOption) (*FilterSetResp, error) {
	out := new(FilterSetResp)
	err := c.cc.Invoke(ctx, "/ChainCollect.ChainCollect/DeleteFilterSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainCollectClient) StorageAbi(ctx context.Context, in *StorageAbiReq, opts ...grpc.CallOption) (*NoResp, error) {
	out := new(NoResp)
	err := c.cc.Invoke(ctx, "/ChainCollect.ChainCollect/StorageAbi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainCollectClient) ParseData(ctx context.Context, in *ParseDataReq, opts ...grpc.CallOption) (*ParseDataResp, error) {
	out := new(ParseDataResp)
	err := c.cc.Invoke(ctx, "/ChainCollect.ChainCollect/ParseData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChainCollectServer is the server API for ChainCollect service.
// All implementations must embed UnimplementedChainCollectServer
// for forward compatibility
type ChainCollectServer interface {
	// 获取区块信息
	GetBlocks(context.Context, *GetBlocksReq) (*GetBlocksResp, error)
	// 获取交易信息
	GetTransactions(context.Context, *GetTransactionsReq) (*GetTransactionsResp, error)
	// 获取事件信息
	GetEvents(context.Context, *GetEventsReq) (*GetEventsResp, error)
	// 增加集合
	AddFilterSet(context.Context, *FilterSetReq) (*FilterSetResp, error)
	// 删除集合
	DeleteFilterSet(context.Context, *FilterSetReq) (*FilterSetResp, error)
	// 存储abi
	StorageAbi(context.Context, *StorageAbiReq) (*NoResp, error)
	// 解析data
	ParseData(context.Context, *ParseDataReq) (*ParseDataResp, error)
	mustEmbedUnimplementedChainCollectServer()
}

// UnimplementedChainCollectServer must be embedded to have forward compatible implementations.
type UnimplementedChainCollectServer struct {
}

func (UnimplementedChainCollectServer) GetBlocks(context.Context, *GetBlocksReq) (*GetBlocksResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (UnimplementedChainCollectServer) GetTransactions(context.Context, *GetTransactionsReq) (*GetTransactionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedChainCollectServer) GetEvents(context.Context, *GetEventsReq) (*GetEventsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedChainCollectServer) AddFilterSet(context.Context, *FilterSetReq) (*FilterSetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFilterSet not implemented")
}
func (UnimplementedChainCollectServer) DeleteFilterSet(context.Context, *FilterSetReq) (*FilterSetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFilterSet not implemented")
}
func (UnimplementedChainCollectServer) StorageAbi(context.Context, *StorageAbiReq) (*NoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorageAbi not implemented")
}
func (UnimplementedChainCollectServer) ParseData(context.Context, *ParseDataReq) (*ParseDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseData not implemented")
}
func (UnimplementedChainCollectServer) mustEmbedUnimplementedChainCollectServer() {}

// UnsafeChainCollectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChainCollectServer will
// result in compilation errors.
type UnsafeChainCollectServer interface {
	mustEmbedUnimplementedChainCollectServer()
}

func RegisterChainCollectServer(s grpc.ServiceRegistrar, srv ChainCollectServer) {
	s.RegisterService(&ChainCollect_ServiceDesc, srv)
}

func _ChainCollect_GetBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlocksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainCollectServer).GetBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChainCollect.ChainCollect/GetBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainCollectServer).GetBlocks(ctx, req.(*GetBlocksReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainCollect_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainCollectServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChainCollect.ChainCollect/GetTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainCollectServer).GetTransactions(ctx, req.(*GetTransactionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainCollect_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainCollectServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChainCollect.ChainCollect/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainCollectServer).GetEvents(ctx, req.(*GetEventsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainCollect_AddFilterSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterSetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainCollectServer).AddFilterSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChainCollect.ChainCollect/AddFilterSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainCollectServer).AddFilterSet(ctx, req.(*FilterSetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainCollect_DeleteFilterSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterSetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainCollectServer).DeleteFilterSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChainCollect.ChainCollect/DeleteFilterSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainCollectServer).DeleteFilterSet(ctx, req.(*FilterSetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainCollect_StorageAbi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageAbiReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainCollectServer).StorageAbi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChainCollect.ChainCollect/StorageAbi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainCollectServer).StorageAbi(ctx, req.(*StorageAbiReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainCollect_ParseData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainCollectServer).ParseData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChainCollect.ChainCollect/ParseData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainCollectServer).ParseData(ctx, req.(*ParseDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ChainCollect_ServiceDesc is the grpc.ServiceDesc for ChainCollect service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChainCollect_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChainCollect.ChainCollect",
	HandlerType: (*ChainCollectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlocks",
			Handler:    _ChainCollect_GetBlocks_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _ChainCollect_GetTransactions_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _ChainCollect_GetEvents_Handler,
		},
		{
			MethodName: "AddFilterSet",
			Handler:    _ChainCollect_AddFilterSet_Handler,
		},
		{
			MethodName: "DeleteFilterSet",
			Handler:    _ChainCollect_DeleteFilterSet_Handler,
		},
		{
			MethodName: "StorageAbi",
			Handler:    _ChainCollect_StorageAbi_Handler,
		},
		{
			MethodName: "ParseData",
			Handler:    _ChainCollect_ParseData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/pbfile/ChainCollect.proto",
}
