package proto

import (
	"bytes"
	"google.golang.org/protobuf/proto"
	"reares/internal/io"
	"reares/protobuf"
)

type SEcho struct {
	*io.MsgBase
	EchoProcessor
	Msg string
}

func NewSEchoWithProcessor(msgProcessor EchoProcessor) *SEcho {
	header := &io.MsgHeader{}
	header.TypeId = 2
	return &SEcho{
		MsgBase:       io.NewMsgBase(header),
		EchoProcessor: msgProcessor,
	}
}

func NewSEcho() *SEcho {
	header := &io.MsgHeader{}
	header.TypeId = 2
	return &SEcho{
		MsgBase: io.NewMsgBase(header),
	}
}

func (echo *SEcho) Decode(src *bytes.Buffer) error {
	tmp := &protobuf.Echo{}
	err := proto.Unmarshal(src.Bytes(), tmp)
	if err != nil {
		return err
	}
	echo.Msg = tmp.Msg
	return nil
}

func (echo *SEcho) Encode(dst *bytes.Buffer) error {
	data, err := proto.Marshal(&protobuf.Echo{Msg: echo.Msg})
	if err != nil {
		return err
	}
	_, err = dst.Write(data)
	return err
}

func (echo *SEcho) Dispatch() {
	echo.Process()
}

func (echo *SEcho) Process() error {
	return echo.ProcessSEcho(echo)
}
