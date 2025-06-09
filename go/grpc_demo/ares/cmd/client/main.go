package main

import (
	pb "ares/proto/gen"
	"context"
	"crypto/tls"
	"crypto/x509"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
)

func main() {
	caCert, err := os.ReadFile("certs/ca.pem")
	if err != nil {
		panic(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}
	conn, err := grpc.NewClient(
		"localhost:443",
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		// 强制用vtprotobuf插件
		grpc.WithDefaultCallOptions(grpc.ForceCodec(vtcodec.Codec{})),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewLinkerClient(conn)

	stream, err := client.Serve(context.TODO())
	if err != nil {
		panic(err)
	}
	err = stream.Send(&pb.Envelope{
		TypeId:  1,
		PvId:    1,
		Payload: nil,
	})
	if err != nil {
		panic(err)
	}
}
