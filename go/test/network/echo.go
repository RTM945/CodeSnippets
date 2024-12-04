package network

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"gotest/network/protobuf"
)

type Echo struct {
	*BaseMsg
	msg string
}

func (e *Echo) Decode(buffer *bytes.Buffer) error {
	tmp := &protobuf.Echo{}
	err := proto.Unmarshal(buffer.Bytes(), tmp)
	if err != nil {
		return err
	}
	e.msg = tmp.Msg
	return nil
}

func (e *Echo) Encode() ([]byte, error) {
	buffer, err := proto.Marshal(&protobuf.Echo{Msg: e.msg})
	return buffer, err
}

func (e *Echo) Process() {

}
