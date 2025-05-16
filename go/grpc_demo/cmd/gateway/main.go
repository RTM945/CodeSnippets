package main

import (
	"crypto/tls"
	"fmt"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_demo/common"
	chatpb "grpc_demo/proto/gen/chat/v1"
	gatewaypb "grpc_demo/proto/gen/gateway/v1"
	hellopb "grpc_demo/proto/gen/hello/v1"
	"io"
	"log/slog"
	"net"
	"os"
)

type server struct {
	hellopb.UnimplementedHelloServiceServer
	chatpb.UnimplementedChatServiceServer
	gatewaypb.UnimplementedGatewayServer
}

func (s *server) Route(stream gatewaypb.Gateway_RouteServer) {
	msgChan := make(chan *gatewaypb.Envelope)
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				slog.Info("Client closed stream")
				return
			}
			if err != nil {
				slog.Error("Server Recv error:", err)
				return
			}
			slog.Info("Server Recv msg:", stream.Context(), msg.String())
			msgChan <- msg
		}
	}()
	go func() {
		for {
			select {
			case msg := <-msgChan:
				if msgCreator, ok := common.MsgCreator[msg.Type]; ok {
					if processor, ok := common.MsgProcessor[msg.Type]; ok {
						newMsg := msgCreator()
						err := newMsg.Unmarshal(msg.Payload)
						if err != nil {
							slog.Error("Unmarshal error:", err)
							continue
						}
						err = processor.Process(newMsg)
						if err != nil {
							slog.Error("Process error:", err)
						}
					}
				}
			}
		}
	}()
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

	serverCert, err := tls.LoadX509KeyPair("certs/gateway.pem", "certs/gateway-key.pem")
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

	//hellopb.RegisterHelloServiceServer(s, gateway)
	//chatpb.RegisterChatServiceServer(s, gateway)
	gatewaypb.RegisterGatewayServer(s, server)

	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve", "error", err)
	}
}
