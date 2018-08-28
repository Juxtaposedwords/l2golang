package PalindromeNumbers

import (
	"testing"
)

type private struct {
	privateField string
}

// Our function is local, but we still want to test it so we declare a version
//    that is exported and so the test will run
var IsPalindrom = isPalindrome

func TestIsPalindrome(t *testing.T) {

	tt := []struct {
		have int
		want bool
	}{
		{121, true},
		{-121, false},
	}
	for _, v := range tt {
		got := isPalindrome(v.have)
		if got != v.want {
			t.Errorf("%v got: %b want: %b", v, got, v.want)
		}
	}
}
