package palindrome

import (
	//	"fmt"
	"strings"
)

/*
"Given a string s, find the longest palindromic substring in s. You may assume
that the maximum length of s is 1000."

The first approach to this was iterative, but ran out of time. So we have a few
options:
    1. Find a new way to solve the problem
    2. Find a way to trade memory for time (typically divie & conquer)


First approach:
    For each character in the word look over every previous letter to see if that
      makes a palindrome.
        * Stop at the first one you find (the others won't be longer)
        *

*/
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	l := strings.Split(s, "")

	//get the longest palidnrome for the next one.
	left := longestPalindrome(strings.Join(l[1:], ""))
	right := firstPalindrome(s)
	if len(left) > len(right) {
		return left
	} else {
		return right
	}
}

func isPalindromic(s string) bool {
	c := strings.Split(s, "")
	n := int(len(c) / 2)
	for i := 0; i < n; i++ {
		if c[i] != c[len(c)-1-i] {
			return false
		}
	}
	return true
}

func firstPalindrome(s string) string {
	c := strings.Split(s, "")
	for i := len(c); i >= 0; i-- {
		e := strings.Join(c[0:i], "")
		if isPalindromic(e) {
			return e
		}
	}

	return ""
}

// What would divide and conquer look like?
//     Recursive, by defintion
//          Early exit?
//  If there is only one
