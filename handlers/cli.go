package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/kierstenSmarty/calc-lib"
)

type Handler struct {
	stdout    io.Writer
	calulator *calc.Addition
}

func NewHandler(stdout io.Writer, calculator *calc.Addition) *Handler {
	return &Handler{
		stdout:    stdout,
		calulator: calculator,
	}
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errWrongNumberOfArgs
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: %w", errInvalidArgument, err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: %w", errInvalidArgument, err)
	}

	calculator := &calc.Addition{}
	result := calculator.Calculate(a, b)

	_, err = fmt.Println(result)
	if err != nil {
		return err
	}

	return nil
}

var errWrongNumberOfArgs = errors.New("usage: calc [a] [b]")
var errInvalidArgument = errors.New("invalid argument")
var errWriterFailure = errors.New("writer failure")
