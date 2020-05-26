package redis

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRESPWriter(t *testing.T) {
	var buf bytes.Buffer
	writer := NewRESPWriter(&buf)
	writer.WriteCommand("GET", "foo")
	fmt.Println(string(buf.Bytes()))
}
