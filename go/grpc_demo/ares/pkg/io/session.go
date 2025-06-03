package io

type Session interface {
	GetSid() uint32
	Send(Msg) error
	Process(Msg)
	OnClose()
}
