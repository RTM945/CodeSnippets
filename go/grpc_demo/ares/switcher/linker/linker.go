package linker

import (
	"ares/logger"
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"context"
	"crypto/tls"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"net"
	"strconv"
	"time"
)

var LOGGER = logger.GetLogger("linker")

type Linker struct {
	grpcServer               *grpc.Server
	certFile, keyFile        string
	kaCheckPeriod, kaTimeout time.Duration
	port                     int
	sessions                 *Sessions
	pb.UnimplementedLinkerServer
}

func New(options ...func(*Linker)) *Linker {
	linker := &Linker{}
	for _, o := range options {
		o(linker)
	}
	return linker
}

// WithCertificate "certs/server.pem" "certs/server-key.pem"
func WithCertificate(certFile, keyFile string) func(*Linker) {
	return func(l *Linker) {
		l.certFile = certFile
		l.keyFile = keyFile
	}
}

func WithPort(port int) func(*Linker) {
	return func(l *Linker) {
		l.port = port
	}
}

// WithKeepAlive checkPeriod 60s timeout 20s
func WithKeepAlive(checkPeriod, timeout time.Duration) func(*Linker) {
	return func(l *Linker) {
		l.kaCheckPeriod = checkPeriod
		l.kaTimeout = timeout
	}
}

func (l *Linker) Start() error {
	serverCert, err := tls.LoadX509KeyPair(l.certFile, l.keyFile)
	if err != nil {
		return err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert, // 单向认证
	}

	kaParams := keepalive.ServerParameters{
		MaxConnectionIdle:     0, // 不管 idle 多久，都不主动优雅关
		MaxConnectionAge:      0, // 不设最大年龄
		MaxConnectionAgeGrace: 0,
		Time:                  l.kaCheckPeriod, // 60s 空闲后发 PING
		Timeout:               l.kaTimeout,     // 等待 20s PING ACK
	}

	lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(l.port)))
	if err != nil {
		return err
	}

	l.grpcServer = grpc.NewServer(
		// 强制用vtprotobuf插件
		grpc.ForceServerCodec(vtcodec.Codec{}),
		// tls
		grpc.Creds(credentials.NewTLS(tlsConfig)),
		// 保活设置 业务心跳给业务层写
		grpc.KeepaliveParams(kaParams),
		grpc.ChainStreamInterceptor(),
	)

	pb.RegisterLinkerServer(l.grpcServer, l)

	return l.grpcServer.Serve(lis)
}

func (l *Linker) Serve(stream pb.Linker_ServeServer) error {

	return nil
}

func (l *Linker) realIPInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return nil
	}
}

func (l *Linker) sessionInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		LOGGER.Infof("[Interceptor] New stream incoming: %s", info.FullMethod)

		session := NewLinkerSession(ss)

		newCtx := context.WithValue(ss.Context(), ares.SessionKey, session)

		go func() {
			<-newCtx.Done()
			LOGGER.Infof("[Interceptor] Stream for session %d finished. Reason: %v", session.GetSid(), newCtx.Err())
			l.sessions.RemoveSession(session.GetSid())
		}()

		err := handler(srv, ss)

		if err != nil {
			LOGGER.Errorf("[Interceptor] Error handling session %d: %v", session.GetSid(), err)
		}
		return err
	}
}
