package cluster

type ServiceDiscovery interface {
	GetServersByType(serverType uint32) (map[uint32]*Server, error)
	GetServerById(id uint32) (*Server, error)
	GetServerByPvId(pvId uint32) (*Server, error)
	GetServers() []*Server
	SyncServers(firstSync bool) error
	AddListener(listener SDListener)
}
