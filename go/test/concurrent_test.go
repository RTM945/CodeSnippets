package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestNoLock(t *testing.T) {
	a := 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a++
		}()
	}
	wg.Wait()
	t.Log(a)
}

func TestLock(t *testing.T) {
	a := 0
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			a++
		}()
	}
	wg.Wait()
	t.Log(a)
}

func TestAtomic(t *testing.T) {
	var a int32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&a, 1)
		}()
	}
	wg.Wait()
	t.Log(a)
}

func TestChanCommunicate(t *testing.T) {
	c := make(chan int, 1)
	a := 0
	c <- a
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			a := <-c
			a++
			c <- a
		}(c)
	}
	wg.Wait()
	t.Log(<-c)
}
