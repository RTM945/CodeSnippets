package network

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
)

type Msg interface {
	Decode(buffer *bytes.Buffer) error
	Encode() ([]byte, error)
	Init(header *MsgHeader, session *Session)
}

type MsgBase struct {
	session *Session
	header  *MsgHeader
	ctx     context.Context
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

var msgCreator = map[int32]func() Msg{
	1: func() Msg { return &Echo{} },
}

func CreateMsg(header *MsgHeader, session *Session, buffer *bytes.Buffer) (Msg, error) {
	create, exists := msgCreator[header.TypeId]
	if !exists {
		return nil, fmt.Errorf("typeId %d not exists", header.TypeId)
	}
	v := create()

	if msg, ok := v.(Msg); ok {
		msg.Init(header, session)
		err := msg.Decode(buffer)
		if err != nil {
			return nil, err
		}
		return msg, nil
	}
	return nil, fmt.Errorf("failed to convert to Msg typeId %d", header.TypeId)
}
