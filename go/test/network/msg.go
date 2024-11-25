package network

import (
	"errors"
	"reflect"
)

type Msg interface {
	GetProtoID() uint32
	Decode(buffer []byte) error
}

var protoTypes = map[uint32]reflect.Type{
	1: reflect.TypeOf(&Echo{}),
}

type MsgCreator func(protoData []byte) Msg

func CreateMsg(protoId uint32, protoData []byte) (Msg, error) {
	protoType, exists := protoTypes[protoId]
	if !exists {
		return nil, errors.New("protoId not exists")
	}
	v := reflect.New(protoType.Elem()).Interface()

	if msg, ok := v.(Msg); ok {
		err := msg.Decode(protoData)
		if err != nil {
			return nil, err
		}
		return msg, nil
	}
	return nil, errors.New("failed to convert to Msg")
}
