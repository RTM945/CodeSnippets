package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range words {
		m[v]++
		//弟弟操作
		//count, ok := m[v]
		//if ok {
		//	count = count + 1
		//} else {
		//	count = 1
		//}
		//m[v] = count

	}
	return m
}

func main() {
	wc.Test(WordCount)
}
