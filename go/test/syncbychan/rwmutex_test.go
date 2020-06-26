package syncbychan

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestRWMutex(t *testing.T) {
	var m sync.RWMutex
	var activity int32 = 0
	done := make(chan bool)
	readers := 10
	go writer(&m, &activity, done)
	var i int
	for i = 0; i < readers/2; i++ {
		go reader(&m, &activity, done)
	}
	go writer(&m, &activity, done)
	for ; i < readers; i++ {
		go reader(&m, &activity, done)
	}
	// 阻塞程序结束
	for i := 0; i < 2+readers; i++ {
		<-done
	}
}

func writer(m *sync.RWMutex, activity *int32, done chan bool) {
	m.Lock()
	n := atomic.AddInt32(activity, 10000)
	if n != 10000 {
		panic(fmt.Sprintf("wlock(%d)\n", n))
	}
	atomic.AddInt32(activity, -10000)
	m.Unlock()
	done <- true
}

func reader(m *sync.RWMutex, activity *int32, done chan bool) {
	m.RLock()
	n := atomic.AddInt32(activity, 1)
	if n < 1 || n >= 10000 {
		panic(fmt.Sprintf("wlock(%d)\n", n))
	}
	atomic.AddInt32(activity, -1)
	m.RUnlock()
	done <- true
}

func TestChanRWMutex(t *testing.T) {
	m := NewRWMutex()
	var activity int32 = 0
	done := make(chan bool)
	readers := 10
	go func(m *RWMutex, activity *int32, done chan bool) {
		m.Lock()
		n := atomic.AddInt32(activity, 10000)
		if n != 10000 {
			panic(fmt.Sprintf("wlock(%d)\n", n))
		}
		atomic.AddInt32(activity, -10000)
		m.Unlock()
		done <- true
	}(m, &activity, done)
	for i := 0; i < readers; i++ {
		go func(m *RWMutex, activity *int32, done chan bool) {
			m.RLock()
			n := atomic.AddInt32(activity, 1)
			if n < 1 || n >= 10000 {
				panic(fmt.Sprintf("wlock(%d)\n", n))
			}
			atomic.AddInt32(activity, -1)
			m.RUnlock()
			done <- true
		}(m, &activity, done)
	}
	// 阻塞程序结束
	for i := 0; i < 1+readers; i++ {
		<-done
	}
}
