package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buff := bytes.Buffer{}
	Greet(&buff, "chris")

	got := buff.String()
	expected := "Hello, chris"

	if got != expected {
		t.Errorf("expected %q, but got %q", expected, got)
	}
}
