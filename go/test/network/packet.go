package network

import (
	"encoding/binary"
	"errors"
	"github.com/golang/protobuf/proto"
)

const HeaderSize = 8 // 4字节消息长度 + 4字节协议号

var PacketTooShortErr = errors.New("packet is too short")
var PacketNotCompleteErr = errors.New("packet is not complete")

// Pack 将消息打包
func Pack(protoID uint32, message proto.Message) ([]byte, error) {
	// 序列化protobuf消息
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	// 计算总长度
	totalLen := HeaderSize + len(data)

	// 创建完整的消息包
	packet := make([]byte, totalLen)

	// 写入消息长度（不包括长度字段本身）
	binary.BigEndian.PutUint32(packet[0:4], uint32(len(data)))

	// 写入协议号
	binary.BigEndian.PutUint32(packet[4:8], protoID)

	// 写入protobuf数据
	copy(packet[HeaderSize:], data)

	return packet, nil
}

// Unpack 解包消息
func Unpack(data []byte) (uint32, uint32, []byte, error) {
	if len(data) < HeaderSize {
		return 0, 0, nil, PacketTooShortErr
	}

	length := binary.BigEndian.Uint32(data[0:4])
	protoID := binary.BigEndian.Uint32(data[4:8])

	if len(data) < int(HeaderSize+length) {
		return 0, 0, nil, PacketNotCompleteErr
	}
	protoData := data[HeaderSize : HeaderSize+length]
	return length, protoID, protoData, nil
}
