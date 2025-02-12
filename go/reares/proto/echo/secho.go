package echo

import (
	"bytes"
	"google.golang.org/protobuf/proto"
	shard "reares/cmd/go-netty"
	"reares/protobuf"
)

type SEcho struct {
	*shard.Msg
	Processor
	Msg string
}

func InitSEcho(msgProcessor Processor) *SEcho {
	header := &shard.MsgHeader{}
	header.TypeId = 101
	return &SEcho{
		Msg:       shard.NewMsg(header),
		Processor: msgProcessor,
	}
}

func NewSEcho() *SEcho {
	header := &shard.MsgHeader{}
	header.TypeId = 101
	return &SEcho{
		Msg: shard.NewMsg(header),
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
