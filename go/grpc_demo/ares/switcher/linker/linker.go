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
	"google.golang.org/grpc/peer"
	"net"
	"time"
)

var LOGGER = logger.GetLogger("linker")

type Linker struct {
	grpcServer               *grpc.Server
	certFile, keyFile        string
	kaCheckPeriod, kaTimeout time.Duration
	address                  string
	sessions                 *Sessions
	sessionHandler           *SessionHandler
	pb.UnimplementedLinkerServer
}

func New(options ...func(*Linker)) *Linker {
	linker := &Linker{}
	for _, o := range options {
		o(linker)
	}
	linker.sessions = NewSessions()
	linker.sessionHandler = NewSessionHandler(linker)
	return linker
}

// WithCertificate "certs/server.pem" "certs/server-key.pem"
func WithCertificate(certFile, keyFile string) func(*Linker) {
	return func(l *Linker) {
		l.certFile = certFile
		l.keyFile = keyFile
	}
}

func WithAddress(address string) func(*Linker) {
	return func(l *Linker) {
		l.address = address
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

	lis, err := net.Listen("tcp", l.address)
	if err != nil {
		return err
	}

	proxyLis := ares.NewPROXYListener(lis)

	l.grpcServer = grpc.NewServer(
		// 强制用vtprotobuf插件
		grpc.ForceServerCodec(vtcodec.Codec{}),
		// tls
		grpc.Creds(credentials.NewTLS(tlsConfig)),
		// 保活设置 业务心跳给业务层写
		grpc.KeepaliveParams(kaParams),
		grpc.ChainStreamInterceptor(
			l.sessionInterceptor(),
		),
	)

	pb.RegisterLinkerServer(l.grpcServer, l)

	return l.grpcServer.Serve(proxyLis)
}

func (l *Linker) Serve(stream pb.Linker_ServeServer) error {
	// 无事可干了
	return nil
}

type streamWrapper struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *streamWrapper) Context() context.Context {
	return w.ctx
}

func (l *Linker) sessionInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx, cancel := context.WithCancel(ss.Context())
		session := NewLinkerSession(ss)
		session.SetCancel(cancel)
		newCtx := context.WithValue(ctx, ares.SessionKey, session)
		if p, ok := peer.FromContext(ss.Context()); ok {
			session.SetRemoteAddr(p.Addr)
		}
		l.sessionHandler.OnAddSession(session)
		defer l.sessionHandler.OnRemoveSession(session)
		LOGGER.Infof("[SessionInterceptor] New session incoming: %v", session)
		wrapper := &streamWrapper{ss, newCtx}

		go session.Start(newCtx)

		return handler(srv, wrapper)
	}
}
