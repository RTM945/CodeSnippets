package processor

import (
	"ares/switcher"
	"ares/switcher/msg"
)

func Ping(ping *msg.Ping) error {
	ping.GetSession().(*switcher.LinkerSession).ResetAlive()
	resp := msg.NewPong()
	resp.TypedPB().Serial = ping.TypedPB().Serial
	return ping.GetSession().Send(resp)
}
