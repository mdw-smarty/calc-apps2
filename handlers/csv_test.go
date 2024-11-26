package handlers

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mdw-smarty/calc-lib2"
)

func TestCSVHandler(t *testing.T) {
	rawInput := `1,+,2`
	var output bytes.Buffer
	var logs bytes.Buffer
	handler := NewCSVHandler(strings.NewReader(rawInput), &output, &logs, map[string]Calculator{"+": &calc.Addition{}})
	err := handler.Handle()
	assertError(t, err, nil)
	assertEqual(t, output.String(), "1,+,2,3\n")
}
