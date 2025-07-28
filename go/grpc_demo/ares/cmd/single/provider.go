package main

import (
	pb "ares/proto/gen"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net"
)

var provider = NewProvider()

type Provider struct {
	pb.UnimplementedProviderServer

	sessions     map[uint32]*ProviderSession
	msgProcessor map[uint32]func(session *ProviderSession, msg proto.Message) error

	auPvIds map[uint32]struct{}
}

type ProviderSession struct {
	stream          grpc.ServerStream
	info            *pb.ProvideeInfo
	checkToProvidee bool
}

func NewProvider() *Provider {
	return &Provider{
		sessions:     make(map[uint32]*ProviderSession),
		msgProcessor: make(map[uint32]func(session *ProviderSession, msg proto.Message) error),
		auPvIds:      make(map[uint32]struct{}),
	}
}

func (provider *Provider) Handler(typeId uint32, handler func(session *ProviderSession, msg proto.Message) error) {
	provider.msgProcessor[typeId] = handler
}

func (provider *Provider) Serve(stream pb.Provider_ServeServer) error {
	if p, ok := peer.FromContext(stream.Context()); ok {
		log.Println("receive providee session", p.Addr.String())
	}
	providerSession := &ProviderSession{
		stream: stream,
		info:   nil,
	}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if processor, ok := provider.msgProcessor[req.TypeId]; ok {
			if err := processor(providerSession, req); err != nil {
				log.Printf("process %d error:%v\n", req.TypeId, err)
			}
		} else {
			if req.PvId != 0 {
				if session, ok := provider.sessions[req.PvId]; ok {
					pDispatch := &pb.PDispatch{
						PvId:    req.PvId,
						TypeId:  req.TypeId,
						Payload: req.Payload,
					}
					payload, err := pDispatch.MarshalVT()
					if err != nil {
						log.Printf("marshal payload error:%v\n", err)
						continue
					}
					err = SendMsg(session.stream, 77, req.PvId, payload)
					if err != nil {
						log.Printf("pdispatch %d to %d error:%v\n", req.TypeId, req.PvId, err)
					}
				}
			}
		}
	}
}

func (provider *Provider) Start(etcdClient *clientv3.Client) {
	providerLis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatal(err)
	}

	providerGrpcServer := grpc.NewServer(
		grpc.ForceServerCodec(vtcodec.Codec{}),
		grpc.Creds(insecure.NewCredentials()),
	)

	pb.RegisterProviderServer(providerGrpcServer, provider)
	key := etcdKey(3, 301, 0)
	etcdAdd(etcdClient, key, "127.0.0.1:5001", 10)
	log.Println("provider grpc server start")
	if err := providerGrpcServer.Serve(providerLis); err != nil {
		log.Fatal(err)
	}
}
