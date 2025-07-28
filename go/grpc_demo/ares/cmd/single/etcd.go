package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
)

// /services/serverType/pvId/serverId
const ectdKeyFormat = "/services/%d/%d/%d"

func etcdInit() *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}})
	if err != nil {
		log.Fatal(err)
	}
	return cli
}

func etcdKey(serverType, pvId, serverId uint32) string {
	return fmt.Sprintf(ectdKeyFormat, serverType, pvId, serverId)
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
