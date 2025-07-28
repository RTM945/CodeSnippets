package main

import (
	ares "ares/pkg/io"
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
	"sync/atomic"
)

var linker = NewLinker()

type Linker struct {
	pb.UnimplementedLinkerServer

	sessions     map[uint32]grpc.ServerStream
	msgProcessor map[uint32]func(msg proto.Message) error
}

func NewLinker() *Linker {
	return &Linker{
		sessions:     make(map[uint32]grpc.ServerStream),
		msgProcessor: make(map[uint32]func(msg proto.Message) error),
	}
}

func (linker *Linker) Handler(typeId uint32, handler func(msg proto.Message) error) {
	linker.msgProcessor[typeId] = handler
}

var genSessionId atomic.Uint32

func (linker *Linker) Serve(stream pb.Linker_ServeServer) error {
	sid := genSessionId.Add(1)
	linker.sessions[sid] = stream
	if p, ok := peer.FromContext(stream.Context()); ok {
		log.Println("receive client session:", p.Addr)
	}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if processor, ok := linker.msgProcessor[req.TypeId]; ok {
			if err := processor(req); err != nil {
				log.Printf("process %d error:%v", req.TypeId, err)
			}
		} else {
			if req.PvId != 0 {
				if session, ok := provider.sessions[req.PvId]; ok {
					dispatch := &pb.Dispatch{
						ClientSid: sid,
						PvId:      req.PvId,
						TypeId:    req.TypeId,
						Payload:   req.Payload,
					}
					payload, err := dispatch.MarshalVT()
					if err != nil {
						log.Printf("marshal payload error:%v", err)
						continue
					}
					err = SendMsg(session.stream, 51, req.PvId, payload)
					if err != nil {
						log.Printf("dispatch %d to %d error:%v", req.TypeId, req.PvId, err)
					}
				} else {
					log.Println("session not found:", req.PvId)
				}
			}
		}
	}
}

func (linker *Linker) Start(etcdClient *clientv3.Client) {
	linkerLis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	linkerLis = ares.NewPROXYListener(linkerLis)

	linkerGrpcServer := grpc.NewServer(
		grpc.ForceServerCodec(vtcodec.Codec{}),
		grpc.Creds(insecure.NewCredentials()),
	)

	pb.RegisterLinkerServer(linkerGrpcServer, linker)

	key := etcdKey(2, 201, 0)
	etcdAdd(etcdClient, key, "127.0.0.1:5000", 10)
	log.Println("linker grpc server start")
	if err := linkerGrpcServer.Serve(linkerLis); err != nil {
		log.Fatal(err)
	}
}

func SendMsg(stream grpc.ServerStream, typeId, pvId uint32, data []byte) error {
	envelope := &pb.Envelope{
		TypeId:  typeId,
		PvId:    pvId,
		Payload: data,
	}
	return stream.SendMsg(envelope)
}
