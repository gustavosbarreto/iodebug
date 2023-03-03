package iodebug

import (
	"fmt"
	"io"
)

type Reader struct {
	io.Reader
}

func (r *Reader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	if err != nil {
		fmt.Printf("io.Reader.Read error: %s\n", err)
	}
	var b = make([]byte, len(p))
	copy(b, p)
	fmt.Printf("io.Reader.Read: %s\n", string(b))
	return n, err
}
