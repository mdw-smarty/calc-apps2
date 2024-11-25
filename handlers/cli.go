package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/mdw-smarty/calc-lib2"
)

type CLIHandler struct {
	calculator *calc.Addition
	stdout     io.Writer
}

func NewCLIHandler(calculator *calc.Addition, stdout io.Writer) *CLIHandler {
	return &CLIHandler{
		calculator: calculator,
		stdout:     stdout,
	}
}

func (this *CLIHandler) Handle(args []string) error {
	if len(args) != 2 {
		return wrongArgCount
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: %w", invalidArg, err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: %w", invalidArg, err)
	}
	c := this.calculator.Calculate(a, b)
	_, err = fmt.Fprintf(this.stdout, "%d", c)
	if err != nil {
		return fmt.Errorf("%w: %w", errWriter, err)
	}
	return nil
}

var (
	wrongArgCount = errors.New("two operands required")
	invalidArg    = errors.New("invalid argument")
	errWriter     = errors.New("writer error")
)
