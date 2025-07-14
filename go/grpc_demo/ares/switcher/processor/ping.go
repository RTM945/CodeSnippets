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
	ping.GetSession().(*switcher.LinkerSession).ResetAlive()
	resp := msg.NewPong()
	resp.TypedPB().Serial = ping.TypedPB().Serial
	return ping.GetSession().Send(resp)
}
