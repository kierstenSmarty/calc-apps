package handlers

import (
	"bytes"
	"errors"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/kierstenSmarty/calc-lib"
)

var csvInput = strings.Join([]string{
	"1,+,2",
	//"2,-,1",
	//"NaN,+,2",
	//"1,+,NaN",
	//"1,nop,2",
	"3,+,4",
	//"20,/,10",
}, "\n")

var csvOutput = strings.Join([]string{
	"1,+,2,3",
	"3,+,4,7",
	"",
}, "\n")

func TestCSVHandler(t *testing.T) {
	var logBuffer bytes.Buffer
	logger := log.New(&logBuffer, "", log.LstdFlags)
	reader := strings.NewReader(csvInput)
	var outputBuffer bytes.Buffer
	calculators := map[string]Calculator{"+": &calc.Addition{}}
	handler := NewCSVHandler(logger, reader, &outputBuffer, calculators)

	err := handler.Handle()

	assertError(t, err, nil)
	if outputBuffer.String() != csvOutput {
		t.Errorf("got %q, want %q", outputBuffer.String(), csvOutput)
	}
	t.Log(logBuffer.String())
}

func TestCSVHandlerWriteError(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)
	reader := strings.NewReader(csvInput)
	output := ErringWriter{err: boink}
	calculators := map[string]Calculator{"+": &calc.Addition{}}
	handler := NewCSVHandler(logger, reader, &output, calculators)

	err := handler.Handle()

	assertError(t, err, boink)
}

func TestCSVHandlerReadError(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)
	reader := ErringReader{err: boink}
	var output bytes.Buffer
	calculators := map[string]Calculator{"+": &calc.Addition{}}
	handler := NewCSVHandler(logger, &reader, &output, calculators)

	err := handler.Handle()

	assertError(t, err, boink)
}

var boink = errors.New("boink")

type ErringReader struct {
	err error
}

func (this *ErringReader) Read(_ []byte) (n int, err error) {
	return 0, this.err
}
