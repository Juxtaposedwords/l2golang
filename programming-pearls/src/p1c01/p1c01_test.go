package p1c01

import (
	"testing"
)

func TestBitmap(t *testing.T) {


	good := []int{45, 15395, 7997, 7190, 11796}
	bad := []int{30000, 30519, 31510}
	b := NewBitmap()
	_ = b.LoadFromFile("input.txt")

	for _, v := range good {
		if got, want := b.GetBit(v), true; got != want {
			t.Errorf("b.GetBit(%d): got %t, want %t", v, got, want)
		}
	}
	for _, v := range bad {
		if got, want := b.GetBit(v), false; got != want {
			t.Errorf("b.GetBit(%d): got %t, want %t", v, got, want)
		}
	}

}
