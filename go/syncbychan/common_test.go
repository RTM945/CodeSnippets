package syncbychan

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestChanAsNonBlockingQueue(t *testing.T) {
	q := make(chan int, 10)
	for i := 0; i < 10; i++ {
		q <- i
		// non-blocking write
		// var ok bool
		// select {
		// case q <- i:
		// 	ok = true
		// default:
		// 	ok = false
		// }
	}
	for {
		var ok bool
		var i int
		select {
		case i, ok = <-q:
		default:
		}
		if ok {
			fmt.Println(i)
		} else {
			break
		}
	}
	fmt.Println("end")
}

func TestTimeout(t *testing.T) {
	start := time.Now()
	done := make(chan struct{})
	timeout := time.Duration(3 * time.Second)
	go func() {
		//do some time-consuming work
		// time.Sleep(time.Duration(3 * time.Second))
		time.Sleep(time.Duration(1 * time.Second))
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(timeout):
	}
	fmt.Println(time.Since(start))
}

func TestTimeoutContext(t *testing.T) {
	start := time.Now()
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	done := make(chan struct{})
	go func() {
		//do some time-consuming work
		// time.Sleep(time.Duration(3 * time.Second))
		time.Sleep(time.Duration(1 * time.Second))
		close(done)
	}()
	select {
	case <-done:
	case <-ctx.Done():
	}
	fmt.Println(time.Since(start))
}

func TestPingPong(t *testing.T) {
	type Ball struct {
		hits int
	}
	table := make(chan *Ball)
	player := func(name string, table chan *Ball) {
		for {
			ball := <-table
			ball.hits++
			fmt.Println(name, ball.hits)
			time.Sleep(100 * time.Millisecond)
			table <- ball
		}
	}
	go player("ping", table)
	go player("pong", table)
	table <- new(Ball)
	time.Sleep(time.Second)
	<-table
}
