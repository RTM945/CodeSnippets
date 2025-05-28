package discovery

import (
	"ares/configs"
	"ares/internel/discovery/etcd"
	"ares/pkg/discovery"
	"fmt"
)

func NewRegistryAndResolver(cfg *configs.Discovery) (discovery.ServiceRegistry, discovery.ServiceResolver, error) {
	switch cfg.Type {
	case "etcd":
		cli, err := etcd.NewClient(cfg.Etcd)
		if err != nil {
			return nil, nil, err
		}
		return etcd.NewRegistry(cli, cfg), etcd.NewResolver(cli, cfg), nil

	default:
		return nil, nil, fmt.Errorf("unsupported discovery type: %s", cfg.Type)
	}
}
