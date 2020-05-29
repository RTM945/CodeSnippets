package syncbychan

import (
	"fmt"
	"testing"
)

func TestOnce(t *testing.T) {
	once := NewOnce()
	for i := 0; i < 5; i++ {
		once.Do(func() {
			fmt.Println("once")
		})
	}
}
