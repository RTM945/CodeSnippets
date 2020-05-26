package redis

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBulk(t *testing.T) {
	str := "$13\r\nHello, World!\r\n"
	b := []byte(str)
	fmt.Printf("%v\n", b)
	if b[0] == '$' {
		i := 1
		for ; ; i++ {
			if b[i] == '\n' && b[i-1] == '\r' {
				i++
				break
			}
		}
		count, _ := strconv.Atoi(string(b[1 : i-2]))
		fmt.Println(count)
		bulk := string(b[i : i+count])
		fmt.Println(bulk)
	}
}
