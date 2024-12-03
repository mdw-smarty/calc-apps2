package handlers

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mdw-smarty/calc-lib2"
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

var rawInput = strings.Join([]string{
	`1,+,2`,
	`2,-,1`,
	`NaN,+,1`,
	`1,+,NaN`,
	`2,+,3`,
}, "\n")

func TestCSVHandler(t *testing.T) {
	gunit.Run(new(CSVHandlerFixture), t)
}

type CSVHandlerFixture struct {
	*gunit.Fixture
}

var csvCalculators = map[string]Calculator{"+": &calc.Addition{}}

func (this *CSVHandlerFixture) TestCSVHandler() {
	var output bytes.Buffer
	handler := NewCSVHandler(strings.NewReader(rawInput), &output, this, csvCalculators)

	err := handler.Handle()

	this.So(err, should.BeNil)
	this.So(output.String(), should.Equal, "1,+,2,3\n2,+,3,5\n")
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
