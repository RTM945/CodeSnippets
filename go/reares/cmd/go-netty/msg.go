package shard

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
)

type Codec interface {
	Encode(dst *bytes.Buffer) error
	Decode(src *bytes.Buffer) error
}

type MsgHeader struct {
	TypeId int32
	PvId   int32
}

func (msgHeader *MsgHeader) Encode(dst *bytes.Buffer) error {
	err := binary.Write(dst, binary.BigEndian, msgHeader.TypeId)
	if err != nil {
		return err
	}

	err = binary.Write(dst, binary.BigEndian, msgHeader.PvId)
	if err != nil {
		return err
	}

	return nil
}

func (msgHeader *MsgHeader) Decode(src *bytes.Buffer) error {
	err := binary.Read(src, binary.BigEndian, &msgHeader.TypeId)
	if err != nil {
		return err
	}
	err = binary.Read(src, binary.BigEndian, &msgHeader.PvId)
	if err != nil {
		return err
	}
	return nil
}

type IMsg interface {
	Codec
	Process() error
	Dispatch()
	GetSession() Session
	SetSession(session Session)
	GetHeader() *MsgHeader
	SetHeader(header *MsgHeader)
	GetContext() context.Context
	SetContext(ctx context.Context)
}

type Msg struct {
	session   Session
	msgHeader *MsgHeader
	ctx       context.Context
}

func NewMsg(msgHeader *MsgHeader) *Msg {
	return &Msg{
		msgHeader: msgHeader,
	}
}

func (msg *Msg) SetSession(session Session) {
	msg.session = session
}

func (msg *Msg) GetSession() Session {
	return msg.session
}

func (msg *Msg) SetContext(ctx context.Context) {
	msg.ctx = ctx
}

func (msg *Msg) GetContext() context.Context {
	return msg.ctx
}

func (msg *Msg) GetHeader() *MsgHeader {
	return msg.msgHeader
}

func (msg *Msg) SetHeader(msgHeader *MsgHeader) {
	msg.msgHeader = msgHeader
}

func (msg *Msg) Dispatch() {

}

func (msg *Msg) Process() error {
	return nil
}

func (msg *Msg) Decode(*bytes.Buffer) error {
	return nil
}

func (msg *Msg) Encode(*bytes.Buffer) error {
	return nil
}

type MsgCreatorFunc[T IMsg] func() T

var MsgCreator = map[int32]MsgCreatorFunc[IMsg]{}

func CreateMsg(header *MsgHeader, session Session, buffer *bytes.Buffer) (IMsg, error) {
	create, exists := MsgCreator[header.TypeId]
	if !exists {
		return nil, fmt.Errorf("typeId %d not exists", header.TypeId)
	}
	msg := create()

	err := msg.Decode(buffer)
	if err != nil {
		return nil, err
	}
	msg.SetSession(session)
	msg.SetHeader(header)
	return msg, nil
}

//var taskQueues = make([]chan Task, 8)
//
//func HashExecute(session *network.Session, msg IMsg) {
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
