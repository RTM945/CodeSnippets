package cluster

import (
	"ares/pkg/logger"
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/client/pkg/v3/logutil"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/namespace"
	"google.golang.org/grpc"
	"time"
)

const ServerETCDPrefix = "/servers/"
const ServerETCDFormat = "/server/%s/%s"

type etcdServiceDiscovery struct {
	cli                *clientv3.Client
	etcdEndpoints      []string
	etcdDialTimeout    time.Duration
	etcdPrefix         string
	running            bool
	syncServersRunning chan bool
}

func (sd *etcdServiceDiscovery) watchEtcdChanges() {
	w := sd.cli.Watch(context.Background(), ServerETCDPrefix, clientv3.WithPrefix())
	failedWatchAttempts := 0
	go func(chn clientv3.WatchChan) {
		for sd.running {
			select {
			// Block here if SyncServers() is running and consume the watcher channel after it's finished, to avoid conflicts
			case syncServersState := <-sd.syncServersRunning:
				for syncServersState {
					syncServersState = <-sd.syncServersRunning
				}
			case wResp, ok := <-chn:
				if wResp.Err() != nil {
					logger.Log.Warnf("etcd watcher response error: %s", wResp.Err())
					time.Sleep(100 * time.Millisecond)
				}
				if !ok {
					logger.Log.Error("etcd watcher died, retrying to watch in 1 second")
					failedWatchAttempts++
					time.Sleep(1000 * time.Millisecond)
					if failedWatchAttempts > 10 {
						if err := sd.InitETCDClient(); err != nil {
							failedWatchAttempts = 0
							continue
						}
						chn = sd.cli.Watch(context.Background(), ServerETCDPrefix, clientv3.WithPrefix())
						failedWatchAttempts = 0
					}
					continue
				}
				failedWatchAttempts = 0
				for _, ev := range wResp.Events {
					_, _, err := parseEtcdKey(string(ev.Kv.Key))
					if err != nil {
						logger.Log.Warnf("failed to parse key from etcd: %s", ev.Kv.Key)
						continue
					}
					switch ev.Type {
					case clientv3.EventTypePut:
						var sv *Server
						var err error
						if sv, err = parseServer(ev.Kv.Value); err != nil {
							logger.Log.Errorf("Failed to parse server from etcd: %v", err)
							continue
						}
						// TODO addServer
						logger.Log.Debugf("server %v added by watcher", sv)
					case clientv3.EventTypeDelete:

					}
				}
			}
		}
	}(w)
}

func parseEtcdKey(key string) (uint32, uint32, error) {
	var serverType, pvId uint32
	_, err := fmt.Sscanf(key, ServerETCDFormat, &serverType, &pvId)
	if err != nil {
		return 0, 0, err
	}
	return serverType, pvId, nil
}

func parseServer(value []byte) (*Server, error) {
	var sv *Server
	err := json.Unmarshal(value, &sv)
	if err != nil {
		logger.Log.Warnf("failed to load server %v, error: %v", sv, err)
		return nil, err
	}
	return sv, nil
}

func (sd *etcdServiceDiscovery) InitETCDClient() error {
	logger.Log.Infof("Initializing ETCD client")
	var cli *clientv3.Client
	var err error
	etcdClientLogger, _ := logutil.CreateDefaultZapLogger(logutil.ConvertToZapLevel("error"))
	config := clientv3.Config{
		Endpoints:   sd.etcdEndpoints,
		DialTimeout: sd.etcdDialTimeout,
		Logger:      etcdClientLogger,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	}
	cli, err = clientv3.New(config)
	if err != nil {
		logger.Log.Errorf("error initializing etcd client: %v", err)
		return err
	}
	sd.cli = cli

	sd.cli.KV = namespace.NewKV(sd.cli.KV, sd.etcdPrefix)
	sd.cli.Watcher = namespace.NewWatcher(sd.cli.Watcher, sd.etcdPrefix)
	sd.cli.Lease = namespace.NewLease(sd.cli.Lease, sd.etcdPrefix)
	return nil
}
