package io

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

type MsgBase struct {
	session   *Session
	msgHeader *MsgHeader
	ctx       context.Context
}

func NewMsgBase(msgHeader *MsgHeader) *MsgBase {
	return &MsgBase{
		msgHeader: msgHeader,
	}
}

func (msgBase *MsgBase) SetSession(session *Session) {
	msgBase.session = session
}

func (msgBase *MsgBase) GetSession() *Session {
	return msgBase.session
}

func (msgBase *MsgBase) SetContext(ctx context.Context) {
	msgBase.ctx = ctx
}

func (msgBase *MsgBase) GetContext() context.Context {
	return msgBase.ctx
}

func (msgBase *MsgBase) GetHeader() *MsgHeader {
	return msgBase.msgHeader
}

func (msgBase *MsgBase) SetHeader(msgHeader *MsgHeader) {
	msgBase.msgHeader = msgHeader
}

func (msgBase *MsgBase) Dispatch() {

}

type Msg interface {
	Codec
	Process() error
	Dispatch()
	GetSession() *Session
	SetSession(session *Session)
	GetHeader() *MsgHeader
	SetHeader(header *MsgHeader)
	GetContext() context.Context
	SetContext(ctx context.Context)
}

type MsgCreatorFunc[T Msg] func() T

var MsgCreator = map[int32]MsgCreatorFunc[Msg]{}

func CreateMsg(header *MsgHeader, session *Session, buffer *bytes.Buffer) (Msg, error) {
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
