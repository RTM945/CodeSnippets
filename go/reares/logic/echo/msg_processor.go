package echo

import (
	"reares/proto/echo"
)

type MsgProcessor struct {
}

func NewMsgProcessor() *MsgProcessor {
	return &MsgProcessor{}
}

func (p *MsgProcessor) ProcessCEcho(echo *proto.CEcho) error {
	return GetEchoLogic(echo.GetSession()).Echo(echo.Msg)
}

func (p *MsgProcessor) ProcessSEcho(echo *proto.SEcho) error {
	return GetEchoLogic(echo.GetSession()).TestSEcho(echo.Msg)
}
