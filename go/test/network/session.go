package network

import (
	"net"
	"sync"
)

var sessions sync.Map

type Session struct {
	Conn net.Conn
}

func NewSession(conn net.Conn) *Session {
	session := &Session{
		Conn: conn,
	}
	sessions.Store(session, 0)
	return session
}

func RemoveSession(session *Session) {
	sessions.Delete(session)
}
