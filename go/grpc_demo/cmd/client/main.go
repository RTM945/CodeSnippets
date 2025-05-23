package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	chatpb "grpc_demo/proto/gen/chat/v1"
	hellopb "grpc_demo/proto/gen/hello/v1"
	"io"
	"log/slog"
	"os"
	"time"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))) // AddSource代码路径

	caCert, err := os.ReadFile("certs/ca.pem")
	if err != nil {
		slog.Error("Error reading CA certificate", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		slog.Error("Error appending CA certificate")
	}

	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	conn, err := grpc.NewClient(
		"localhost:443",
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		// 强制用vtprotobuf插件
		grpc.WithDefaultCallOptions(grpc.ForceCodec(vtcodec.Codec{})),
	)
	if err != nil {
		slog.Error("Failed to connect to gateway", "error", err)
		return
	}
	defer conn.Close()

	helloClient := hellopb.NewHelloServiceClient(conn)

	helloReply, err := helloClient.SayHello(context.Background(), &hellopb.HelloRequest{Name: "rtm"})

	if err != nil {
		slog.Error("SayHello failed", "error", err)
		return
	}

	slog.Info(helloReply.String())

	chatClient := chatpb.NewChatServiceClient(conn)
	stream, err := chatClient.ChatStream(context.Background())
	if err != nil {
		slog.Error("ChatStream failed", "error", err)
		return
	}

	done := make(chan struct{})
	sendDone := make(chan struct{})

	go func() {
		for i := 0; i < 3; i++ {
			msg := &chatpb.ChatMessage{
				User:    "test",
				Content: fmt.Sprintf("hello %d", i),
			}
			err := stream.Send(msg)
			if err != nil {
				slog.Error("Client send error:", err)
				break
			}
		}
		close(sendDone)
	}()

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				slog.Info("Server closed stream")
				break
			}
			if err != nil {
				slog.Error("Client recv error:", err)
				break
			}
			slog.Info("Client recv msg:", msg.String())
		}
		close(done)
	}()

	select {
	case <-sendDone:
		_ = stream.CloseSend()
		slog.Info("Send finished, closed stream")
	case <-time.After(10 * time.Second):
		slog.Warn("Send timeout")
		_ = stream.CloseSend()
	}
	time.Sleep(time.Second * 10)
	<-done
}
