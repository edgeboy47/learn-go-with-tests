package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Name")

	got := buffer.String()
	want := "Hello Name"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
