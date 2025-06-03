package msg

import "ares/pkg/io"

var Creator = map[string]func() io.Msg{
	"type.googleapis.com/switcher.Ping": func() io.Msg { return NewPing() },
}
