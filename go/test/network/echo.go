package proto

import (
	"bytes"
	"google.golang.org/protobuf/proto"
	"reares/internal/io"
	"reares/protobuf"
)

type Echo struct {
	*io.BaseMsg
	Msg string
}

func NewEcho() *Echo {
	header := &io.MsgHeader{}
	header.TypeId = 1
	return &Echo{
		BaseMsg: io.NewBaseMsg(header),
	}
}

func (msg *Echo) Decode(src *bytes.Buffer) error {
	tmp := &protobuf.Echo{}
	err := proto.Unmarshal(src.Bytes(), tmp)
	if err != nil {
		return err
	}
	msg.Msg = tmp.Msg
	return nil
}

func (msg *Echo) Encode(dst *bytes.Buffer) error {
	data, err := proto.Marshal(&protobuf.Echo{Msg: msg.Msg})
	dst.Write(data)
	return err
}

func (msg *Echo) Process() {
	panic("implement me")
}
