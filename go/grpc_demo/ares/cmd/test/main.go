package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

//type server struct {
//	pb.UnimplementedLinkerServer
//}
//
//func (s *server) Process(stream pb.Linker_ServeServer) error {
//
//	return nil
//}

//func main() {
//	ping := pb.Ping{Serial: 1}
//	a, _ := anypb.New(&ping)
//	fmt.Println(a.TypeUrl)
//
//	data, err := os.ReadFile("proto/gen/protos.desc")
//	if err != nil {
//		panic(err)
//	}
//	fds := &descriptorpb.FileDescriptorSet{}
//	if err := proto.Unmarshal(data, fds); err != nil {
//		panic(err)
//	}
//	files, err := protodesc.NewFiles(fds)
//	if err != nil {
//		panic(err)
//	}
//	files.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
//		messages := fd.Messages()
//		for i := 0; i < messages.Len(); i++ {
//			md := messages.Get(i)
//			opts := md.Options().(*descriptorpb.MessageOptions)
//			if opts != nil {
//				ext := proto.GetExtension(opts, pb.E_TypeId)
//				if typeId, ok := ext.(uint32); ok {
//					fmt.Printf("name : %s typeID: %d\n", md.FullName(), typeId)
//				}
//			}
//
//		}
//		return true
//	})
//
//	serverCert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server-key.pem")
//	if err != nil {
//		panic(err)
//	}
//
//	tlsConfig := &tls.Config{
//		Certificates: []tls.Certificate{serverCert},
//		ClientAuth:   tls.NoClientCert, // 单向认证
//	}
//
//	lis, err := net.Listen("tcp", ":50051")
//
//	if err != nil {
//		panic(err)
//	}
//
//	proxyLis := io.NewPROXYListener(lis)
//
//	// 强制用vtprotobuf插件
//	s := grpc.NewServer(
//		grpc.ForceServerCodec(vtcodec.Codec{}),
//		grpc.Creds(credentials.NewTLS(tlsConfig)),
//		grpc.ChainStreamInterceptor(Interceptor()),
//	)
//
//	server := &server{}
//
//	pb.RegisterLinkerServer(s, server)
//
//	if err := s.Serve(proxyLis); err != nil {
//		panic(err)
//	}
//}
//
//func Interceptor() grpc.StreamServerInterceptor {
//	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
//		if p, ok := peer.FromContext(ss.Context()); ok {
//			raw := p.Addr.String()
//			if idx := strings.LastIndex(raw, ":"); idx != -1 {
//				clientIP := raw[:idx]
//				port := raw[idx+1:]
//				fmt.Println(clientIP)
//				fmt.Println(port)
//			} else {
//				fmt.Println(raw)
//			}
//		}
//
//		return nil
//	}
//}
//
//type IMsg interface {
//	TypeId() uint32
//}
//
//type Msg struct {
//	typeId uint32
//}
//
//func (msg Msg) TypeId() uint32 {
//	return msg.typeId
//}
//
//type A struct {
//	*Msg
//}
//
//func NewA() *A {
//	return &A{
//		Msg: &Msg{
//			typeId: 1,
//		},
//	}
//}
//
//type TypedMsgProcessor[T IMsg] struct {
//	processor func(T) error
//}
//
//func (t TypedMsgProcessor[T]) Process(msg IMsg) error {
//	var typed T
//	typed = msg.(T)
//	return t.processor(typed)
//}
//
//type RawProcessor interface {
//	Process(msg IMsg) error
//}
//
//func NewTypedMsgProcessor[T IMsg](logicProcessor interface{}) RawProcessor {
//	typed := logicProcessor.(interface{ Process(T) error })
//	return TypedMsgProcessor[T]{
//		processor: typed.Process,
//	}
//}
//
//type IMsgProcessor interface {
//	Register(id uint32, f RawProcessor)
//	Process(msg IMsg) error
//}
//
//type MsgProcessor struct {
//	register map[uint32]RawProcessor
//}
//
//func NewMsgProcessor() *MsgProcessor {
//	return &MsgProcessor{
//		register: make(map[uint32]RawProcessor),
//	}
//}
//
//func (mp *MsgProcessor) Register(id uint32, f RawProcessor) {
//	mp.register[id] = f
//}
//
//var NoMsgProcessorErr = errors.New("no msg processor")
//
//func (mp *MsgProcessor) Process(msg IMsg) error {
//	if proc, ok := mp.register[msg.TypeId()]; ok {
//		return proc.Process(msg)
//	}
//	return NoMsgProcessorErr
//}
//
//type AProcessor struct {
//}
//
//func (ap *AProcessor) Process(msg *A) error {
//	fmt.Println("success A!")
//	return nil
//}

//func main() {
//ma := MsgCreator[*A]{func() *A {
//	return &A{Msg: &Msg{1}}
//}}
//processor := NewTypedMsgProcessor[*A](&AProcessor{})
//msgProcessor := NewMsgProcessor()
//msgProcessor.Register(1, processor)
//
//a := NewA()
//fmt.Println(msgProcessor.Process(a))
//}

type Discovery struct {
	cli    *clientv3.Client
	prefix string // "/services/backend/"
	watchC clientv3.WatchChan
}

func NewDiscovery(endpoints []string, prefix string) (*Discovery, error) {
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		return nil, err
	}
	wc := cli.Watch(context.Background(), prefix, clientv3.WithPrefix(), clientv3.WithPrevKV())
	return &Discovery{cli, prefix, wc}, nil
}

// Run onAdd/onDel 参数会收到完整的 etcd key，例如 "/services/backend/42/10.0.0.5:50051"
func (d *Discovery) Run(onAdd func(key, val string), onDel func(key, val string)) {
	go func() {
		for wr := range d.watchC {
			for _, ev := range wr.Events {
				switch ev.Type {
				case mvccpb.PUT:
					onAdd(string(ev.Kv.Key), string(ev.Kv.Value))
				case mvccpb.DELETE:
					onDel(string(ev.PrevKv.Key), string(ev.PrevKv.Value))
				}
			}
		}
	}()
}

func (d *Discovery) Add(key, val string, ttl int64) {
	leaseResp, err := d.cli.Grant(context.Background(), ttl)
	if err != nil {
		log.Fatalf("Failed to create lease: %v", err)
	}
	key = fmt.Sprintf("%s%s", d.prefix, key)
	if _, err := d.cli.Put(context.Background(), key, val, clientv3.WithLease(leaseResp.ID)); err != nil {
		log.Fatalf("etcd put failed: %v", err)
	}
	ch, err := d.cli.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		log.Fatalf("etcd keepalive failed: %v", err)
	}
	go func() {
		for ka := range ch {
			if ka == nil {
				log.Printf("etcd lease %d expired", leaseResp.ID)
				return
			}
		}
	}()
}

type ConnManager struct {
	mu    sync.RWMutex
	conns map[string]map[string]*grpc.ClientConn // pvId → (address → conn)
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		conns: make(map[string]map[string]*grpc.ClientConn),
	}
}

func (m *ConnManager) Add(key, val string) {
	pvId := key
	addr := val

	m.mu.Lock()
	defer m.mu.Unlock()
	if m.conns[pvId] == nil {
		m.conns[pvId] = make(map[string]*grpc.ClientConn)
	}
	if _, exists := m.conns[pvId][addr]; exists {
		return
	}

	cc, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("dial %s error: %v", addr, err)
		return
	}
	m.conns[pvId][addr] = cc
}

func (m *ConnManager) Delete(key, val string) {
	pvId := key
	addr := val

	m.mu.Lock()
	defer m.mu.Unlock()
	if mp, ok := m.conns[pvId]; ok {
		if cc, exists := mp[addr]; exists {
			cc.Close()
			delete(mp, addr)
		}
		if len(mp) == 0 {
			delete(m.conns, pvId)
		}
	}
}

func (m *ConnManager) Pick(pvId string) *grpc.ClientConn {
	m.mu.RLock()
	defer m.mu.RUnlock()
	mp, ok := m.conns[pvId]
	if !ok || len(mp) == 0 {
		return nil
	}
	// 轮询或随机
	for _, cc := range mp {
		return cc
	}
	return nil
}

func main() {
	disc, err := NewDiscovery([]string{"127.0.0.1:2379"}, "/services/backend/")
	if err != nil {
		log.Fatal(err)
	}
	cm := NewConnManager()
	disc.Run(cm.Add, cm.Delete)

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	healthCheckServer := health.NewServer()
	healthCheckServer.SetServingStatus("test", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, healthCheckServer)

	disc.Add("test", "127.0.0.1:5000", 10)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		log.Printf("server stopped")
	}()

	go func() {
		key := fmt.Sprintf("%s%s", disc.prefix, "test")
		healthClient := healthpb.NewHealthClient(cm.Pick(key))
		req := &healthpb.HealthCheckRequest{
			Service: "",
		}
		resp, err := healthClient.Check(context.Background(), req)
		if err != nil {
			log.Fatalf("grpc health check failed: %v", err)
		}
		log.Println(resp)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("closing...")

	s.GracefulStop()
}
