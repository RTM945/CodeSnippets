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

// 假设有一个超长的切片，切片的元素类型为int，切片中的元素为乱序排列。
// 限时5秒，使用多个goroutine查找切片中是否存在给定值，
// 在找到目标值或者超时后立刻结束所有goroutine的执行。
// 比如切片为：[23, 32, 78, 43, 76, 65, 345, 762, …… 915, 86]，
// 查找的目标值为345，如果切片中存在目标值程序输出:"Found it!"
// 并且立即取消仍在执行查找任务的goroutine。
// 如果在超时时间未找到目标值程序输出:"Timeout! Not Found"，
// 同时立即取消仍在执行查找任务的goroutine。
func TestSearch(t *testing.T) {
	data := []int{1, 2, 3, 10, 999, 8, 345, 7, 98, 33, 66, 77, 88, 68, 96}
	dataLen := len(data)
	step := 3
	target := 345
	timeout := time.NewTimer(5 * time.Second)
	resultC := make(chan struct{}, 1)
	ctx, cancle := context.WithCancel(context.Background())

	for i := 0; i < dataLen; i += step {
		end := i + step
		if end > dataLen {
			end = dataLen - 1
		}
		go func(ctx context.Context, s []int, target int, c chan<- struct{}) {
			for _, v := range s {
				select {
				case <-ctx.Done():
					fmt.Println("task cancled!")
					return
				default:
				}
				fmt.Printf("v: %d\n", v)
				time.Sleep(1 * time.Second)
				if v == target {
					resultC <- struct{}{}
					return
				}
			}
		}(ctx, data[i:end], target, resultC)
	}

	select {
	case <-resultC:
		fmt.Println("Found it!")
		cancle()
	case <-timeout.C:
		fmt.Println("Timeout! Not Found.")
		cancle()
	}
	time.Sleep(time.Second)
}

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
