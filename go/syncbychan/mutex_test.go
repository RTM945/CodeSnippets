package syncbychan

import (
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	mutex := NewMutex()
	// var mutex sync.Mutex
	var wg sync.WaitGroup
	s := 0
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func(i int) {
			mutex.Lock()
			s += i
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	if got, want := s, (1000+1)*1000/2; got != want {
		t.Errorf("s = %d, want %d", got, want)
	}
}
