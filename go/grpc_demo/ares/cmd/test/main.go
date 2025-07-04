package main

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

type MsgCreatorFunc[T IMsg] func() T

type A struct {
	*Msg
}

type MsgCreator[T IMsg] struct {
	f MsgCreatorFunc[T]
}

func (mc MsgCreator[T]) Create() T {
	return mc.f()
}

func main() {
	ma := MsgCreator[*A]{func() *A {
		return &A{Msg: &Msg{1}}
	}}

}
