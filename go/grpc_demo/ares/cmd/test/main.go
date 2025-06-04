package main

import (
	pb "ares/proto/gen"
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/anypb"
	"os"
)

func main() {
	ping := pb.Ping{Serial: 1}
	a, _ := anypb.New(&ping)
	fmt.Println(a.TypeUrl)

	data, err := os.ReadFile("proto/gen/protos.desc")
	if err != nil {
		panic(err)
	}
	fds := &descriptorpb.FileDescriptorSet{}
	if err := proto.Unmarshal(data, fds); err != nil {
		panic(err)
	}
	files, err := protodesc.NewFiles(fds)
	if err != nil {
		panic(err)
	}
	files.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		messages := fd.Messages()
		for i := 0; i < messages.Len(); i++ {
			md := messages.Get(i)
			opts := md.Options().(*descriptorpb.MessageOptions)
			if opts != nil {
				ext := proto.GetExtension(opts, pb.E_TypeId)
				if typeId, ok := ext.(uint32); ok {
					fmt.Printf("name : %s typeID: %d\n", md.FullName(), typeId)
				}
			}

		}
		return true
	})
}
