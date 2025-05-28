package io

type Session interface {
	GetSid() string
	Send(Msg) error
	GetExecutor()
}
