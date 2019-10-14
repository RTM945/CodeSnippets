package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map
	go func() {
		for i := 0; i < 100; i++ {
			go func(i int) {
				m.Store(i, i)
				//fmt.Printf("add %d\n", i)
			}(i)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println("匹配开始")
			var s []int
			m.Range(func(key, value interface{}) bool {
				fmt.Printf("%d = %d\n", key.(int), value.(int))
				i := value.(int)
				s = append(s, i)
				length := len(s)
				fmt.Printf("s length = %d\n", length)
				if length == 5 {
					fmt.Printf("%v匹配完成", s)
					return false
				}
				return true
			})
			if len(s) == 5 {
				for _, v := range s {
					m.Delete(v)
				}
			}
			fmt.Println("匹配结束")
			time.Sleep(time.Second)
		}
	}()

	<-make(chan int)
}
