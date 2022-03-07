// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pkg/rpc/customer/proto/customer.proto

package customer

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

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/utl.grpc.customer.CustomerService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/utl.grpc.customer.CustomerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) Update(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/utl.grpc.customer.CustomerService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) Delete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/utl.grpc.customer.CustomerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations must embed UnimplementedCustomerServiceServer
// for forward compatibility
type CustomerServiceServer interface {
	List(context.Context, *Empty) (*Response, error)
	Create(context.Context, *Request) (*Response, error)
	Update(context.Context, *Request) (*Response, error)
	Delete(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedCustomerServiceServer()
}

// UnimplementedCustomerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (UnimplementedCustomerServiceServer) List(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCustomerServiceServer) Create(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCustomerServiceServer) Update(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCustomerServiceServer) Delete(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCustomerServiceServer) mustEmbedUnimplementedCustomerServiceServer() {}

// UnsafeCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceServer will
// result in compilation errors.
type UnsafeCustomerServiceServer interface {
	mustEmbedUnimplementedCustomerServiceServer()
}

func RegisterCustomerServiceServer(s grpc.ServiceRegistrar, srv CustomerServiceServer) {
	s.RegisterService(&CustomerService_ServiceDesc, srv)
}

func _CustomerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/utl.grpc.customer.CustomerService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/utl.grpc.customer.CustomerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Create(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/utl.grpc.customer.CustomerService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Update(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/utl.grpc.customer.CustomerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Delete(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "utl.grpc.customer.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CustomerService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CustomerService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CustomerService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CustomerService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/rpc/customer/proto/customer.proto",
}
