package main

import (
	"ares/switcher"
	"ares/switcher/msg"
)

func main() {
	linker := switcher.GetLinker()
	linker.OnSessionError = func(session *switcher.LinkerSession, code uint32) error {
		sessionError := msg.NewSessionError()
		sessionError.TypedPB().Code = code
		err := session.Send0(sessionError)
		session.Close()
		return err
	}
	linker.OnServerError = func(session *switcher.LinkerSession, pvId, code uint32) error {
		serverError := msg.NewServerError()
		serverError.TypedPB().PvId = pvId
		serverError.TypedPB().Code = code
		return session.Send0(serverError)
	}
	linker.OnDispatch = func(session *switcher.ProviderSession, pvId, typeId uint32, payload []byte) error {
		dispatch := msg.NewDispatch()
		dispatch.TypedPB().PvId = pvId
		dispatch.TypedPB().TypeId = typeId
		dispatch.TypedPB().Payload = payload
		return session.Send(dispatch)
	}
}
