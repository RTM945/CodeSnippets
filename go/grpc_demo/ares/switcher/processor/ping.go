package processor

import (
	"ares/switcher"
	"ares/switcher/msg"
)

type PingProcessor struct {
}

func NewPingProcessor() *PingProcessor {
	return &PingProcessor{}
}

func (p *PingProcessor) Process(ping *msg.Ping) error {
	linkerSession := ping.GetSession().(*switcher.LinkerSession)
	linkerSession.ResetAlive()
	resp := msg.NewPong()
	resp.TypedPB().Serial = ping.TypedPB().Serial
	return linkerSession.Send(resp)
}
