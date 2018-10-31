package binary

import (
	"testing"
)

var AddBinary = addBinary

func TestAddBinary(t *testing.T) {

	tt := []struct {
		have [2]string
		want string
	}{
		{[2]string{"11", "1"}, "100"},
		{[2]string{"1010", "1011"}, "10101"},
		{[2]string{"1111", "1111"}, "11110"},
	}

	for _, v := range tt {
		got := addBinary(v.have[0], v.have[1])
		if got != v.want {
			t.Errorf("\nlongestPalindrome failed with\n  test: %#v\n  got: %#v\n", v, got)
		}
	}
}
