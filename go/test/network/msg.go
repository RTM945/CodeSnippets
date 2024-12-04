package network

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"hash/fnv"
)

type Coder interface {
	Decode(buffer *bytes.Buffer) error
	Encode() ([]byte, error)
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
	GetHeader() *MsgHeader
	SetHeader(h *MsgHeader)
	GetContext() context.Context
	SetContext(ctx context.Context)
}

type BaseMsg struct {
	session *Session
	header  *MsgHeader
	ctx     context.Context
}

func NewBaseMsg(session *Session, header *MsgHeader) *BaseMsg {
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

func (m *BaseMsg) GetHeader() *MsgHeader {
	return m.header
}

func (m *BaseMsg) SetHeader(h *MsgHeader) {
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

type MsgHeader struct {
	TypeId int32
	PvId   int32
}

func (m *MsgHeader) Decode(buffer *bytes.Buffer) error {
	err := binary.Read(buffer, binary.BigEndian, &m.TypeId)
	if err != nil {
		return err
	}
	err = binary.Read(buffer, binary.BigEndian, &m.PvId)
	if err != nil {
		return err
	}
	return nil
}

func (m *MsgHeader) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, m.TypeId)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, m.PvId)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var msgCreator = map[int32]func(header *MsgHeader, session *Session) Msg{
	1: func(header *MsgHeader, session *Session) Msg {
		return &Echo{
			BaseMsg: NewBaseMsg(session, header),
		}
	},
}

var taskQueues = make([]chan Task, 8)

func HashExecute(session *Session, msg Msg) {
	hash := fnv.New32a()
	hash.Write([]byte(fmt.Sprintf("%v", session)))
	idx := hash.Sum32() & (8 - 1)
	taskQueues[idx] <- msg
}

func StartExecuteTasks() {
	for i := 0; i < len(taskQueues); i++ {
		queue := taskQueues[i]
		go startExecuteTask(queue)
	}
}

func startExecuteTask(ch chan Task) {
	for {
		select {
		case msg := <-ch:
			msg.Execute()
		default:
			continue
		}
	}
}

func CreateMsg(header *MsgHeader, session *Session, buffer *bytes.Buffer) (Msg, error) {
	create, exists := msgCreator[header.TypeId]
	if !exists {
		return nil, fmt.Errorf("typeId %d not exists", header.TypeId)
	}
	v := create(header, session)

	if msg, ok := v.(Msg); ok {
		err := msg.Decode(buffer)
		if err != nil {
			return nil, err
		}
		return msg, nil
	}
	return nil, fmt.Errorf("failed to convert to Msg typeId %d", header.TypeId)
}
