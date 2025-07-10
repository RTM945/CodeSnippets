package main

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	switchermsg "ares/switcher/msg"
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"

	_ "ares/proto/gen"
)

var TypeIDMap = map[uint32]proto.Message{}
var typeIDMap = map[protoreflect.FullName]uint32{}

func init() {
	protoregistry.GlobalTypes.RangeMessages(func(md protoreflect.MessageType) bool {
		desc := md.Descriptor()
		opts := desc.Options().(*descriptorpb.MessageOptions)
		raw := proto.GetExtension(opts, pb.E_TypeId)
		typeID, ok := raw.(uint32)
		if !ok {
			fmt.Printf("unexpected type_id type for %s: %T\n", desc.FullName(), raw)
			return true
		}
		if typeID == 0 {
			return true
		}
		msg := md.New().Interface()
		TypeIDMap[typeID] = msg
		typeIDMap[desc.FullName()] = typeID
		fmt.Printf("Registered message %s with type_id %d\n", desc.FullName(), typeID)
		return true
	})

}

func Init(node ares.INode) {
	node.MsgCreator().Register(4, switchermsg.PingCreator)
	node.MsgCreator().Register(51, switchermsg.DispatchCreator)
	node.MsgCreator().Register(53, switchermsg.ProvideeKickCreator)
}
