package syncbychan

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	var once sync.Once
	for i := 0; i < 5; i++ {
		once.Do(func() {
			fmt.Println("once")
		})
	}
}

func TestChanOnce(t *testing.T) {
	once := NewOnce()
	for i := 0; i < 5; i++ {
		once.Do(func() {
			fmt.Println("once")
		})
	}
}
