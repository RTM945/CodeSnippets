package switcher

import (
	"ares/pkg/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var LOGGER = logger.GetLogger("switcher")

var linker *Linker
var provider *Provider

func GetLinker() *Linker {
	return linker
}

func GetProvider() *Provider {
	return provider
}

var etcdClient *clientv3.Client
