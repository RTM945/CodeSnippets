package etcd

import (
	"ares/config"
	"ares/pkg/discovery"
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
)

type Resolver struct {
	sync.Mutex
	cli       *clientv3.Client
	instances []discovery.ServiceInstance
	cfg       *config.Discovery
}

func NewResolver(cli *clientv3.Client, cfg *config.Discovery) discovery.ServiceResolver {
	return &Resolver{
		cli:       cli,
		instances: make([]discovery.ServiceInstance, 0),
		cfg:       cfg,
	}
}

func (r *Resolver) GetServiceInstances() []discovery.ServiceInstance {
	r.Lock()
	defer r.Unlock()
	return r.instances
}

func (r *Resolver) List(ctx context.Context, serviceName string) ([]discovery.ServiceInstance, error) {
	prefix := fmt.Sprintf(r.cfg.KeyPrefix, serviceName)
	resp, err := r.cli.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	instances := make([]discovery.ServiceInstance, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		instances = append(instances, discovery.ServiceInstance{
			ID:      string(kv.Key[len(prefix):]),
			Name:    serviceName,
			Address: string(kv.Value),
		})
	}
	return instances, nil
}

func (r *Resolver) Watch(ctx context.Context, serviceName string) (<-chan []discovery.ServiceInstance, error) {
	prefix := fmt.Sprintf(r.cfg.KeyPrefix, serviceName)
	out := make(chan []discovery.ServiceInstance)
	// 先推送一次当前列表
	initial, err := r.List(ctx, serviceName)
	if err != nil {
		close(out)
		return nil, err
	}
	go func() {
		out <- initial
		watchCh := r.cli.Watch(ctx, prefix, clientv3.WithPrefix())
		for range watchCh {
			// 每次事件发生后重新拉取最新列表
			list, err := r.List(ctx, serviceName)
			if err == nil {
				out <- list
			}
		}
		close(out)
	}()
	return out, nil
}

//func (r *Resolver) WatchService(ctx context.Context, serviceName string) error {
//	prefix := fmt.Sprintf(r.cfg.KeyPrefix, serviceName)
//	initial, err := r.List(ctx, serviceName)
//	if err != nil {
//		return err
//	}
//	r.SetServiceInstances(initial)
//	go func() {
//		watchCh := r.cli.Watch(ctx, prefix, clientv3.WithPrefix())
//		for wr := range watchCh {
//			LOGGER.Infof("watch event: %v", wr.Events)
//			// 每次事件发生后重新拉取最新列表
//			list, err := r.List(ctx, serviceName)
//			if err != nil {
//				LOGGER.Errorf("list err: %v", err)
//			} else {
//				r.SetServiceInstances(list)
//			}
//		}
//	}()
//
//	return nil
//}
//
//func (r *Resolver) SetServiceInstances(serviceInstances []discovery.ServiceInstance) {
//	r.Lock()
//	defer r.Unlock()
//	r.instances = serviceInstances
//}
