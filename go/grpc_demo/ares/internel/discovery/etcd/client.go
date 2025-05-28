package etcd

import (
	"ares/configs"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewClient(cfg *configs.Etcd) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: cfg.DialTimeout,
	})
}
