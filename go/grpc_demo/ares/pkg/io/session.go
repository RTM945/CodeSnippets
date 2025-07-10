package io

import (
	pb "ares/proto/gen"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"net"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

type ISession interface {
	GetSid() uint32
	Process(IMsg)
	Send(IMsg) error
	Close()
	Node() INode
}

type Session struct {
	State
	sid         uint32
	stream      grpc.ServerStream
	node        INode
	remoteAddr  net.Addr
	ctx         context.Context
	cancel      context.CancelFunc
	sendChan    chan *pb.Envelope
	processChan chan IMsg
}

var genSessionId atomic.Uint32

var ChanSize = 64

func NewSession(stream grpc.ServerStream, node INode) *Session {
	session := &Session{
		stream:      stream,
		node:        node,
		sid:         genSessionId.Add(1),
		sendChan:    make(chan *pb.Envelope, ChanSize),
		processChan: make(chan IMsg, ChanSize),
	}
	if p, ok := peer.FromContext(stream.Context()); ok {
		session.remoteAddr = p.Addr
	}
	session.ctx, session.cancel = context.WithCancel(stream.Context())
	return session
}

func (s *Session) GetSid() uint32 {
	return s.sid
}

func (s *Session) Process(msg IMsg) {
	s.processChan <- msg
}

func (s *Session) Send(msg IMsg) error {
	payload, err := msg.Marshal()
	if err != nil {
		return err
	}

	envelope := &pb.Envelope{
		TypeId:  msg.GetType(),
		PvId:    msg.GetPvId(),
		Payload: payload,
	}

	s.sendChan <- envelope
	return nil
}

func (s *Session) Send0(msg IMsg) error {
	payload, err := msg.Marshal()
	if err != nil {
		return err
	}

	envelope := &pb.Envelope{
		TypeId:  msg.GetType(),
		PvId:    msg.GetPvId(),
		Payload: payload,
	}
	return s.stream.SendMsg(envelope)
}

func (s *Session) StartProcess() {
	defer LOGGER.Infof("session[%v] process goroutine stopped", s)
	processFunc := func(msg IMsg) {
		defer func() {
			if r := recover(); r != nil {
				LOGGER.Errorf("session[%v] panic processing msg %v: %v\n%s", s, msg, r, string(debug.Stack()))
			}
		}()

		if err := msg.Process(); err != nil {
			LOGGER.Errorf("session[%v] process msg %v err: %v", s, msg, err)
		}
	}
	for {
		if err := s.Context().Err(); err != nil {
			for m := range s.processChan {
				processFunc(m)
			}
			return
		}

		select {
		case m := <-s.processChan:
			processFunc(m)
		}
	}
}

func (s *Session) StartSend() {
	defer LOGGER.Infof("session[%v] send goroutine stopped", s)
	for {
		if err := s.Context().Err(); err != nil {
			// session close 后就不发了
			return
		}
		select {
		case envelope := <-s.sendChan:
			if err := s.stream.SendMsg(envelope); err != nil {
				LOGGER.Errorf("session[%v] send err: %v", s, err)
				return
			}
		}
	}
}

func (s *Session) String() string {
	return fmt.Sprintf("[sid = %d, remoteAddr = %s]", s.GetSid(), s.RemoteAddr())
}

func (s *Session) RemoteAddr() net.Addr {
	return s.remoteAddr
}

func (s *Session) Close() {
	s.cancel()
}

func (s *Session) Context() context.Context {
	return s.ctx
}

func (s *Session) Node() INode {
	return s.node
}

func (s *Session) OnClose() {}

type IState interface {
	AddState(state int)
	RemoveState(state int)
	CheckState(state int) bool
	GetState() int
}

type State struct {
	sync.Mutex
	state int
}

func (s *State) AddState(state int) {
	s.Lock()
	defer s.Unlock()
	s.state |= state
}

func (s *State) RemoveState(state int) {
	s.Lock()
	defer s.Unlock()
	s.state &^= state
}

func (s *State) CheckState(state int) bool {
	return s.state&state == state
}

func (s *State) GetState() int {
	return s.state
}

type ISessions interface {
	OnAddSession(ISession)
	OnRemoveSession(ISession)
	GetSession(uint32) ISession
	AllSessions() []ISession
}

type Sessions struct {
	sync.RWMutex
	allSessions map[uint32]ISession
	sessionCNT  atomic.Uint32
}

func NewSessions() *Sessions {
	return &Sessions{
		allSessions: make(map[uint32]ISession),
	}
}

func (s *Sessions) GetSession(sid uint32) ISession {
	s.RLock()
	defer s.RUnlock()
	return s.allSessions[sid]
}

func (s *Sessions) AllSessions() []ISession {
	res := make([]ISession, 0, len(s.allSessions))
	s.RLock()
	defer s.RUnlock()
	for _, v := range s.allSessions {
		res = append(res, v)
	}
	return res
}

func (s *Sessions) Size() uint32 {
	return s.sessionCNT.Load()
}

func (s *Sessions) OnAddSession(session ISession) {
	s.Lock()
	defer s.Unlock()
	s.sessionCNT.Add(1)
	s.allSessions[session.GetSid()] = session
}

func (s *Sessions) OnRemoveSession(session ISession) {
	s.Lock()
	defer s.Unlock()
	s.sessionCNT.Add(^uint32(0))
	delete(s.allSessions, session.GetSid())
}

func (s *Sessions) Stop() {
	s.Lock()
	defer s.Unlock()
	for _, v := range s.allSessions {
		v.Close()
	}
	clear(s.allSessions)
}
