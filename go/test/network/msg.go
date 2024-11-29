package network

import (
	"errors"
)

type Msg interface {
	GetProtoID() uint32
	Decode(buffer []byte) error
}

var msgCreator = map[uint32]func() Msg{
	1: func() Msg { return &Echo{} },
}

func CreateMsg(protoId uint32, protoData []byte) (Msg, error) {
	create, exists := msgCreator[protoId]
	if !exists {
		return nil, errors.New("protoId not exists")
	}
	v := create()

	if msg, ok := v.(Msg); ok {
		err := msg.Decode(protoData)
		if err != nil {
			return nil, err
		}
		return msg, nil
	}
	return nil, errors.New("failed to convert to Msg")
}
