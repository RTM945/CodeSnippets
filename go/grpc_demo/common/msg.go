package common

import (
	"context"
	"google.golang.org/protobuf/proto"
)

type Msg interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	GetPB() proto.Message
	GetType() uint32
	GetPvId() uint32
	GetContext() context.Context
	Dispatch() error
	Process() error
}
