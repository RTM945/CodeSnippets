package main

import (
	"ares/pkg/logger"
	pb "ares/proto/gen"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"
	"io"
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

var providerLogger = logger.GetLogger("provider")

func (provider *Provider) Handler(typeId uint32, handler func(session *ProviderSession, msg proto.Message) error) {
	provider.msgProcessor[typeId] = handler
	providerLogger.Infof("register handler type: %v", typeId)
}

func (provider *Provider) Serve(stream pb.Provider_ServeServer) error {
	if p, ok := peer.FromContext(stream.Context()); ok {
		providerLogger.Infof("receive providee session: %s", p.Addr.String())
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
				providerLogger.Errorf("process %d error:%v", req.TypeId, err)
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
						providerLogger.Errorf("marshal payload error:%v", err)
						continue
					}
					err = SendMsg(session.stream, 77, req.PvId, payload)
					if err != nil {
						providerLogger.Errorf("pdispatch %d to %d error:%v", req.TypeId, req.PvId, err)
					}
				}
			}
		}
	}
}

func (provider *Provider) Start(etcdClient *clientv3.Client) {
	providerLis, err := net.Listen("tcp", ":5001")
	if err != nil {
		providerLogger.Errorf("listen error:%v", err)
		return
	}

	providerGrpcServer := grpc.NewServer(
		grpc.ForceServerCodec(vtcodec.Codec{}),
		grpc.Creds(insecure.NewCredentials()),
	)

	pb.RegisterProviderServer(providerGrpcServer, provider)
	host := "127.0.0.1:5001"
	key := "/services/provider/3/301"
	etcdAdd(etcdClient, key, host, 10)
	providerLogger.Infof("provider grpc server start at %s", host)
	if err := providerGrpcServer.Serve(providerLis); err != nil {
		providerLogger.Errorf("provider grpc server start error:%v", err)
		return
	}
}
