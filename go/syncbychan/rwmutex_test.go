package syncbychan

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestRWMutex(t *testing.T) {
	count := 0
	mutex := NewRWMutex()
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			mutex.Lock()
			count += i
			mutex.UnLock()
			wg.Done()
		}(i)
	}

	go func() {
		for {
			mutex.RLock()
			t.Logf("count = %d\n", count)
			mutex.RUnlock()
			time.Sleep(10 * time.Millisecond)
		}
	}()

	wg.Wait()
	if got, want := count, (100+1)*100/2; got != want {
		t.Errorf("s = %d, want %d", got, want)
	}
}
