package main

import (
	"ares/logger"
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var etcdLogger = logger.GetLogger("etcd")

func etcdInit() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}})
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func etcdAdd(cli *clientv3.Client, key, val string, ttl int64) error {
	leaseResp, err := cli.Grant(context.Background(), ttl)
	if err != nil {
		return err
	}
	_, err = cli.Put(context.Background(), key, val, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return err
	}
	ch, err := cli.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		return err
	}
	go func() {
		for ka := range ch {
			if ka == nil {
				etcdLogger.Errorf("etcd lease %d expired", leaseResp.ID)
				return
			}
		}
	}()
	return nil
}
