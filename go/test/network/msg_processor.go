package io

type MsgProcessorFunc[T Msg] func(msg T) error

var MsgProcessor = map[int32]MsgProcessorFunc[Msg]{}
