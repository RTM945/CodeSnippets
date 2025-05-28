package etcd

import (
	"ares/configs"
	"ares/logger"
	"ares/pkg/discovery"
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var LOGGER = logger.GetLogger("etcd registry")

type Registry struct {
	cli      *clientv3.Client
	leaseTTL int64
	leaseID  clientv3.LeaseID
	key      string
	val      string
	cfg      *configs.Discovery
}

func NewRegistry(cli *clientv3.Client, cfg *configs.Discovery) discovery.ServiceRegistry {
	return &Registry{cli: cli, leaseTTL: cfg.Etcd.LeaseTTL, cfg: cfg}
}

func (r *Registry) Register(ctx context.Context, inst discovery.ServiceInstance) error {
	r.key = fmt.Sprintf(r.cfg.KeyFormat, inst.Name, inst.ID)
	r.val = inst.Address
	// 租约
	resp, err := r.cli.Grant(ctx, r.leaseTTL)
	if err != nil {
		return err
	}
	r.leaseID = resp.ID
	if _, err := r.cli.Put(ctx, r.key, r.val, clientv3.WithLease(r.leaseID)); err != nil {
		return err
	}
	// 自动续约
	ch, err := r.cli.KeepAlive(ctx, r.leaseID)
	if err != nil {
		return err
	}
	// 消费 KeepAlive 响应，防止阻塞
	go func() {
		for leaseKeepResp := range ch {
			LOGGER.Debugf("renewing lease: %v", leaseKeepResp)
		}
		LOGGER.Infof("close lease: %v", r.leaseID)
	}()
	return nil
}

func (r *Registry) Deregister(ctx context.Context, inst discovery.ServiceInstance) error {
	_, err := r.cli.Revoke(ctx, r.leaseID)
	return err
}
