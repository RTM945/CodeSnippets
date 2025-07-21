package io

type INode interface {
	Sessions() ISessions
	MsgCreator() IMsgCreator
	MsgProcessor() IMsgProcessor
	//Ports() map[string]IPort
}
