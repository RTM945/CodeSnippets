package main

import (
	"context"
	"fmt"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	chatpb "grpc_demo/proto/gen/chat/v1"
	hellopb "grpc_demo/proto/gen/hello/v1"
	"io"
	"log/slog"
	"os"
	"time"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))) // AddSource代码路径
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 强制用vtprotobuf插件
		grpc.WithDefaultCallOptions(grpc.ForceCodec(vtcodec.Codec{})),
	)
	if err != nil {
		slog.Error("Failed to connect to server", "error", err)
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

	<-done
}
