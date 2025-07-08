package main

import (
	ares "ares/pkg/io"
	switchermsg "ares/switcher/msg"
	switcherprocessor "ares/switcher/processor"
)

//func init() {
//	switchermsg.DispatcherProcessor = switcherprocessor.Dispatch
//	switchermsg.PingProcessor = switcherprocessor.Ping
//	switchermsg.ProvideeKickProcessor = switcherprocessor.ProvideeKick
//}

func Init(node ares.INode) {
	node.MsgCreator().Register(4, switchermsg.PingCreator)
	node.MsgCreator().Register(51, switchermsg.DispatchCreator)
	node.MsgCreator().Register(53, switchermsg.ProvideeKickCreator)

	ares.RegisterMsgProcessor[*switchermsg.Ping](node.MsgProcessor(), 4, switcherprocessor.Ping)
}
