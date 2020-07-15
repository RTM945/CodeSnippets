package algo

import (
	"fmt"
	"testing"
)

func BenchmarkSearch(B *testing.B) {
	for i := 0; i < B.N; i++ {
		s := []int{5, 4, 3, 2, 1}
		target := 3
		fmt.Println(Search(s, target))
	}
}

func TestSearch(t *testing.T) {
	s := []int{5, 4, 3, 2, 1}
	target := 3
	fmt.Println(Search(s, target))
}

func TestBinarySearch(t *testing.T) {
	s := []int{5, 4, 3, 2, 1}
	target := 3
	fmt.Println(BinarySearch(s, target))
	fmt.Println(BinarySearch1(s, target))
}
