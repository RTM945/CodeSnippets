package main

import (
	_ "ares/proto/gen"
	pb "ares/proto/gen"
	"google.golang.org/protobuf/proto"
)

//var TypeIDMap = map[uint32]proto.Message{}
//var typeIDMap = map[protoreflect.FullName]uint32{}

//func init() {
//	protoregistry.GlobalTypes.RangeMessages(func(md protoreflect.MessageType) bool {
//		desc := md.Descriptor()
//		opts := desc.Options().(*descriptorpb.MessageOptions)
//		raw := proto.GetExtension(opts, pb.E_TypeId)
//		typeID, ok := raw.(uint32)
//		if !ok {
//			fmt.Printf("unexpected type_id type for %s: %T\n", desc.FullName(), raw)
//			return true
//		}
//		if typeID == 0 {
//			return true
//		}
//		msg := md.New().Interface()
//		TypeIDMap[typeID] = msg
//		typeIDMap[desc.FullName()] = typeID
//		fmt.Printf("Registered message %s with type_id %d\n", desc.FullName(), typeID)
//		return true
//	})
//	// TODO 用反射获取到了 type_id 和 pb 对象 下面的 creator 注册 可以尝试用反射生成
//}
//
//func Init(node ares.INode) {
//	node.MsgCreator().Register(4, switchermsg.PingCreator)
//	node.MsgCreator().Register(51, switchermsg.DispatchCreator)
//	node.MsgCreator().Register(53, switchermsg.ProvideeKickCreator)
//	node.MsgCreator().Register(73, switchermsg.SendToClientCreator)
//
//	node.MsgProcessor().Register(4, ares.NewTypedMsgProcessor[*switchermsg.Ping](switcherprocessor.NewPingProcessor()))
//	node.MsgProcessor().Register(51, ares.NewTypedMsgProcessor[*switchermsg.Dispatch](switcherprocessor.NewDispatchProcessor()))
//	node.MsgProcessor().Register(53, ares.NewTypedMsgProcessor[*switchermsg.ProvideeKick](switcherprocessor.NewProvideeKickProcessor()))
//	node.MsgProcessor().Register(73, ares.NewTypedMsgProcessor[*switchermsg.SendToClient](switcherprocessor.NewSendToClientProcessor()))
//
//}

func init() {
	provider.msgProcessor[52] = func(session *ProviderSession, msg proto.Message) error {
		bindPvId := msg.(*pb.BindPvId)
		session.info = bindPvId.Info
		session.checkToProvidee = bindPvId.CheckToProvidee
		provider.sessions[bindPvId.Info.PvId] = session
		if bindPvId.Info.ServerType == uint32(pb.ServerType_AU) {
			provider.auPvIds[bindPvId.Info.PvId] = struct{}{}
		} else if bindPvId.Info.ServerType == uint32(pb.ServerType_PHANTOM) {
			// 将linker ip端口和连接数 provider ip端口信息都发送给Phantom
			// 还需要将这个provider管理的其他providee信息发送给Phantom
			// 但服务发现功能由Phantom转给etcd的情况下还需要这样做吗?
		} else if bindPvId.Info.PvId == uint32(pb.ServerType_LOGIC) {
			// 游戏服原来的处理是将每个linker和该游戏服组装个linker-gs信息的对象发给Phantom
			// 方便客户端进行选服
			// 同样在使用etcd的情况下这个做法要不要沿用呢
		}

		// 通知所有Phantom provideeBind 换成etcd后还需要吗

		return nil
	}
}
