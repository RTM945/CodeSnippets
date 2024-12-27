package client_switch

import (
	"bytes"
	"google.golang.org/protobuf/proto"
	shard "reares/cmd/go-netty"
	"reares/protobuf"
)

type RSAKeyExchange struct {
	*shard.MsgBase
	Processor
	Key []byte
}

func InitRSAKeyExchange(msgProcessor Processor) *RSAKeyExchange {
	header := &shard.MsgHeader{}
	header.TypeId = 1
	return &RSAKeyExchange{
		MsgBase:   shard.NewMsgBase(header),
		Processor: msgProcessor,
	}
}

func NewRSAKeyExchange() *RSAKeyExchange {
	header := &shard.MsgHeader{}
	header.TypeId = 1
	return &RSAKeyExchange{
		MsgBase: shard.NewMsgBase(header),
	}
}

func (msg *RSAKeyExchange) Decode(src *bytes.Buffer) error {
	tmp := &protobuf.RSAKeyExchange{}
	err := proto.Unmarshal(src.Bytes(), tmp)
	if err != nil {
		return err
	}
	msg.Key = tmp.Key
	return nil
}

func (msg *RSAKeyExchange) Encode(dst *bytes.Buffer) error {
	tmp := &protobuf.RSAKeyExchange{
		Key: msg.Key,
	}
	data, err := proto.Marshal(tmp)
	if err != nil {
		return err
	}
	_, err = dst.Write(data)
	return err
}

func (msg *RSAKeyExchange) Dispatch() {
	msg.Process()
}

func (msg *RSAKeyExchange) Process() error {
	return msg.ProcessRSAKeyExchange(msg)
}
