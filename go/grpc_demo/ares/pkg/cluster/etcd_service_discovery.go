package cluster

import (
	"ares/config"
	"ares/pkg/logger"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"go.etcd.io/etcd/client/pkg/v3/logutil"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"strconv"
	"sync"
	"time"
)

const ServerETCDKeyPrefix = "/servers/"

// FrontendETCDKeyPrefix
// /servers/linker/127.0.0.1:9090
// /servers/provider/127.0.0.1:9091
// /servers/au/101
// /servers/phantom/301
// /servers/game/5001->5001

// ServerETCDKeyFormat key = /servers/type/<id or host>  value = pvId
const ServerETCDKeyFormat = "/servers/%s/%s"

var ErrEtcdGrantLeaseTimeout = errors.New("timed out waiting for etcd lease grant")

type etcdServiceDiscovery struct {
	cli                 *clientv3.Client
	syncServersInterval time.Duration
	heartbeatTTL        time.Duration
	logHeartbeat        bool
	etcdEndpoints       []string
	etcdDialTimeout     time.Duration
	running             bool
	leaseID             clientv3.LeaseID

	syncServersRunning chan bool
	stopChan           chan bool
	stopLeaseChan      chan bool
	lastSyncTime       time.Time
	appDieChan         chan bool

	mapByTypeLock   sync.RWMutex
	serverMapByType map[uint32]map[uint32]*Server
	serverMapByID   sync.Map

	server               *Server
	revokeTimeout        time.Duration
	grantLeaseTimeout    time.Duration
	grantLeaseMaxRetries int
	grantLeaseInterval   time.Duration
	shutdownDelay        time.Duration

	listeners []SDListener
}

func NewEtcdServiceDiscovery(config config.EtcdServiceDiscoveryConfig, server *Server, appDieChan chan bool) (ServiceDiscovery, error) {
	sd := &etcdServiceDiscovery{
		running:            false,
		server:             server,
		serverMapByType:    make(map[uint32]map[uint32]*Server),
		listeners:          make([]SDListener, 0),
		stopChan:           make(chan bool),
		stopLeaseChan:      make(chan bool),
		appDieChan:         appDieChan,
		syncServersRunning: make(chan bool),
	}

	sd.configure(config)

	return sd, nil
}

func (sd *etcdServiceDiscovery) configure(config config.EtcdServiceDiscoveryConfig) {
	sd.etcdEndpoints = config.Endpoints
	sd.etcdDialTimeout = config.DialTimeout
	sd.heartbeatTTL = config.Heartbeat.TTL
	sd.logHeartbeat = config.Heartbeat.Log
	sd.syncServersInterval = config.SyncServers.Interval
	sd.revokeTimeout = config.Revoke.Timeout
	sd.grantLeaseTimeout = config.GrantLease.Timeout
	sd.grantLeaseMaxRetries = config.GrantLease.MaxRetries
	sd.grantLeaseInterval = config.GrantLease.RetryInterval
	sd.shutdownDelay = config.Shutdown.Delay
}

func (sd *etcdServiceDiscovery) Init() error {
	sd.running = true
	var err error

	err = sd.InitETCDClient()
	if err != nil {
		return err
	}

	if sd.server.Frontend {
		// 如果是 frontend 需要获取其他frontend
		sd.SyncServers(true)

	} else {

	}

	go sd.watchEtcdChanges()

	if err = sd.bootstrap(); err != nil {
		return err
	}

	// update servers
	syncServersTicker := time.NewTicker(sd.syncServersInterval)
	go func() {
		for sd.running {
			select {
			case <-syncServersTicker.C:
				err := sd.SyncServers(false)
				if err != nil {
					logger.Log.Errorf("error resyncing servers: %s", err.Error())
				}
			case <-sd.stopChan:
				return
			}
		}
	}()

	return nil
}

func (sd *etcdServiceDiscovery) GetServersByType(serverType uint32) (map[uint32]*Server, error) {
	//TODO implement me
	panic("implement me")
}

func (sd *etcdServiceDiscovery) GetServerById(id uint32) (*Server, error) {
	//TODO implement me
	panic("implement me")
}

func (sd *etcdServiceDiscovery) GetServerByPvId(pvId uint32) (*Server, error) {
	//TODO implement me
	panic("implement me")
}

func (sd *etcdServiceDiscovery) GetServers() []*Server {
	//TODO implement me
	panic("implement me")
}

func (sd *etcdServiceDiscovery) AddListener(listener SDListener) {
	//TODO implement me
	panic("implement me")
}

func (sd *etcdServiceDiscovery) bootstrap() error {
	if err := sd.grantLease(); err != nil {
		return err
	}

	if err := sd.bootstrapServer(sd.server); err != nil {
		return err
	}

	return nil
}

func (sd *etcdServiceDiscovery) grantLease() error {
	// grab lease
	ctx, cancel := context.WithTimeout(context.Background(), sd.etcdDialTimeout)
	defer cancel()
	l, err := sd.cli.Grant(ctx, int64(sd.heartbeatTTL.Seconds()))
	if err != nil {
		return err
	}
	sd.leaseID = l.ID
	logger.Log.Debugf("sd: got leaseID: %x", l.ID)
	// this will keep alive forever, when channel c is closed
	// it means we probably have to rebootstrap the lease
	c, err := sd.cli.KeepAlive(context.TODO(), sd.leaseID)
	if err != nil {
		return err
	}
	// need to receive here as per etcd docs
	<-c
	go sd.watchLeaseChan(c)
	return nil
}

func (sd *etcdServiceDiscovery) watchLeaseChan(c <-chan *clientv3.LeaseKeepAliveResponse) {
	failedGrantLeaseAttempts := 0
	for {
		select {
		case <-sd.stopChan:
			return
		case <-sd.stopLeaseChan:
			return
		case leaseKeepAliveResponse, ok := <-c:
			if !ok {
				logger.Log.Error("ETCD lease KeepAlive died, retrying in 10 seconds")
				time.Sleep(10000 * time.Millisecond)
			}
			if leaseKeepAliveResponse != nil {
				if sd.logHeartbeat {
					logger.Log.Debugf("sd: etcd lease %x renewed", leaseKeepAliveResponse.ID)
				}
				failedGrantLeaseAttempts = 0
				continue
			}
			logger.Log.Warn("sd: error renewing etcd lease, reconfiguring")
			for {
				err := sd.renewLease()
				if err != nil {
					failedGrantLeaseAttempts = failedGrantLeaseAttempts + 1
					if err == ErrEtcdGrantLeaseTimeout {
						logger.Log.Warn("sd: timed out trying to grant etcd lease")
						if sd.appDieChan != nil {
							sd.appDieChan <- true
						}
						return
					}
					if failedGrantLeaseAttempts >= sd.grantLeaseMaxRetries {
						logger.Log.Warn("sd: exceeded max attempts to renew etcd lease")
						if sd.appDieChan != nil {
							sd.appDieChan <- true
						}
						return
					}
					logger.Log.Warnf("sd: error granting etcd lease, will retry in %d seconds", uint64(sd.grantLeaseInterval.Seconds()))
					time.Sleep(sd.grantLeaseInterval)
					continue
				}
				return
			}
		}
	}
}

// renewLease reestablishes connection with etcd
func (sd *etcdServiceDiscovery) renewLease() error {
	c := make(chan error, 1)
	go func() {
		defer close(c)
		logger.Log.Infof("waiting for etcd lease")
		err := sd.grantLease()
		if err != nil {
			c <- err
			return
		}
		err = sd.bootstrapServer(sd.server)
		c <- err
	}()
	select {
	case err := <-c:
		return err
	case <-time.After(sd.grantLeaseTimeout):
		return ErrEtcdGrantLeaseTimeout
	}
}

func (sd *etcdServiceDiscovery) bootstrapServer(server *Server) error {
	if err := sd.addServerIntoEtcd(server); err != nil {
		return err
	}

	sd.SyncServers(true)
	return nil
}

func (sd *etcdServiceDiscovery) addServerIntoEtcd(server *Server) error {
	_, err := sd.cli.Put(
		context.TODO(),
		getKey(server.ID, server.Type),
		strconv.Itoa(int(server.PvId)),
		clientv3.WithLease(sd.leaseID),
	)
	return err
}

func getServerFromEtcd(cli *clientv3.Client, serverType, serverID string) (*Server, error) {
	svKey := getKey(serverID, serverType)
	svEInfo, err := cli.Get(context.TODO(), svKey)
	if err != nil {
		return nil, fmt.Errorf("error getting server: %s from etcd, error: %v", svKey, err)
	}
	if len(svEInfo.Kvs) == 0 {
		return nil, fmt.Errorf("didn't found server: %s in etcd", svKey)
	}
	return parseServer(svEInfo.Kvs[0].Value)
}

func getKey(serverID, serverType uint32) string {
	return fmt.Sprintf(ServerETCDKeyFormat, strconv.Itoa(int(serverType)), strconv.Itoa(int(serverID)))
}

func (sd *etcdServiceDiscovery) watchEtcdChanges() {
	w := sd.cli.Watch(context.Background(), ServerETCDKeyPrefix, clientv3.WithPrefix())
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
						chn = sd.cli.Watch(context.Background(), ServerETCDKeyPrefix, clientv3.WithPrefix())
						failedWatchAttempts = 0
					}
					continue
				}
				failedWatchAttempts = 0
				for _, ev := range wResp.Events {
					_, svID, err := parseEtcdKey(string(ev.Kv.Key))
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
						sd.addServer(sv)
						logger.Log.Debugf("server %s added by watcher", ev.Kv.Key)
						sd.printServers()
					case clientv3.EventTypeDelete:
						sd.deleteServer(svID)
						logger.Log.Debugf("server %s deleted by watcher", ev.Kv.Key)
						sd.printServers()
					}
				}
			case <-sd.stopChan:
				return
			}
		}
	}(w)
}

func parseEtcdKey(key string) (string, string, error) {
	var svType, svID string
	_, err := fmt.Sscanf(key, ServerETCDKeyFormat, &svType, &svID)
	if err != nil {
		return "", "", err
	}
	return svType, svID, nil
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
	return nil
}

func (sd *etcdServiceDiscovery) addServer(sv *Server) {
	if _, loaded := sd.serverMapByID.LoadOrStore(sv.ID, sv); !loaded {
		sd.writeLockScope(func() {
			mapSvByType, ok := sd.serverMapByType[sv.Type]
			if !ok {
				mapSvByType = make(map[string]*Server)
				sd.serverMapByType[sv.Type] = mapSvByType
			}
			mapSvByType[sv.ID] = sv
		})
		if sv.ID != sd.server.ID {
			sd.notifyListeners(ADD, sv)
		}
	}
}

func (sd *etcdServiceDiscovery) deleteServer(serverID string) {
	if actual, ok := sd.serverMapByID.Load(serverID); ok {
		sv := actual.(*Server)
		sd.serverMapByID.Delete(sv.ID)
		sd.writeLockScope(func() {
			if svMap, ok := sd.serverMapByType[sv.Type]; ok {
				delete(svMap, sv.ID)
			}
		})
		sd.notifyListeners(DEL, sv)
	}
}

func (sd *etcdServiceDiscovery) writeLockScope(f func()) {
	sd.mapByTypeLock.Lock()
	defer sd.mapByTypeLock.Unlock()
	f()
}

func (sd *etcdServiceDiscovery) notifyListeners(act Action, sv *Server) {
	for _, l := range sd.listeners {
		if act == DEL {
			l.RemoveServer(sv)
		} else if act == ADD {
			l.AddServer(sv)
		}
	}
}

func (sd *etcdServiceDiscovery) printServers() {
	sd.mapByTypeLock.RLock()
	defer sd.mapByTypeLock.RUnlock()
	for k, v := range sd.serverMapByType {
		logger.Log.Debugf("type: %s, servers: %+v", k, v)
	}
}

// SyncServers gets all servers from etcd
func (sd *etcdServiceDiscovery) SyncServers(firstSync bool) error {
	sd.syncServersRunning <- true
	defer func() {
		sd.syncServersRunning <- false
	}()
	start := time.Now()
	var kvs *clientv3.GetResponse
	var err error
	kvs, err = sd.cli.Get(
		context.TODO(),
		ServerETCDKeyPrefix,
		clientv3.WithPrefix(),
	)
	if err != nil {
		logger.Log.Errorf("Error querying etcd server: %v", err)
		return err
	}

	var allIds = make([]string, 0)

	for _, kv := range kvs.Kvs {
		svType, svID, err := parseEtcdKey(string(kv.Key))
		if err != nil {
			logger.Log.Warnf("failed to parse etcd key %s, error: %s", kv.Key, err.Error())
			continue
		}

		allIds = append(allIds, svID)

		if _, exists := sd.serverMapByID.Load(svID); !exists {
			var sv *Server
			sv, err = parseServer(kv.Value)
			if err != nil {
				logger.Log.Errorf("error loading server %s/%s: %s", svType, svID, err)
				continue
			}
			logger.Log.Debugf("adding server %v", sv)
			sd.addServer(sv)
		}
	}

	for _, server := range servers {
		logger.Log.Debugf("adding server %v", server)
		sd.addServer(server)
	}

	sd.deleteLocalInvalidServers(allIds)

	sd.printServers()
	sd.lastSyncTime = time.Now()
	elapsed := time.Since(start)
	logger.Log.Infof("SyncServers took : %s to run", elapsed)
	return nil
}

func (sd *etcdServiceDiscovery) deleteLocalInvalidServers(actualServers []string) {
	sd.serverMapByID.Range(func(key interface{}, value interface{}) bool {
		k := key.(string)
		if !lo.Contains(actualServers, k) {
			logger.Log.Warnf("deleting invalid local server %s", k)
			sd.deleteServer(k)
		}
		return true
	})
}
