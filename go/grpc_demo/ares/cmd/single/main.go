package main

import (
	"time"
)

func main() {
	etcdClient, err := etcdInit()
	if err != nil {
		return
	}

	go func() {
		linker.Start(etcdClient)
	}()
	go func() {
		provider.Start(etcdClient)
	}()
	time.Sleep(5 * time.Second)
	go func() {
		NewProvidee(5, 501, 2501).Start(etcdClient)
	}()
	go func() {
		NewProvidee(5, 502, 2502).Start(etcdClient)
	}()

	select {}
}
