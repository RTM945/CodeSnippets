package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	count := 1 * 1000
	start := time.Now()
	for i := 0; i < count; i++ {
		wg.Add(1)
		go request(func() {
			wg.Done()
		})
	}
	wg.Wait()
	fmt.Printf("spend %v\n", time.Since(start))
}

//run ../webapp/main.go first
func request(cb func()) {
	resp, err := http.Get("http://localhost:9090/test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	expected := "1"
	if string(body) != expected {
		fmt.Printf("resp: %s\n", string(body))
	}
	cb()
}
