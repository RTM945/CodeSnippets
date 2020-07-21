package solution095

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	trees := generateTrees(3)
	for _, t := range trees {
		fmt.Printf("%v\n", *t)
	}

}
