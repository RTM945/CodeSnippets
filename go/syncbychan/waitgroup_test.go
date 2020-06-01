package syncbychan

import (
	"fmt"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	wg := NewWaitGroup()
	mutex := NewMutex()
	s := 0
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			mutex.Lock()
			s += i
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(s)
}
