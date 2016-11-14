package p1c01

import (
	"testing"
)

func TestBitmap(t *testing.T) {

	b := NewBitmap()

	good := []int{0, 1, 1026, 1802, 29108}
	bad := []int{2, 3, 16, 1024, 1025, 1027, 29107, 29109}
	for _, v := range good {
		b.SetBit(v)
	}

	for _, v := range good {
		if got, want := b.CheckBit(v), true; got != want {
			t.Errorf("b.CheckBit(%d): got %t, want %t", v, got, want)
		}
	}
	for _, v := range bad {
		if got, want := b.CheckBit(v), false; got != want {
			t.Errorf("b.CheckBit(%d): got %t, want %t", v, got, want)
		}
	}

}