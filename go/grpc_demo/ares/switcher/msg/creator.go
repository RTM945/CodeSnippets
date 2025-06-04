package msg

import (
	"ares/pkg/io"
	pb "ares/proto/gen"
)

type CreatorFunc func(session io.Session, envelope *pb.Envelope) (io.Msg, error)

var Creator = map[uint32]CreatorFunc{
	4: PingCreator,
}
