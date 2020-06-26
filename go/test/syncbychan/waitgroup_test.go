package syncbychan

import (
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	// Run the same test a few times to ensure barrier is in a proper state.
	for i := 0; i != 8; i++ {
		n := 16
		wg1.Add(n)
		wg2.Add(n)
		exited := make(chan bool, n)
		for i := 0; i != n; i++ {
			go func() {
				wg1.Done()
				wg2.Wait()
				exited <- true
			}()
		}
		wg1.Wait()
		for i := 0; i != n; i++ {
			select {
			case <-exited:
				t.Fatal("WaitGroup released group too soon")
			default:
			}
			wg2.Done()
		}
		for i := 0; i != n; i++ {
			<-exited // Will block if barrier fails to unlock someone.
		}
	}
}

func TestChanWaitGroup(t *testing.T) {
	wg1 := NewWaitGroup()
	wg2 := NewWaitGroup()
	// Run the same test a few times to ensure barrier is in a proper state.
	for i := 0; i != 8; i++ {
		n := 16
		wg1.Add(n)
		wg2.Add(n)
		exited := make(chan bool, n)
		for i := 0; i != n; i++ {
			go func() {
				wg1.Done()
				wg2.Wait()
				exited <- true
			}()
		}
		wg1.Wait()
		for i := 0; i != n; i++ {
			select {
			case <-exited:
				t.Fatal("WaitGroup released group too soon")
			default:
			}
			wg2.Done()
		}
		for i := 0; i != n; i++ {
			<-exited // Will block if barrier fails to unlock someone.
		}
	}
}
