package client_switch

import (
	"bytes"
	"google.golang.org/protobuf/proto"
	"reares/internal/io"
	"reares/protobuf"
)

type KeyExchange struct {
	*io.MsgBase
	Processor
	Key []byte
}

func InitKeyExchange(msgProcessor Processor) *KeyExchange {
	header := &io.MsgHeader{}
	header.TypeId = 2
	return &KeyExchange{
		MsgBase:   io.NewMsgBase(header),
		Processor: msgProcessor,
	}
}

func NewKeyExchange() *KeyExchange {
	header := &io.MsgHeader{}
	header.TypeId = 2
	return &KeyExchange{
		MsgBase: io.NewMsgBase(header),
	}
}

func (msg *KeyExchange) Decode(src *bytes.Buffer) error {
	tmp := &protobuf.KeyExchange{}
	err := proto.Unmarshal(src.Bytes(), tmp)
	if err != nil {
		return err
	}
	msg.Key = tmp.Key
	return nil
}

func (msg *KeyExchange) Encode(dst *bytes.Buffer) error {
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

func (msg *KeyExchange) Dispatch() {
	msg.Process()
}

func (msg *KeyExchange) Process() error {
	return msg.ProcessKeyExchange(msg)
}
