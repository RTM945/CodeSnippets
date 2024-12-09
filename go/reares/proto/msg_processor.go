package proto

type EchoProcessor interface {
	ProcessCEcho(cEcho *CEcho) error
	ProcessSEcho(sEcho *SEcho) error
}
