package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13r *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13r.r.Read(b)
	for i := range b {
		if (b[i] >= 'A' && b[i] <= 'Z'-13) || (b[i] >= 'a' && b[i] <= 'z'-13) {
			b[i] += 13
		} else if (b[i] >= 'A'+13 && b[i] <= 'Z') || (b[i] >= 'A'+13 && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
