package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	count := 100 * 10000
	start := time.Now()
	for i := 0; i < count; i++ {
		go func() {
			time.Sleep(10 * time.Second)
		}()
	}

	fmt.Printf("time: %v NumGoroutine: %d\n", time.Since(start), runtime.NumGoroutine())
	PrintMemUsage()
	for {
		n := runtime.NumGoroutine()
		if n < 2 {
			break
		}
		time.Sleep(time.Millisecond)
	}
	fmt.Printf("finish in %v\n", time.Since(start))
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %vMiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %vMiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %vMiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
