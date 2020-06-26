package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChanWithoutBuffer1(t *testing.T) {
	// 不带缓冲的chan无论发送与接收都会阻塞当前线程
	c := make(chan struct{})
	c <- struct{}{}
	fmt.Println(<-c)
}

func TestChanWithoutBuffer2(t *testing.T) {
	// 不带缓冲的chan无论发送与接收都会阻塞当前线程
	c := make(chan struct{})
	// 使用goroutine避免阻塞
	go func(c chan struct{}) {
		time.Sleep(time.Second)
		c <- struct{}{}
	}(c)
	fmt.Println(<-c)
}

func TestChanWithBuffer(t *testing.T) {
	// 使用带缓冲的chan
	c := make(chan struct{}, 1)
	c <- struct{}{}
	fmt.Println(<-c)
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
