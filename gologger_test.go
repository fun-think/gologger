package gologger_test

import (
	"io"
	"os"
	"testing"

	"github.com/fun-think/gologger"
	"github.com/fun-think/gologger/format"
	"github.com/fun-think/gologger/writer"
)

func TestGologger(t *testing.T) {
	Logger := &gologger.Logger{
		Level:  gologger.DEBUG,
		Format: new(format.TextFormat),
		Output: io.MultiWriter(
			&writer.DailyFileWriter{
				Name:     "cache/logs",
				MaxCount: 7,
			},
			os.Stdout,
		),
	}
	Logger.Error("err")
	t.Fatal("err")
}
