package main

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"context"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"go.etcd.io/etcd/api/v3/mvccpb"
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
				log.Printf("process %d error:%v", req.TypeId, err)
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
						log.Printf("marshal payload error:%v", err)
						continue
					}
					err = SendMsg(session.stream, 77, req.PvId, payload)
					if err != nil {
						log.Printf("pdispatch %d to %d error:%v", req.TypeId, req.PvId, err)
					}
				}
			}
		}
	}
}

var linker = NewLinker()
var provider = NewProvider()

// Phantom 启动 onProvideeInitDone, 但由于没有连接 return
// switcher 启动, 连接Phantom, Phantom MasterNode ProvideeNode 触发onAddSession发送BindPvId
// switcher收到BindPvId, 走到给Phantom发送SPRegisgerLinker、SPRegisgerProvider、OtherProvidees、ProvideeBind
// OtherProvidees发送的是Phantom自己的providee信息
// Phantom OtherProvidees走到provideeBind, 流程中走到setOrderSession(pvid, bindPs); 向自己发送 PhantomUpdateOrderSession
// 向自己发送也是要通过provider转发的, 通过SendToProvidee协议包装先发到provider, provider再发回Phantom
// Phantom provideeBind 又调用onProvideeInitDone, 因为有连接继续, 给自己发送ProvideeInitDone

// gs流程:
// 启动connector, 从Phantom获取provider列表后连接, provider不会主动发消息
// 由providee的onAddSession给provider发送BindPvId协议, provider给所有Phantom发ProvideeBind协议
// Phantom phantomSendProvideeBind 由于本服务类型不是Phantom而退出
// 而在Phantom的流程基本是广播给已知的所有providee广播ProvideeBind和OtherProvidees
// ProvideeBind负责通知其他服, 这个服启动了
// OtherProvidees负责通知这个服, 所有已经启动的其他服
// 而下面被注入的provideeContext.provideeBind(info)会执行每个服自己的provideeBind逻辑

func main() {
	linkerLis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	linkerLis = ares.NewPROXYListener(linkerLis)

	providerLis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatal(err)
	}

	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}})
	if err != nil {
		log.Fatal(err)
	}
	linkerGrpcServer := grpc.NewServer(
		grpc.ForceServerCodec(vtcodec.Codec{}),
		grpc.Creds(insecure.NewCredentials()),
	)

	pb.RegisterLinkerServer(linkerGrpcServer, linker)

	providerGrpcServer := grpc.NewServer(
		grpc.ForceServerCodec(vtcodec.Codec{}),
		grpc.Creds(insecure.NewCredentials()),
	)

	pb.RegisterProviderServer(providerGrpcServer, provider)

	// linker
	go func() {
		etcdAdd(cli, "/services/linker", "127.0.0.1:5000", 10)
		log.Println("linker grpc server start")
		if err := linkerGrpcServer.Serve(linkerLis); err != nil {
			log.Fatal(err)
		}
	}()

	// provider
	go func() {
		etcdAdd(cli, "/services/provider", "127.0.0.1:5001", 10)
		log.Println("provider grpc server start")
		if err := providerGrpcServer.Serve(providerLis); err != nil {
			log.Fatal(err)
		}
	}()

	wc := cli.Watch(context.Background(), "/services/", clientv3.WithPrefix(), clientv3.WithPrevKV())
	// etcd watcher for client
	go func() {
		for wr := range wc {
			for _, ev := range wr.Events {
				switch ev.Type {
				case mvccpb.PUT:
					switch string(ev.Kv.Key) {
					case "/services/provider":
						newAU(string(ev.Kv.Value))
					}

				case mvccpb.DELETE:

				}
			}
		}
	}()

}

type AU struct {
	stream grpc.ClientStream
}

func newAU(addr string) *AU {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(vtcodec.Codec{})),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewProviderClient(conn)
	stream, err := client.Serve(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// 发BindPvId
	bindPvId := &pb.BindPvId{
		Info: &pb.ProvideeInfo{
			PvId:       101,
			ServerType: uint32(pb.ServerType_AU),
			ServerId:   0,
			Topics:     []string{"MsgTopic_254", "MsgTopic_253"},
			Ip:         2130706433, // 127.0.0.1
		},
		DefaultState:    0,
		CheckToProvidee: false,
	}
	payload, err := bindPvId.MarshalVT()
	if err != nil {
		log.Printf("marshal payload error:%v", err)
		return nil
	}
	err = stream.Send(&pb.Envelope{
		TypeId:  52,
		PvId:    0,
		Payload: payload,
	})
	if err != nil {
		log.Printf("send payload error:%v", err)
		return nil
	}

	return &AU{stream}
}

func etcdAdd(cli *clientv3.Client, key, val string, ttl int64) {
	leaseResp, err := cli.Grant(context.Background(), ttl)
	if err != nil {
		log.Fatalf("Failed to create lease: %v", err)
	}
	_, err = cli.Put(context.Background(), key, val, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		log.Fatalf("Failed to put lease: %v", err)
	}
	ch, err := cli.KeepAlive(context.Background(), leaseResp.ID)
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

func SendMsg(stream grpc.ServerStream, typeId, pvId uint32, data []byte) error {
	envelope := &pb.Envelope{
		TypeId:  typeId,
		PvId:    pvId,
		Payload: data,
	}
	return stream.SendMsg(envelope)
}
