package io

import (
	"google.golang.org/protobuf/proto"
)

type Msg interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	GetPB() proto.Message
	GetType() string
	GetPvId() uint32
	GetContext() any
	Dispatch()
	Process() error
	SetSession(session Session)
	GetSession() Session
}
