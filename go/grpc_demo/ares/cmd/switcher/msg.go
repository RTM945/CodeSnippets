package main

import (
	switchermsg "ares/msg/switcher"
	switcherprocessor "ares/switcher/processor"
)

func init() {
	switchermsg.DispatcherProcessor = switcherprocessor.Dispatch
	switchermsg.PingProcessor = switcherprocessor.Ping
	switchermsg.ProvideeKickProcessor = switcherprocessor.ProvideeKick
}
