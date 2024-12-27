package echo

import (
	"bytes"
	"google.golang.org/protobuf/proto"
	shard "reares/cmd/go-netty"
	"reares/protobuf"
)

type CEcho struct {
	*shard.MsgBase
	Processor
	Msg string
}

func InitCEcho(msgProcessor Processor) *CEcho {
	header := &shard.MsgHeader{}
	header.TypeId = 100
	return &CEcho{
		MsgBase:   shard.NewMsgBase(header),
		Processor: msgProcessor,
	}
}

func NewCEcho() *CEcho {
	header := &shard.MsgHeader{}
	header.TypeId = 100
	return &CEcho{
		MsgBase: shard.NewMsgBase(header),
	}
}

func (echo *CEcho) Decode(src *bytes.Buffer) error {
	tmp := &protobuf.Echo{}
	err := proto.Unmarshal(src.Bytes(), tmp)
	if err != nil {
		return err
	}
	echo.Msg = tmp.Msg
	return nil
}

func (echo *CEcho) Encode(dst *bytes.Buffer) error {
	data, err := proto.Marshal(&protobuf.Echo{Msg: echo.Msg})
	if err != nil {
		return err
	}
	_, err = dst.Write(data)
	return err
}

func (echo *CEcho) Dispatch() {
	echo.Process()
}

func (echo *CEcho) Process() error {
	return echo.ProcessCEcho(echo)
}
