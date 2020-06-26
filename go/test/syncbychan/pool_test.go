package syncbychan

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"runtime/debug"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	// disable GC so we can control when it happens.
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var p sync.Pool
	if p.Get() != nil {
		t.Fatal("expected empty")
	}

	p.Put("a")
	p.Put("b")
	if g := p.Get(); g != "a" {
		t.Fatalf("got %#v; want a", g)
	}
	if g := p.Get(); g != "b" {
		t.Fatalf("got %#v; want b", g)
	}
	if g := p.Get(); g != nil {
		t.Fatalf("got %#v; want nil", g)
	}
	// Put in a large number of objects so they spill into
	// stealable space.
	for i := 0; i < 100; i++ {
		p.Put("c")
	}
	// After one GC, the victim cache should keep them alive.
	runtime.GC()
	if g := p.Get(); g != "c" {
		t.Fatalf("got %#v; want c after GC", g)
	}
	// A second GC should drop the victim cache.
	runtime.GC()
	if g := p.Get(); g != nil {
		t.Fatalf("got %#v; want nil after second GC", g)
	}
}

func TestPoolUse(t *testing.T) {
	var m sync.Mutex
	holder := make(map[string]bool)

	pool := sync.Pool{
		New: func() interface{} {
			buffer := make([]byte, 1024)
			return &buffer
		},
	}

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get("http://rtmsoft.me")
			if err != nil {
				t.Error(err)
			}
			defer resp.Body.Close()
			tmp := pool.Get().(*[]byte)
			key := fmt.Sprintf("%p", tmp)
			m.Lock()
			_, ok := holder[key]
			if !ok {
				holder[key] = true
			}
			m.Unlock()
			_, err = io.ReadFull(resp.Body, *tmp)
			if err != nil {
				t.Error(err)
			}
			pool.Put(tmp)
		}()
	}
	wg.Wait()
	// for key, val := range holder {
	// 	fmt.Println("Key:", key, "Value:", val)
	// }
	// 为了测试池中临时对象的复用
	// 本次测试输出14 小于20
	fmt.Println(len(holder))
}

func TestChanPool(t *testing.T) {
	i := 0
	p := NewPool(1024,
		// func() interface{} {
		// 	return make([]byte, 0, 10)
		// },
		// func(i interface{}) interface{} {
		// 	return i.([]byte)[:0]
		// }, //这个没什么测头...
		func() interface{} {
			i++
			return i
		},
		nil,
	)
	if v := p.Get(); v != 1 {
		t.Fatalf("got %v; want 1", v)
	}
	if v := p.Get(); v != 2 {
		t.Fatalf("got %v; want 2", v)
	}
}
