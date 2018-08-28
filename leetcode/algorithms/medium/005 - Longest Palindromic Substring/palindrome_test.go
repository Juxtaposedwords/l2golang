package palindrome

import (
	"testing"
)

var LongestPalindrome = longestPalindrome

func TestLongestPalindrome(t *testing.T) {

	tt := []struct {
		have string
		want string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
	}

	for _, v := range tt {
		got := longestPalindrome(v.have)
		if got != v.want {
			t.Errorf("\nlongestPalindrome failed with\n  test: %#v\n  got: %#v\n", v, got)
		}
	}

}
