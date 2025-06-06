package io

import "sync"

type Session interface {
	GetSid() uint32
	Send(Msg) error
	Process(Msg)
	Close()
	OnClose()
}
type sessionKeyType struct{}

var SessionKey = sessionKeyType{}

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
