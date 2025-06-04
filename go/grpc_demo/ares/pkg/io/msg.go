package io

import (
	"google.golang.org/protobuf/proto"
)

type Msg interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	GetPB() proto.Message
	GetType() uint32
	GetPvId() uint32
	SetPvId(pvId uint32)
	GetContext() any
	SetContext(context any)
	Dispatch()
	Process() error
	SetSession(session Session)
	GetSession() Session
}
