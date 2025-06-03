package main

import (
	linkerpb "ares/proto/switcher"
	"fmt"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	ping := linkerpb.Ping{Serial: 1}
	a, _ := anypb.New(&ping)
	fmt.Println(a.TypeUrl)
}
