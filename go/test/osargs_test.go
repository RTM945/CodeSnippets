package main

import (
	"fmt"
	"os"
	"testing"
)

func TestOsArgs(t *testing.T) {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}
