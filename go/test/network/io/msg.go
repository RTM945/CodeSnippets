package io

import (
	"bytes"
	"context"
	"encoding/binary"
)

type Coder interface {
	Decode(src *bytes.Buffer) error
	Encode(dst *bytes.Buffer) error
}

type Task interface {
	Execute()
}

type Msg interface {
	Coder
	Task
	Process()
	Dispatch()
	GetSession() *Session
	SetSession(session *Session)
	GetHeader() *Header
	SetHeader(h *Header)
	GetContext() context.Context
	SetContext(ctx context.Context)
}

type BaseMsg struct {
	session *Session
	header  *Header
	ctx     context.Context
}

func NewBaseMsg(session *Session, header *Header) *BaseMsg {
	return &BaseMsg{
		session: session,
		header:  header,
	}
}

func (m *BaseMsg) SetSession(session *Session) {
	m.session = session
}

func (m *BaseMsg) GetSession() *Session {
	return m.session
}

func (m *BaseMsg) SetContext(ctx context.Context) {
	m.ctx = ctx
}

func (m *BaseMsg) GetContext() context.Context {
	return m.ctx
}

func (m *BaseMsg) GetHeader() *Header {
	return m.header
}

func (m *BaseMsg) SetHeader(h *Header) {
	m.header = h
}

func (m *BaseMsg) Dispatch() {

}

func (m *BaseMsg) Process() {
	panic("implement me")
}

func (m *BaseMsg) Execute() {
	m.Process()
}

type Header struct {
	TypeId int32
	PvId   int32
}

func (m *Header) Decode(src *bytes.Buffer) error {
	err := binary.Read(src, binary.BigEndian, &m.TypeId)
	if err != nil {
		return err
	}
	err = binary.Read(src, binary.BigEndian, &m.PvId)
	if err != nil {
		return err
	}
	return nil
}

func (m *Header) Encode(dst *bytes.Buffer) error {
	err := binary.Write(dst, binary.BigEndian, m.TypeId)
	if err != nil {
		return err
	}

	err = binary.Write(dst, binary.BigEndian, m.PvId)
	if err != nil {
		return err
	}

	return nil
}

//var taskQueues = make([]chan Task, 8)
//
//func HashExecute(session *network.Session, msg Msg) {
//	hash := fnv.New32a()
//	hash.Write([]byte(fmt.Sprintf("%v", session)))
//	idx := hash.Sum32() & (8 - 1)
//	taskQueues[idx] <- msg
//}
//
//func StartExecuteTasks() {
//	for i := 0; i < len(taskQueues); i++ {
//		queue := taskQueues[i]
//		go startExecuteTask(queue)
//	}
//}
//
//func startExecuteTask(ch chan Task) {
//	for {
//		select {
//		case msg := <-ch:
//			msg.Execute()
//		default:
//			continue
//		}
//	}
//}
