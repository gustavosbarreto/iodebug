package iodebug

import (
	"fmt"
	"io"
)

type Writer struct {
	io.Writer
}

func (w *Writer) Write(p []byte) (int, error) {
	n, err := w.Writer.Write(p)
	if err != nil {
		fmt.Printf("io.Writer.Write error: %s\n", err)
	}
	fmt.Printf("io.Writer.Write: %s\n", p)
	return n, err
}
