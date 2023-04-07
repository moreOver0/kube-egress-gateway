// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pkg/cniprotocol/v1/cni.proto

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

// NicServiceClient is the client API for NicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NicServiceClient interface {
	// NicAdd: send pod public key and return gateway public key and endpoint ip
	NicAdd(ctx context.Context, in *NicAddRequest, opts ...grpc.CallOption) (*NicAddResponse, error)
	// NicDel: delete pod endpoint resource
	NicDel(ctx context.Context, in *NicDelRequest, opts ...grpc.CallOption) (*NicDelResponse, error)
	// PodRetrieve: send pod information and return pod information
	PodRetrieve(ctx context.Context, in *PodRetrieveRequest, opts ...grpc.CallOption) (*PodRetrieveResponse, error)
}

type nicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNicServiceClient(cc grpc.ClientConnInterface) NicServiceClient {
	return &nicServiceClient{cc}
}

func (c *nicServiceClient) NicAdd(ctx context.Context, in *NicAddRequest, opts ...grpc.CallOption) (*NicAddResponse, error) {
	out := new(NicAddResponse)
	err := c.cc.Invoke(ctx, "/pkg.cniprotocol.v1.NicService/NicAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nicServiceClient) NicDel(ctx context.Context, in *NicDelRequest, opts ...grpc.CallOption) (*NicDelResponse, error) {
	out := new(NicDelResponse)
	err := c.cc.Invoke(ctx, "/pkg.cniprotocol.v1.NicService/NicDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nicServiceClient) PodRetrieve(ctx context.Context, in *PodRetrieveRequest, opts ...grpc.CallOption) (*PodRetrieveResponse, error) {
	out := new(PodRetrieveResponse)
	err := c.cc.Invoke(ctx, "/pkg.cniprotocol.v1.NicService/PodRetrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NicServiceServer is the server API for NicService service.
// All implementations must embed UnimplementedNicServiceServer
// for forward compatibility
type NicServiceServer interface {
	// NicAdd: send pod public key and return gateway public key and endpoint ip
	NicAdd(context.Context, *NicAddRequest) (*NicAddResponse, error)
	// NicDel: delete pod endpoint resource
	NicDel(context.Context, *NicDelRequest) (*NicDelResponse, error)
	// PodRetrieve: send pod information and return pod information
	PodRetrieve(context.Context, *PodRetrieveRequest) (*PodRetrieveResponse, error)
	mustEmbedUnimplementedNicServiceServer()
}

// UnimplementedNicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNicServiceServer struct {
}

func (UnimplementedNicServiceServer) NicAdd(context.Context, *NicAddRequest) (*NicAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NicAdd not implemented")
}
func (UnimplementedNicServiceServer) NicDel(context.Context, *NicDelRequest) (*NicDelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NicDel not implemented")
}
func (UnimplementedNicServiceServer) PodRetrieve(context.Context, *PodRetrieveRequest) (*PodRetrieveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PodRetrieve not implemented")
}
func (UnimplementedNicServiceServer) mustEmbedUnimplementedNicServiceServer() {}

// UnsafeNicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NicServiceServer will
// result in compilation errors.
type UnsafeNicServiceServer interface {
	mustEmbedUnimplementedNicServiceServer()
}

func RegisterNicServiceServer(s grpc.ServiceRegistrar, srv NicServiceServer) {
	s.RegisterService(&NicService_ServiceDesc, srv)
}

func _NicService_NicAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NicAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NicServiceServer).NicAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.cniprotocol.v1.NicService/NicAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NicServiceServer).NicAdd(ctx, req.(*NicAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NicService_NicDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NicDelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NicServiceServer).NicDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.cniprotocol.v1.NicService/NicDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NicServiceServer).NicDel(ctx, req.(*NicDelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NicService_PodRetrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NicServiceServer).PodRetrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.cniprotocol.v1.NicService/PodRetrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NicServiceServer).PodRetrieve(ctx, req.(*PodRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NicService_ServiceDesc is the grpc.ServiceDesc for NicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.cniprotocol.v1.NicService",
	HandlerType: (*NicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NicAdd",
			Handler:    _NicService_NicAdd_Handler,
		},
		{
			MethodName: "NicDel",
			Handler:    _NicService_NicDel_Handler,
		},
		{
			MethodName: "PodRetrieve",
			Handler:    _NicService_PodRetrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/cniprotocol/v1/cni.proto",
}
