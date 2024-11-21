package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Calculator interface {
	Calculate(a, b int) int
}

type Handler struct {
	stdout    io.Writer
	calulator Calculator
}

func NewHandler(stdout io.Writer, calculator Calculator) *Handler {
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

	calculator := this.calulator
	result := calculator.Calculate(a, b)

	_, err = fmt.Fprint(this.stdout, result)
	if err != nil {
		return fmt.Errorf("%w: %w", errWriterFailure, err)
	}

	return nil
}

var errWrongNumberOfArgs = errors.New("usage: calc [a] [b]")
var errInvalidArgument = errors.New("invalid argument")
var errWriterFailure = errors.New("writer failure")
