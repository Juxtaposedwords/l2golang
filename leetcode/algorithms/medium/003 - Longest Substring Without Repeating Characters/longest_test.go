package longest

import (
	"testing"
)

var LengthOfLongestSubstring = lengthOfLongestSubstring

func TestLengthOfLongestSubstring(t *testing.T) {
	tt := []struct {
		have string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"dvdf", 3},
	}

	for _, v := range tt {
		got := LengthOfLongestSubstring(v.have)
		if got != v.want {
			t.Errorf("LengthOfLongestSubstring failed:\n got: %d \nwant:%d\n have: %s", got, v.want, v.have)
		}
	}
}
