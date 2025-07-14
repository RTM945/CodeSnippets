package main

import (
	"errors"
	"fmt"
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

type IMsg interface {
	TypeId() uint32
}

type Msg struct {
	typeId uint32
}

func (msg Msg) TypeId() uint32 {
	return msg.typeId
}

type A struct {
	*Msg
}

func NewA() *A {
	return &A{
		Msg: &Msg{
			typeId: 1,
		},
	}
}

type TypedMsgProcessor[T IMsg] struct {
	processor func(T) error
}

func (t TypedMsgProcessor[T]) Process(msg IMsg) error {
	var typed T
	typed = msg.(T)
	return t.processor(typed)
}

type RawProcessor interface {
	Process(msg IMsg) error
}

func NewTypedMsgProcessor[T IMsg](logicProcessor interface{}) RawProcessor {
	typed := logicProcessor.(interface{ Process(T) error })
	return TypedMsgProcessor[T]{
		processor: typed.Process,
	}
}

type IMsgProcessor interface {
	Register(id uint32, f RawProcessor)
	Process(msg IMsg) error
}

type MsgProcessor struct {
	register map[uint32]RawProcessor
}

func NewMsgProcessor() *MsgProcessor {
	return &MsgProcessor{
		register: make(map[uint32]RawProcessor),
	}
}

func (mp *MsgProcessor) Register(id uint32, f RawProcessor) {
	mp.register[id] = f
}

var NoMsgProcessorErr = errors.New("no msg processor")

func (mp *MsgProcessor) Process(msg IMsg) error {
	if proc, ok := mp.register[msg.TypeId()]; ok {
		return proc.Process(msg)
	}
	return NoMsgProcessorErr
}

type AProcessor struct {
}

func (ap *AProcessor) Process(msg *A) error {
	fmt.Println("success A!")
	return nil
}

func main() {
	//ma := MsgCreator[*A]{func() *A {
	//	return &A{Msg: &Msg{1}}
	//}}
	processor := NewTypedMsgProcessor[*A](&AProcessor{})
	msgProcessor := NewMsgProcessor()
	msgProcessor.Register(1, processor)

	a := NewA()
	fmt.Println(msgProcessor.Process(a))
}
