package main

import (
	pb "ares/proto/gen"
	"crypto/tls"
	"fmt"
	"github.com/pires/go-proxyproto"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/anypb"
	"net"
	"os"
	"strings"
)

type server struct {
	pb.UnimplementedLinkerServer
}

func (s *server) Process(stream pb.Linker_ProcessServer) error {

	return nil
}

func main() {
	ping := pb.Ping{Serial: 1}
	a, _ := anypb.New(&ping)
	fmt.Println(a.TypeUrl)

	data, err := os.ReadFile("proto/gen/protos.desc")
	if err != nil {
		panic(err)
	}
	fds := &descriptorpb.FileDescriptorSet{}
	if err := proto.Unmarshal(data, fds); err != nil {
		panic(err)
	}
	files, err := protodesc.NewFiles(fds)
	if err != nil {
		panic(err)
	}
	files.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		messages := fd.Messages()
		for i := 0; i < messages.Len(); i++ {
			md := messages.Get(i)
			opts := md.Options().(*descriptorpb.MessageOptions)
			if opts != nil {
				ext := proto.GetExtension(opts, pb.E_TypeId)
				if typeId, ok := ext.(uint32); ok {
					fmt.Printf("name : %s typeID: %d\n", md.FullName(), typeId)
				}
			}

		}
		return true
	})

	serverCert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server-key.pem")
	if err != nil {
		panic(err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert, // 单向认证
	}

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	proxyLis := &proxyproto.Listener{Listener: lis}

	// 强制用vtprotobuf插件
	s := grpc.NewServer(
		grpc.ForceServerCodec(vtcodec.Codec{}),
		grpc.Creds(credentials.NewTLS(tlsConfig)),
		grpc.ChainStreamInterceptor(Interceptor()),
	)

	server := &server{}

	pb.RegisterLinkerServer(s, server)

	if err := s.Serve(proxyLis); err != nil {
		panic(err)
	}
}

func Interceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if p, ok := peer.FromContext(ss.Context()); ok {
			raw := p.Addr.String()
			if idx := strings.LastIndex(raw, ":"); idx != -1 {
				clientIP := raw[:idx]
				port := raw[idx+1:]
				fmt.Println(clientIP)
				fmt.Println(port)
			} else {
				fmt.Println(raw)
			}
		}

		//if md, ok := metadata.FromIncomingContext(ss.Context()); ok {
		//	if xff := md.Get("x-forwarded-for"); len(xff) > 0 {
		//		clientIP := strings.TrimSpace(strings.Split(xff[0], ",")[0])
		//		fmt.Println("Stream RPC 真实客户端 IP (X-Forwarded-For) =", clientIP)
		//	} else if xri := md.Get("x-real-ip"); len(xri) > 0 {
		//		fmt.Println("Stream RPC 真实客户端 IP (X-Real-IP) =", xri[0])
		//	}
		//}
		return nil
	}
}
