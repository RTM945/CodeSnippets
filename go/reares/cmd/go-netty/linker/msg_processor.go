package main

import (
	shard "reares/cmd/go-netty"
	client_switch "reares/cmd/go-netty/proto/client-switch"
)

type MsgProcessor struct{}

func Init() {
	processor := &MsgProcessor{}
	shard.MsgCreator[1] = func() shard.IMsg { return client_switch.InitRSAKeyExchange(processor) }
	shard.MsgCreator[2] = func() shard.IMsg { return client_switch.InitKeyExchange(processor) }
}

func (MsgProcessor) ProcessRSAKeyExchange(msg *client_switch.RSAKeyExchange) error {
	session := msg.GetSession().(*Session)
	return session.SendKeyExchange(msg.Key)
}

func (MsgProcessor) ProcessKeyExchange(msg *client_switch.KeyExchange) error {
	session := msg.GetSession().(*Session)
	return session.SetClientKey(msg.Key)
}
