package io

type INode interface {
	Sessions() ISessions
	MsgCreator() IMsgCreator
}
