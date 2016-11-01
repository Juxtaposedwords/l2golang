package main

import (
	"testing"
)

func bitMapCreate(max int) {
}
func bitMapSet(number int) {
}
func bitMapGet(number int) {
}
func TestHello(t *testing.T) {
	expectedStr := "Hello, Testing!"
	result := hello()
	if result != expectedStr {
		t.Fatalf("Expected %s, got %s", expectedStr, result)
	}
}
