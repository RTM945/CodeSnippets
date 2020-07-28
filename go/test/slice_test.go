package main

import "testing"

func TestSlice(t *testing.T) {
	a := make([]int, 0)
	f := func(slice []int) {
		slice = append(slice, 1)
	}
	f(a)
	t.Log(len(a)) //0
	//ineffectual assignment to `slice`
}
