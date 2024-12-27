package echo

import (
	"reares/cmd/go-netty/proto/echo"
)

type MsgProcessor struct {
}

func NewMsgProcessor() *MsgProcessor {
	return &MsgProcessor{}
}

func (p *MsgProcessor) ProcessCEcho(echo *echo.CEcho) error {
	return GetEchoLogic(echo.GetSession()).Echo(echo.Msg)
}

func (p *MsgProcessor) ProcessSEcho(echo *echo.SEcho) error {
	return GetEchoLogic(echo.GetSession()).TestSEcho(echo.Msg)
}
