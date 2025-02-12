package echo

type Processor interface {
	ProcessCEcho(msg *CEcho) error
	ProcessSEcho(msg *SEcho) error
}
