package main

import (
	pb "ares/proto/gen"
	"context"
	"encoding/binary"
	"fmt"
	vtcodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"strconv"
)

type IPorvideeContext interface {
	OnProvideeInitDone(provideeInfo *pb.ProvideeInfo)
	OnProvideeBroken(provideeInfo *pb.ProvideeInfo)
}

type Providee struct {
	serverType uint32
	pvId       uint32
	serverId   uint32
	sessions   map[string]grpc.ClientStream

	logger *zap.SugaredLogger
}

func NewProvidee(serverType, pvId, serverId uint32, logger *zap.SugaredLogger) *Providee {
	return &Providee{
		serverType: serverType,
		pvId:       pvId,
		serverId:   serverId,
		sessions:   make(map[string]grpc.ClientStream),
		logger:     logger,
	}
}

func (providee *Providee) Start(etcdClient *clientv3.Client) error {
	// 先获取provider
	resp, err := etcdClient.Get(context.Background(), "/services/provider", clientv3.WithPrefix())
	if err != nil {
		return err
	}
	currentRevision := resp.Header.Revision

	for _, kv := range resp.Kvs {
		// 启动grpc client连接provider
		providerInfo, err := etcdEventToProviderInfo(kv)
		if err != nil {
			providee.logger.Errorf("Error converting to provider info: %v", err)
			continue
		}
		err = providee.Connect(providerInfo)
		if err != nil {
			providee.logger.Errorf("Error connecting to provider: %v", err)
			continue
		}
	}

	// watch provider, 有新的就连接
	providerWC := etcdClient.Watch(context.Background(), "/services/provider", clientv3.WithPrefix(), clientv3.WithPrevKV(), clientv3.WithRev(currentRevision+1))
	go func() {
		for wr := range providerWC {
			for _, ev := range wr.Events {
				switch ev.Type {
				case mvccpb.PUT:
					providerInfo, err := etcdEventToProviderInfo(ev.Kv)
					if err != nil {
						providee.logger.Errorf("Error converting to provider info: %v", err)
						continue
					}
					err = providee.Connect(providerInfo)
					if err != nil {
						providee.logger.Errorf("Error connecting to provider: %v", err)
						continue
					}
				case mvccpb.DELETE:
					providee.logger.Infof("provider %v has been deleted", ev.PrevKv)
				}
			}
		}
	}()

	// 连上所有provider了，向etcd注册自己
	key := fmt.Sprintf("/services/providee/%d/%d/%d", providee.serverType, providee.pvId, providee.serverId)

	etcdAdd(etcdClient, key, "127.0.0.1", 10)

	resp, err = etcdClient.Get(context.Background(), "/services/providee", clientv3.WithPrefix())
	if err != nil {
		return err
	}
	currentRevision = resp.Header.Revision

	for _, kv := range resp.Kvs {
		// 获取其他providee启动事件
		provideeInfo, err := etcdEventToProvideeInfo(kv)
		if err != nil {
			providee.logger.Errorf("Error converting to providee info: %v", err)
			continue
		}
		providee.OnProvideeInitDone(provideeInfo)
	}

	// watch 其他 providee, 获取事件
	provideeWC := etcdClient.Watch(context.Background(), "/services/providee", clientv3.WithPrefix(), clientv3.WithPrevKV(), clientv3.WithRev(currentRevision+1))
	go func() {
		for wr := range provideeWC {
			for _, ev := range wr.Events {
				switch ev.Type {
				case mvccpb.PUT:
					putProvideeInfo, err := etcdEventToProvideeInfo(ev.Kv)
					if err != nil {
						providee.logger.Errorf("Error converting to providee info: %v", err)
						continue
					}
					providee.OnProvideeInitDone(putProvideeInfo)
				case mvccpb.DELETE:
					delProvideeInfo, err := etcdEventToProvideeInfo(ev.PrevKv)
					if err != nil {
						providee.logger.Errorf("Error converting to providee info: %v", err)
						continue
					}
					providee.OnProvideeInitDone(delProvideeInfo)
				}
			}
		}
	}()
	return nil
}

func etcdEventToProviderInfo(kv *mvccpb.KeyValue) (*pb.ProviderInfo, error) {
	v := string(kv.Value)
	host, portStr, err := net.SplitHostPort(v)
	if err != nil {
		return nil, err
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}
	return &pb.ProviderInfo{
		Ip:   host,
		Port: uint32(port),
	}, nil
}

func etcdEventToProvideeInfo(kv *mvccpb.KeyValue) (*pb.ProvideeInfo, error) {
	k := string(kv.Key)
	v := string(kv.Value)
	var serverType, pvId, serverId uint32
	_, err := fmt.Sscanf(k, "/services/providee/%d/%d/%d", &serverType, &pvId, &serverId)
	if err != nil {
		return nil, err
	}

	return &pb.ProvideeInfo{
		PvId:       pvId,
		ServerType: serverType,
		ServerId:   serverId,
		Topics:     make([]string, 0),
		Ip:         binary.BigEndian.Uint32(net.ParseIP(v).To4()),
	}, nil
}

func (providee *Providee) Connect(providerInfo *pb.ProviderInfo) error {
	host := net.JoinHostPort(providerInfo.Ip, strconv.Itoa(int(providerInfo.Port)))
	if providee.sessions[host] != nil {
		return fmt.Errorf("provider %v has already been bind", providerInfo)
	}
	// 启动grpc client连接provider
	conn, err := grpc.NewClient(
		host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(vtcodec.Codec{})),
	)
	if err != nil {
		return err
	}
	client := pb.NewProviderClient(conn)
	stream, err := client.Serve(context.TODO())
	if err != nil {
		return err
	}
	providee.sessions[host] = stream
	return nil
}

func (providee *Providee) OnProvideeInitDone(provideeInfo *pb.ProvideeInfo) {
	providee.logger.Infof("providee: %v init done", provideeInfo)
}

func (providee *Providee) OnProvideeBroken(provideeInfo *pb.ProvideeInfo) {
	providee.logger.Infof("providee: %v broken", provideeInfo)
}
