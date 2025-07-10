package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"crypto/tls"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"io"
	"net"
	"time"
)

type Linker struct {
	grpcServer               *grpc.Server
	certFile, keyFile        string
	kaCheckPeriod, kaTimeout time.Duration
	address                  string
	sessions                 ares.ISessions
	maxSession               uint32
	sessionTimeout           int64
	whiteIps, blackIps       []string
	rateMin, rateMax         int
	msgCreator               ares.IMsgCreator

	OnSessionError func(session *LinkerSession, code uint32) error
	OnServerError  func(session *LinkerSession, pvId, code uint32) error
	OnDispatch     func(session *ProviderSession, pvId, typeId uint32, payload []byte) error

	pb.UnimplementedLinkerServer
}

func (l *Linker) Sessions() ares.ISessions {
	return l.sessions
}

func (l *Linker) MsgCreator() ares.IMsgCreator {
	return l.msgCreator
}

func New(options ...func(*Linker)) *Linker {
	linker := &Linker{}
	for _, o := range options {
		o(linker)
	}
	linker.sessions = NewLinkerSessions()
	linker.msgCreator = NewLinkerMsgCreator()
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

func WithRateLimit(min, max int) func(*Linker) {
	return func(l *Linker) {
		l.rateMin = min
		l.rateMax = max
	}
}

func WithWhiteIps(whiteIps []string) func(*Linker) {
	return func(l *Linker) {
		l.whiteIps = whiteIps
	}
}

func WithBlackIps(blackIps []string) func(*Linker) {
	return func(l *Linker) {
		l.blackIps = blackIps
	}
}

func WithSessionTimeout(timeout int64) func(*Linker) {
	return func(l *Linker) {
		l.sessionTimeout = timeout
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
	)

	pb.RegisterLinkerServer(l.grpcServer, l)

	return l.grpcServer.Serve(proxyLis)
}

func (l *Linker) Serve(stream pb.Linker_ServeServer) error {
	session := NewLinkerSession(stream, l)
	l.sessions.OnAddSession(session)
	defer l.sessions.OnRemoveSession(session)

	go session.StartSend()
	go session.StartProcess()

	for {
		select {
		case <-session.Context().Done():
			// session close 后就不收了
			return nil
		default:
			envelope, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					return nil // 客户端正常关闭
				}
				return err
			}

			session.HandleEnvelope(envelope)
		}
	}
}

func (l *Linker) GetWhiteIps() map[string]struct{} {
	tmp := make(map[string]struct{}, len(l.whiteIps))
	for _, ip := range l.whiteIps {
		tmp[ip] = struct{}{}
	}
	return tmp
}

func (l *Linker) GetBlackIps() map[string]struct{} {
	tmp := make(map[string]struct{}, len(l.blackIps))
	for _, ip := range l.blackIps {
		tmp[ip] = struct{}{}
	}
	return tmp
}
