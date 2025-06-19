package switcher

import "ares/pkg/io"

type ILinker interface {
	GetSessions() io.ISessions
}

type IProvider interface {
	GetSessions() io.ISessions
}
