package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_demo/common"
	"grpc_demo/common/logger"
	chatpb "grpc_demo/proto/gen/chat/v1"
	gatewaypb "grpc_demo/proto/gen/gateway/v1"
	"io"
	"net"
)

var LOGGER = logger.GetLogger("IO")

type server struct {
	gatewaypb.UnimplementedGatewayServer
}

// Route handles a bidirectional streaming RPC by reading incoming messages,
// processing them via registered handlers, and cleanly shutting down when done.
func (s *server) Route(stream gatewaypb.Gateway_RouteServer) error {
	ctx := stream.Context()
	// Use errgroup to coordinate reader and processor
	eg, ctx := errgroup.WithContext(ctx)

	// Channel for pipelining messages; buffered to smooth bursts
	envelopeChan := make(chan *gatewaypb.Envelope, 100)

	// Reader: receive from stream and push into channel
	eg.Go(func() error {
		defer close(envelopeChan)
		for {
			envelope, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					LOGGER.Info("client closed stream")
					return nil
				}
				return fmt.Errorf("recv error: %w", err)
			}

			select {
			case envelopeChan <- envelope:
				// enqueued
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	// Processor: consume from channel and handle payloads
	eg.Go(func() error {
		for envelope := range envelopeChan {
			creator, ok := common.MsgCreator[envelope.Type]
			if !ok {
				LOGGER.Warn("no MsgCreator for type", "type", envelope.Type)
				continue
			}

			msg := creator()
			if err := msg.Unmarshal(envelope.Payload); err != nil {
				LOGGER.Error("unmarshal payload", err)
				continue
			}
			err := msg.Dispatch()
			if err != nil {
				LOGGER.Error("dispatch payload", err)
				continue
			}

			// Respect cancellation between messages
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
		}
		return nil
	})

	// Wait for both goroutines to finish
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (s *server) ChatStream(stream chatpb.ChatService_ChatStreamServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			LOGGER.Info("Client closed stream")
			return nil
		}
		if err != nil {
			LOGGER.Error("Server recv error:", err)
			return err
		}
		LOGGER.Info("Server recv msg:", stream.Context(), msg.String())
		reply := &chatpb.ChatMessage{
			User:    msg.User,
			Content: fmt.Sprintf("Echo: %s", msg.Content),
		}
		if err := stream.Send(reply); err != nil {
			LOGGER.Error("Server send error:", err)
			return err
		}
	}
}

func main() {
	serverCert, err := tls.LoadX509KeyPair("certs/gateway.pem", "certs/gateway-key.pem")
	if err != nil {
		LOGGER.Error("Server load cert/key err:", err)
		return
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert, // 单向认证
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		LOGGER.Error("failed to listen", "error", err)
		return
	}

	// 强制用vtprotobuf插件
	s := grpc.NewServer(grpc.ForceServerCodec(vtcodec.Codec{}), grpc.Creds(credentials.NewTLS(tlsConfig)))

	server := &server{}

	gatewaypb.RegisterGatewayServer(s, server)

	if err := s.Serve(lis); err != nil {
		LOGGER.Error("failed to serve", "error", err)
	}
}
