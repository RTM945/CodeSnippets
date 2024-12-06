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

type Msg interface {
	Coder
	Process() error
	Dispatch()
	GetSession() *Session
	SetSession(session *Session)
	GetHeader() *MsgHeader
	SetHeader(h *MsgHeader)
	GetContext() context.Context
	SetContext(ctx context.Context)
}

type MsgBase struct {
	session *Session
	header  *MsgHeader
	ctx     context.Context
}

func NewMsgBase(header *MsgHeader) *MsgBase {
	return &MsgBase{
		header: header,
	}
}

func (m *MsgBase) SetSession(session *Session) {
	m.session = session
}

func (m *MsgBase) GetSession() *Session {
	return m.session
}

func (m *MsgBase) SetContext(ctx context.Context) {
	m.ctx = ctx
}

func (m *MsgBase) GetContext() context.Context {
	return m.ctx
}

func (m *MsgBase) GetHeader() *MsgHeader {
	return m.header
}

func (m *MsgBase) SetHeader(h *MsgHeader) {
	m.header = h
}

type MsgHeader struct {
	TypeId int32
	PvId   int32
}

func (m *MsgHeader) Decode(src *bytes.Buffer) error {
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

func (m *MsgHeader) Encode(dst *bytes.Buffer) error {
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
