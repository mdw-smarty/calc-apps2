package handlers

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mdw-smarty/calc-lib2"
)

var rawInput = strings.Join([]string{
	`1,+,2`,
	`2,-,1`,
	`NaN,+,1`,
	`1,+,NaN`,
	`2,+,3`,
}, "\n")

var csvCalculators = map[string]Calculator{"+": &calc.Addition{}}

func TestCSVHandler(t *testing.T) {
	var output bytes.Buffer
	var logs bytes.Buffer
	handler := NewCSVHandler(strings.NewReader(rawInput), &output, &logs, csvCalculators)

	err := handler.Handle()

	assertError(t, err, nil)
	assertEqual(t, output.String(), "1,+,2,3\n2,+,3,5\n")
	if logs.Len() > 0 {
		t.Logf("\n%s", logs.String())
	}
}
func TestCSVHandler_ReaderError(t *testing.T) {
	reader := &ErringReader{err: boink}
	var output bytes.Buffer
	var logs bytes.Buffer
	handler := NewCSVHandler(reader, &output, &logs, csvCalculators)

	err := handler.Handle()

	assertError(t, err, boink)
	if logs.Len() > 0 {
		t.Logf("\n%s", logs.String())
	}
}
func TestCSVHandler_WriterError(t *testing.T) {
	output := &ErringWriter{err: boink}
	var logs bytes.Buffer
	handler := NewCSVHandler(strings.NewReader(rawInput), output, &logs, csvCalculators)

	err := handler.Handle()

	assertError(t, err, boink)
	if logs.Len() > 0 {
		t.Logf("\n%s", logs.String())
	}
}
