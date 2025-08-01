package etcd

import (
	"ares/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewClient(cfg *config.Etcd) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: cfg.DialTimeout,
	})
}
