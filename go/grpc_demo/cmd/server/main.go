package main

import (
	"context"
	"crypto/tls"
	"fmt"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	chatpb "grpc_demo/proto/gen/chat/v1"
	hellopb "grpc_demo/proto/gen/hello/v1"
	"io"
	"log/slog"
	"net"
	"os"
)

type server struct {
	hellopb.UnimplementedHelloServiceServer
	chatpb.UnimplementedChatServiceServer
}

func (s *server) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloReply, error) {
	reply := &hellopb.HelloReply{
		Message: "Hello, " + req.Name,
	}

	return reply, nil
}

func (s *server) ChatStream(stream chatpb.ChatService_ChatStreamServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			slog.Info("Client closed stream")
			return nil
		}
		if err != nil {
			slog.Error("Server recv error:", err)
			return err
		}
		slog.Info("Server recv msg:", stream.Context(), msg.String())
		reply := &chatpb.ChatMessage{
			User:    msg.User,
			Content: fmt.Sprintf("Echo: %s", msg.Content),
		}
		if err := stream.Send(reply); err != nil {
			slog.Error("Server send error:", err)
			return err
		}
	}
}

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))) // AddSource代码路径

	serverCert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server-key.pem")
	if err != nil {
		slog.Error("Server load cert/key err:", err)
		return
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert, // 单向认证
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		slog.Error("failed to listen", "error", err)
	}

	// 强制用vtprotobuf插件
	s := grpc.NewServer(grpc.ForceServerCodec(vtcodec.Codec{}), grpc.Creds(credentials.NewTLS(tlsConfig)))

	server := &server{}

	hellopb.RegisterHelloServiceServer(s, server)
	chatpb.RegisterChatServiceServer(s, server)

	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve", "error", err)
	}
}
