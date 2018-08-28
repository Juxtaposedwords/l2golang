package ReverseInteger

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		have int
		want int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
	}

	for _, v := range tt {
		got := Reverse(v.have)
		if got != v.want {
			t.Errorf("Reverse failed for %v:  got: %d want: %d ", v, got, v.want)
		}
	}
}
