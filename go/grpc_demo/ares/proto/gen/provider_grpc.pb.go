// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: provider.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Provider_Serve_FullMethodName = "/proto.Provider/Serve"
)

// ProviderClient is the client API for Provider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// providee <-> switcher
type ProviderClient interface {
	Serve(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Envelope, Envelope], error)
}

type providerClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderClient(cc grpc.ClientConnInterface) ProviderClient {
	return &providerClient{cc}
}

func (c *providerClient) Serve(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Envelope, Envelope], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Provider_ServiceDesc.Streams[0], Provider_Serve_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Envelope, Envelope]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Provider_ServeClient = grpc.BidiStreamingClient[Envelope, Envelope]

// ProviderServer is the server API for Provider service.
// All implementations must embed UnimplementedProviderServer
// for forward compatibility.
//
// providee <-> switcher
type ProviderServer interface {
	Serve(grpc.BidiStreamingServer[Envelope, Envelope]) error
	mustEmbedUnimplementedProviderServer()
}

// UnimplementedProviderServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProviderServer struct{}

func (UnimplementedProviderServer) Serve(grpc.BidiStreamingServer[Envelope, Envelope]) error {
	return status.Errorf(codes.Unimplemented, "method Serve not implemented")
}
func (UnimplementedProviderServer) mustEmbedUnimplementedProviderServer() {}
func (UnimplementedProviderServer) testEmbeddedByValue()                  {}

// UnsafeProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServer will
// result in compilation errors.
type UnsafeProviderServer interface {
	mustEmbedUnimplementedProviderServer()
}

func RegisterProviderServer(s grpc.ServiceRegistrar, srv ProviderServer) {
	// If the following call pancis, it indicates UnimplementedProviderServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Provider_ServiceDesc, srv)
}

func _Provider_Serve_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProviderServer).Serve(&grpc.GenericServerStream[Envelope, Envelope]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Provider_ServeServer = grpc.BidiStreamingServer[Envelope, Envelope]

// Provider_ServiceDesc is the grpc.ServiceDesc for Provider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Provider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Provider",
	HandlerType: (*ProviderServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Serve",
			Handler:       _Provider_Serve_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "provider.proto",
}
