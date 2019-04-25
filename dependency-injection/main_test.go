package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	name := "Chris"
	buffer := bytes.Buffer{}
	Greet(&buffer, name)

	got := buffer.String()
	want := "Hello, " + name

	if got != want {
		t.Errorf("got '%s' want %s'", got, want)
	}
}
