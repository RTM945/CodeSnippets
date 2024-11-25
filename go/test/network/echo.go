package network

import (
	"github.com/golang/protobuf/proto"
	"gotest/network/protobuf"
)

type Echo struct {
	msg string
}

func (e *Echo) Decode(buffer []byte) error {
	tmp := &protobuf.Echo{}
	err := proto.Unmarshal(buffer, tmp)
	if err != nil {
		return err
	}
	e.msg = tmp.Msg
	return nil
}

func (e *Echo) GetProtoID() uint32 {
	return 1
}
