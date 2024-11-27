package handlers

import (
	"errors"
	"reflect"
	"testing"
)

var boink = errors.New("boink")

/////////////////////////////////////////////////////

func assertEqual(t *testing.T, actual, expected any) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got:\n%v\n\nwant:\n%v", actual, expected)
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
