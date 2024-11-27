package handlers

import (
	"errors"
	"testing"
)

var boink = errors.New("boink")

/////////////////////////////////////////////////////

func assertEqual(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("got:\n%s\n\nwant:\n%s", actual, expected)
	}
}
func assertError(t *testing.T, err, expected error) {
	if !errors.Is(err, expected) {
		t.Errorf("expected err to wrap %v, but it didn't", expected)
	}
}

/////////////////////////////////////////////////////

type ErringReader struct {
	err error
}

func (this *ErringReader) Read(p []byte) (n int, err error) {
	return 0, this.err
}

/////////////////////////////////////////////////////

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}

/////////////////////////////////////////////////////
