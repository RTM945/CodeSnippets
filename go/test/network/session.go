package io

import (
	"log"
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

func (s *Session) Send(msg Msg) {
	buffer := GetBuffer()
	err := EncodeMsg(s, buffer, msg)
	if err != nil {
		log.Println("encode msg error:", err)
		PutBuffer(buffer)
		return
	}
	_, err = s.Conn.Write(buffer.Bytes())
	if err != nil {
		log.Println("send msg error:", err)
		PutBuffer(buffer)
		return
	}
	PutBuffer(buffer)
}
