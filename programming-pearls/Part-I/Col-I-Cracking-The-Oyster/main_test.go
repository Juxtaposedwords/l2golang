package main

import (
	"testing"
)

func TestSet(t *testing.T) {
	if got,want :=hello(), "Hello, Testing!"; got != want {
		  t.Errorf("hello():  got %q, want %q", got, want)
	}
}
	