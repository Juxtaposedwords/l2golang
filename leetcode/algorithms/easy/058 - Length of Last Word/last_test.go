package last

import (
	"testing"
)

var LengthOfLastWord = lengthOfLastWord

func TestLengthOfLastWord(t *testing.T) {

	tt := []struct {
		have string
		want int
	}{
		{"Hello World", 5},
		{"HelloWorld", 10},
		{"helloWorld ?", 1},
		{"a", 1},
		{"a ", 1},
	}

	for _, v := range tt {
		got := lengthOfLastWord(v.have)
		if got != v.want {
			t.Errorf("\nlongestPalindrome failed with\n  test: %#v\n  got: %#v\n", v, got)
		}
	}
}
