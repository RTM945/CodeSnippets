package proto

import (
	"bytes"
	"reares/internal/io"
	"reares/protobuf"

	"google.golang.org/protobuf/proto"
)

type Echo struct {
	*io.MsgBase
	Msg string
}

func NewEcho() *Echo {
	header := &io.MsgHeader{}
	header.TypeId = 1
	return &Echo{
		MsgBase: io.NewMsgBase(header),
	}
}

func (echo *Echo) Decode(src *bytes.Buffer) error {
	tmp := &protobuf.Echo{}
	err := proto.Unmarshal(src.Bytes(), tmp)
	if err != nil {
		return err
	}
	echo.Msg = tmp.Msg
	return nil
}

func (echo *Echo) Encode(dst *bytes.Buffer) error {
	data, err := proto.Marshal(&protobuf.Echo{Msg: echo.Msg})
	if err != nil {
		return err
	}
	_, err = dst.Write(data)
	return err
}

func (echo *Echo) Dispatch() {
	echo.Process()
}

func (echo *Echo) Process() error {
	return io.MsgProcessor[echo.GetHeader().TypeId](echo)
}
