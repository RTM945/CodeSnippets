package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	//var wg sync.WaitGroup
	count := 1 * 10000
	start := time.Now()
	ch := make(chan int)
	for i := 0; i < count; i++ {
		//wg.Add(1)
		go request(func() {
			//wg.Done()
			ch <- 1
		})
	}
	//wg.Wait()
	for _ = range ch {
		<-ch
	}
	fmt.Printf("cost %v", time.Since(start))
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
		fmt.Print(err)
		return
	}

	expected := "1"
	if string(body) != expected {
		fmt.Printf("resp: %s", string(body))
	}
	cb()
}
