package syncbychan

import (
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	var m sync.Mutex
	s := 0
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			s += i
			wg.Done()
		}(i)
	}
	wg.Wait()
	if got, want := s, (1+100)*100/2; got != want {
		t.Errorf("get %d, want %d", got, want)
	}
}

func TestChanMutex(t *testing.T) {
	m := NewMutex()
	s := 0
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			s += i
			wg.Done()
		}(i)
	}
	wg.Wait()
	if got, want := s, (1+100)*100/2; got != want {
		t.Errorf("get %d, want %d", got, want)
	}

}
