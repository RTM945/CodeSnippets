package main

import "fmt"

type Team []int

func main() {
	var t Team
	t = append(t, 1)
	fmt.Println(t)
}
