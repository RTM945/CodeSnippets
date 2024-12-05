package io

import (
	"bytes"
	"fmt"
)

var MsgCreator = map[int32]func(header *Header, session *Session) Msg{
	//1: func(header *Header, session *Session) Msg {
	//	return &proto.Echo{
	//		BaseMsg: NewBaseMsg(session, header),
	//	}
	//},
}

func CreateMsg(header *Header, session *Session, buffer *bytes.Buffer) (Msg, error) {
	create, exists := MsgCreator[header.TypeId]
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
