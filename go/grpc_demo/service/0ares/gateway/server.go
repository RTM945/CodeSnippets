package gateway

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_demo/common/logger"
	"grpc_demo/common/msg/gen"
	gatewaypb "grpc_demo/proto/gen/gateway/v1"
	"io"
	"net"
)

var LOGGER = logger.GetLogger("gateway")

type server struct {
	gatewaypb.UnimplementedGatewayServer
}

func (s *server) Route(stream gatewaypb.Gateway_RouteServer) error {
	ctx := stream.Context()

	eg, ctx := errgroup.WithContext(ctx)

	recvch := make(chan *gatewaypb.Envelope, 100)

	eg.Go(func() error {
		defer close(recvch)
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
			case recvch <- envelope:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	eg.Go(func() error {
		for envelope := range recvch {
			creator, ok := gen.Creator[envelope.TypeUrl]
			if !ok {
				LOGGER.Warn("no Creator for type", "type", envelope.TypeUrl)
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

			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
		}
		return nil
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func StartServer() {
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
