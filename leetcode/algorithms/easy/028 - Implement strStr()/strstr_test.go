package strstr

import (
	"testing"
)

var StrStr = strStr

func TestStrStr(t *testing.T) {
	type h struct {
		haystack string
		needle   string
	}
	tt := []struct {
		have *h
		want int
	}{
		{&h{"hello", "ll"}, 2},
		{&h{"aaaaa", "bba"}, -1},
		{&h{"aa", "a"}, 0},
		{&h{"mississippi", "pi"}, 9},
	}
	for _, v := range tt {
		got := strStr(v.have.haystack, v.have.needle)
		if got != v.want {
			t.Errorf("strStr have: %#v want: %d got: %d", v.have, v.want, got)
		}
	}
}
