package cluster

type Server struct {
	ServerType string
	PvId       uint32
	ServerId   uint32
	Host       string
}

func NewServer(serverType string, pvId uint32, serverId uint32, host string) *Server {
	return &Server{
		ServerType: serverType,
		PvId:       pvId,
		ServerId:   serverId,
		Host:       host,
	}
}
