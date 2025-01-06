package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Kirill")

	got := buffer.String()
	want := "Hello, Kirill"

	if got != want {
		t.Errorf("Got %q Want %q", got, want)
	}
}
