package main

import "fmt"

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		tmp := first
		first, second = second, first+second
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
