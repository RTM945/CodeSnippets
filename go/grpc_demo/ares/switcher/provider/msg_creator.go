package provider

import (
	"ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/msg"
)

type MsgCreatorFunc func(session io.Session, envelope *pb.Envelope) (io.Msg, error)

var MsgCreator = map[uint32]MsgCreatorFunc{
	73: msg.SendToClientCreator,
}
