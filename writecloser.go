package iodebug

import (
	"fmt"
	"io"
)

type WriteCloser struct {
	io.WriteCloser
}

func (wc *WriteCloser) Write(p []byte) (int, error) {
	n, err := wc.WriteCloser.Write(p)
	if err != nil {
		fmt.Printf("io.WriteCloser.Write error: %s\n", err)
	}
	fmt.Printf("io.WriteCloser.Write: %s\n", p)
	return n, err
}

func (wc *WriteCloser) Close() error {
	if err := wc.WriteCloser.Close(); err != nil {
		fmt.Printf("io.WriteCloser.Close error: %s\n", err)
		return err
	}
	return nil
}
