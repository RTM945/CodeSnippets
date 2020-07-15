package algo

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	s := []int{1, 5, 3, 2, 4}
	InsertionSort(s[:])
	fmt.Println(s)
	InsertionSort1(s[:])
	fmt.Println(s)
	InsertionSort2(s[:])
	fmt.Println(s)
	InsertionSort3(s[:])
	fmt.Println(s)
	BinaryInsertionSort(s[:])
	fmt.Println(s)
	BinaryInsertionSort1(s[:])
	fmt.Println(s)
	BinaryInsertionSort2(s[:])
	fmt.Println(s)
	BinaryInsertionSort3(s[:])
	fmt.Println(s)
}
