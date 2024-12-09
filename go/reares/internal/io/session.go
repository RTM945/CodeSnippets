package io

import (
	"net"
	"sync"
	"sync/atomic"
)

var sessions sync.Map
var genSessionId int32

type Session struct {
	conn net.Conn
	sid  int32
}

func NewSession(conn net.Conn) *Session {
	session := &Session{
		conn: conn,
		sid:  atomic.AddInt32(&genSessionId, 1),
	}
	sessions.Store(session.sid, session)
	return session
}

func RemoveSession(session *Session) {
	sessions.Delete(session.sid)
}

func (session *Session) SetConn(conn net.Conn) {
	session.conn = conn
}

func (session *Session) Send(msg Msg) error {
	buffer := GetBuffer()
	err := EncodeMsg(buffer, msg)
	if err != nil {
		PutBuffer(buffer)
		return err
	}
	_, err = session.conn.Write(buffer.Bytes())
	if err != nil {
		PutBuffer(buffer)
		return err
	}
	PutBuffer(buffer)
	return nil
}

// type Session interface {
// 	Send(msg Msg) error
// 	SetConn(conn net.Conn)
// }

// type SessionBase struct {
// 	conn net.Conn
// 	sid  int32
// }

// func (session *SessionBase) Send(msg Msg) error {
// 	buffer := GetBuffer()
// 	err := EncodeMsg(session, buffer, msg)
// 	if err != nil {
// 		PutBuffer(buffer)
// 		return err
// 	}
// 	_, err = session.conn.Write(buffer.Bytes())
// 	if err != nil {
// 		PutBuffer(buffer)
// 		return err
// 	}
// 	PutBuffer(buffer)
// 	return nil
// }

// func (session *SessionBase) SetConn(conn net.Conn) {
// 	session.conn = conn
// }
