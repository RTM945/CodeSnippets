package main

import (
	"context"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"google.golang.org/grpc"
	"grpc_demo/common"
	hellopb "grpc_demo/proto/gen/hello/v1"
	"net"
)

var logger = common.GetLogger()

type server struct {
	hellopb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloReply, error) {
	reply := &hellopb.HelloReply{
		Message: "Hello, " + req.Name,
	}

	return reply, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		logger.Error("failed to listen", "error", err)
	}

	// 注册到网关

	s := grpc.NewServer(grpc.ForceServerCodec(vtcodec.Codec{}))
	server := &server{}

	hellopb.RegisterHelloServiceServer(s, server)

	go func() {
		if err := s.Serve(lis); err != nil {

		}
	}()
	if err := s.Serve(lis); err != nil {
		logger.Error("failed to serve", "error", err)
	}

}
