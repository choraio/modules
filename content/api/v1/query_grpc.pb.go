// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: chora/content/v1/query.proto

package contentv1

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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Content queries a content by id.
	Content(ctx context.Context, in *QueryContentRequest, opts ...grpc.CallOption) (*QueryContentResponse, error)
	// Contents queries all contents.
	Contents(ctx context.Context, in *QueryContentsRequest, opts ...grpc.CallOption) (*QueryContentsResponse, error)
	// ContentsByCurator queries contents by curator.
	ContentsByCurator(ctx context.Context, in *QueryContentsByCuratorRequest, opts ...grpc.CallOption) (*QueryContentsByCuratorResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Content(ctx context.Context, in *QueryContentRequest, opts ...grpc.CallOption) (*QueryContentResponse, error) {
	out := new(QueryContentResponse)
	err := c.cc.Invoke(ctx, "/chora.content.v1.Query/Content", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Contents(ctx context.Context, in *QueryContentsRequest, opts ...grpc.CallOption) (*QueryContentsResponse, error) {
	out := new(QueryContentsResponse)
	err := c.cc.Invoke(ctx, "/chora.content.v1.Query/Contents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContentsByCurator(ctx context.Context, in *QueryContentsByCuratorRequest, opts ...grpc.CallOption) (*QueryContentsByCuratorResponse, error) {
	out := new(QueryContentsByCuratorResponse)
	err := c.cc.Invoke(ctx, "/chora.content.v1.Query/ContentsByCurator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Content queries a content by id.
	Content(context.Context, *QueryContentRequest) (*QueryContentResponse, error)
	// Contents queries all contents.
	Contents(context.Context, *QueryContentsRequest) (*QueryContentsResponse, error)
	// ContentsByCurator queries contents by curator.
	ContentsByCurator(context.Context, *QueryContentsByCuratorRequest) (*QueryContentsByCuratorResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Content(context.Context, *QueryContentRequest) (*QueryContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Content not implemented")
}
func (UnimplementedQueryServer) Contents(context.Context, *QueryContentsRequest) (*QueryContentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Contents not implemented")
}
func (UnimplementedQueryServer) ContentsByCurator(context.Context, *QueryContentsByCuratorRequest) (*QueryContentsByCuratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContentsByCurator not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Content_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Content(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chora.content.v1.Query/Content",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Content(ctx, req.(*QueryContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Contents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Contents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chora.content.v1.Query/Contents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Contents(ctx, req.(*QueryContentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContentsByCurator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContentsByCuratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContentsByCurator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chora.content.v1.Query/ContentsByCurator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContentsByCurator(ctx, req.(*QueryContentsByCuratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chora.content.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Content",
			Handler:    _Query_Content_Handler,
		},
		{
			MethodName: "Contents",
			Handler:    _Query_Contents_Handler,
		},
		{
			MethodName: "ContentsByCurator",
			Handler:    _Query_ContentsByCurator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chora/content/v1/query.proto",
}