package zeroOneGame

import (
	"reflect"
	"testing"
)

func TestZeroOneGame(t *testing.T) {
	tt := []struct {
		have []int
		want string
	}{
		{[]int{1, 0, 0, 1}, "Bob"},
		{[]int{1, 0, 1, 0, 1}, "Alice"},
		{[]int{0, 0, 0, 0, 0, 0}, "Bob"},
	}
	for _, e := range tt {
		got := zeroOneGame(e.have)
		if e.want != got {
			t.Errorf("There was an issue with: %+v, got: %s", e, got)
		}
	}
}
func TestValidPlays(t *testing.T) {
	tt := []struct {
		have []int
		want []int
	}{
		{[]int{0, 0, 0, 0, 1, 0}, []int{1, 2, 4}},
		{[]int{1, 0, 0, 1}, []int{}},
	}

	for _, e := range tt {
		got := validPlays(e.have)
		if !reflect.DeepEqual(e.want, got) {
			t.Errorf("There was an error with %+v, got: %+v", e, got)
		}
	}
}
