package main

import (
	pb "ares/proto/gen"
	"context"
	"fmt"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Providee struct {
	serverType uint32
	pvId       uint32
	serverId   uint32
	sessions   map[uint32]grpc.ClientStream
}

func NewProvidee(serverType, pvId, serverId uint32) *Providee {
	return &Providee{
		serverType: serverType,
		pvId:       pvId,
		serverId:   serverId,
		sessions:   make(map[uint32]grpc.ClientStream),
	}
}

func (p *Providee) Start(etcdClient *clientv3.Client) {
	// 先获取provider
	resp, err := etcdClient.Get(context.Background(), "/services/3", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}
	currentRevision := resp.Header.Revision

	for _, kv := range resp.Kvs {
		// 启动grpc client连接provider
		p.NewClient(kv)
	}

	// watch provider, 有新的就连接
	providerWC := etcdClient.Watch(context.Background(), "/services/3", clientv3.WithPrefix(), clientv3.WithPrevKV(), clientv3.WithRev(currentRevision+1))
	go func() {
		for wr := range providerWC {
			for _, ev := range wr.Events {
				switch ev.Type {
				case mvccpb.PUT:
					p.NewClient(ev.Kv)
				case mvccpb.DELETE:
					log.Printf("provider %v has been deleted\n", ev.PrevKv)
				}
			}
		}
	}()

	// 连上所有provider了，向etcd注册自己
	key := etcdKey(p.serverType, p.pvId, p.serverId)

	etcdAdd(etcdClient, key, "", 10)

	resp, err = etcdClient.Get(context.Background(), "/services", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}
	currentRevision = resp.Header.Revision

	for _, kv := range resp.Kvs {
		// 获取其他providee启动事件
		p.OnProvideeInitDone(kv)
	}

	// watch 其他 providee, 获取事件
	provideeWC := etcdClient.Watch(context.Background(), "/services", clientv3.WithPrefix(), clientv3.WithPrevKV(), clientv3.WithRev(currentRevision+1))
	go func() {
		for wr := range provideeWC {
			for _, ev := range wr.Events {
				switch ev.Type {
				case mvccpb.PUT:
					p.OnProvideeInitDone(ev.Kv)
				case mvccpb.DELETE:
					p.OnProvideeInitDone(ev.PrevKv)
				}
			}
		}
	}()
}

func (p *Providee) NewClient(kv *mvccpb.KeyValue) {
	k := string(kv.Key)
	v := string(kv.Value)
	var serverType, pvId, serverId uint32
	_, err := fmt.Sscanf(k, ectdKeyFormat, &serverType, &pvId, &serverId)
	if err != nil {
		log.Println(err)
		return
	}
	if pvId == p.pvId {
		return
	}
	if p.sessions[pvId] != nil {
		return
	}
	// 启动grpc client连接provider
	conn, err := grpc.NewClient(
		v,
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
	p.sessions[pvId] = stream
	log.Printf("providee %d bind provider %s=%s\n", p.pvId, k, v)
}

func (p *Providee) OnProvideeInitDone(kv *mvccpb.KeyValue) {
	log.Printf("providee: %v init done\n", kv)
}

func (p *Providee) OnProvideeBroken(kv *mvccpb.KeyValue) {
	log.Printf("providee: %v broken\n", kv)
}
