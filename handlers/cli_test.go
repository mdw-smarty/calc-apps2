package handlers

import (
	"bytes"
	"errors"
	"strconv"
	"testing"

	"github.com/mdw-smarty/calc-lib2"
)

func assertEqual(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("got: [%s], want: [%s]", actual, expected)
	}
}
func assertError(t *testing.T, err, expected error) {
	if !errors.Is(err, expected) {
		t.Errorf("expected err to wrap %v, but it didn't", expected)
	}
}

func TestHandler_NotEnoughArgs(t *testing.T) {
	handler := NewCLIHandler(nil, nil)
	err := handler.Handle(nil)
	assertError(t, err, wrongArgCount)
}
func TestHandler_InvalidFirstArgument(t *testing.T) {
	handler := NewCLIHandler(nil, nil)
	err := handler.Handle([]string{"NaN", "1"})
	assertError(t, err, invalidArg)
	assertError(t, err, strconv.ErrSyntax)
}
func TestHandler_InvalidSecondArgument(t *testing.T) {
	handler := NewCLIHandler(nil, nil)
	err := handler.Handle([]string{"1", "NaN"})
	assertError(t, err, invalidArg)
	assertError(t, err, strconv.ErrSyntax)
}
func TestHandler_CalculationResultWritten(t *testing.T) {
	var output bytes.Buffer
	handler := NewCLIHandler(&calc.Addition{}, &output)
	err := handler.Handle([]string{"1", "2"})
	assertError(t, err, nil)
	assertEqual(t, output.String(), "3")
}
func TestHandler_WriterError(t *testing.T) {
	boink := errors.New("boink")
	output := &ErringWriter{err: boink}
	handler := NewCLIHandler(&calc.Addition{}, output)
	err := handler.Handle([]string{"1", "2"})
	assertError(t, err, boink)
	assertError(t, err, errWriter)
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}
