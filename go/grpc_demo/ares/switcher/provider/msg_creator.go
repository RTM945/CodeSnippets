package provider

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/msg"
)

type MsgCreatorFunc func(session ares.ISession, envelope *pb.Envelope) (ares.IMsg, error)

var MsgCreator = map[uint32]MsgCreatorFunc{
	73: msg.SendToClientCreator,
}
