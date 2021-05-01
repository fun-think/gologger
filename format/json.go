package format

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
	"github.com/fun-think/gologger"
)

// JSONFormat is a json formatter
type JSONFormat struct {
	AppName    string
	TimeFormat string

	init sync.Once
	host string
	pid  int
}

// Format implements log.Formatter
func (f *JSONFormat) Format(level gologger.Level, msg string, logger *gologger.Logger) []byte {
	// output fields: time level host app pid file line msg

	f.init.Do(func() {
		if f.AppName == "" {
			f.AppName = filepath.Base(os.Args[0])
		}
		if f.TimeFormat == "" {
			f.TimeFormat = "2006-01-02 15:04:05.000"
		}

		f.host, _ = os.Hostname()
		f.pid = os.Getpid()
	})

	data := make(map[string]interface{}, 8)

	// file, line
	file, line := FilelineCaller(4)

	data["time"] = time.Now().Format(f.TimeFormat)
	data["level"] = level.String()
	data["host"] = f.host
	data["app"] = f.AppName
	data["pid"] = f.pid
	data["file"] = file
	data["line"] = line
	data["msg"] = msg

	serialized, err := marshal(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal json, %v\n", err)
	}
	return serialized
}

func marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}