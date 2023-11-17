package format

import (
	"io"
	"os"
	"runtime"
	"strings"
)

// FileLineCaller returns file and line for caller
func FileLineCaller(skip int) (file string, line int) {
	for i := 0; i < 10; i++ {
		_, file, line, ok := runtime.Caller(skip + i)
		if !ok {
			return "???", 0
		}

		// file = pkg/file.go
		n := 0
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				n++
				if n >= 2 {
					file = file[i+1:]
					break
				}
			}
		}

		if !strings.HasPrefix(file, "logs/") {
			return file, line
		}
	}

	return "???", 0
}

// IsTerminal returns whether is a valid tty for io.Writer
func IsTerminal(w io.Writer) bool {
	switch w.(type) {
	case *os.File:
		return true
	default:
		return false
	}
}
