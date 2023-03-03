# iodebug

The iodebug package is a tool for debugging input and output (I/O) operations in Go. It implements interfaces like io.Reader and io.WriteCloser, which can be used as wrappers to monitor the data flow at runtime.

## Installation

To install the package, use the following command:

```
go get github.com/gustavosbarreto/iodebug
```

## How to Use

The iodebug package offers following types of wrappers:

* `Reader` for `io.Reader`
* `Writer` for `io.Writer`
* `WriteCloser` for `io.WriteCloser`

### Using Reader

The Reader wrapper is used to monitor read operations. To use it, you need to create an object of type Reader, passing the object you want to monitor as a parameter:

```go
package main

import (
    "fmt"
    "io"
    "os"

    "github.com/gustavosbarreto/iodebug"
)

func main() {
    f, err := os.Open("file.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    r := iodebug.Reader{Reader: f}

    buf := make([]byte, 1024)
    for {
        n, err := r.Read(buf)
        if err != nil && err != io.EOF {
            fmt.Println(err)
            return
        }
        if n == 0 {
            break
        }
        fmt.Print(string(buf[:n]))
    }
}
```

In the example above, the object f is opened using the os.Open function and passed as a parameter to the Reader wrapper. The wrapper is then used in the read operation with the Read function, which returns the bytes read and a possible error.

### Using Writer

The Writer wrapper is used to monitor write operations. To use it, you need to create an object of type Writer, passing the object you want to monitor as a parameter:

```go
package main

import (
	"fmt"
	"os"

	"github.com/gustavosbarreto/iodebug"
)

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	w := iodebug.Writer{Writer: f}

	msg := "Hello, world!"
	_, err = fmt.Fprint(w, msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = w.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
```

In the example above, the object f is created using the os.Create function and passed as a parameter to the Writer wrapper. The wrapper is then used in the write operation with the Fprint function from the fmt package, which receives the bytes to be written.

### Using WriteCloser

The WriteCloser wrapper is used to monitor write and close operations. To use it, you need to create an object of type WriteCloser, passing the object you want to monitor as a parameter:

```go
package main

import (
    "fmt"
    "io"
    "os"

    "github.com/gustavosbarreto/iodebug"
)

func main() {
    f, err := os.Create("file.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    wc := iodebug.WriteCloser{WriteCloser: f}

    msg := "Hello, world!"
    _, err = wc.Write([]byte(msg))
    if err != nil {
        fmt.Println(err)
        return
    }

    err = wc.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}
```

In the example above, the object f is created using the os.Create function and passed as a parameter to the WriteCloser wrapper. The wrapper is then used in the write operation with the Write function, which receives the bytes to be written, and in the close operation with the Close function.

## Contributing

Contributions are always welcome! If you have found a bug or have any suggestions for improvement, please open an issue or pull request on the GitHub repository.
