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
        * Use divide and conquer

Second approach:
    Split the work and callign longestPlaindrome recursively so that
    longestPalindrome("tacocat")
    |                           \
    firstPalindrome("tacocat")  longestPlaindrome("acocat")
                                 |                          \
                                 firstPalindrome("acocat")   longestPalindrome("cocat")
    ...
Note:
    This doesn't solve our problem and doesn't help the asymptotic growth.
        So that means we need to reconsider the data-struct involved as it will
            inform the algorithm

Third approach:

*/
func longestPalindrome(s string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return s
	case 2:
		return firstPalindrome(s)
	}

	sl := strings.Split(s, "")
	n := int(len(sl) / 2)
	var o string
	l, r := strings.Split(sl[:n], ""), strings.Split(sl[n-1:], "")

	//get the longest palidnrome for the next one.
	left := longestPalindrome(strings.Join(l[:], ""))
	right := longestPalindrome(strings.Join(l[:], ""))
	if len(left) > len(right) {
		return left
	} else {
		return right
	}
}

func isPalindromic(s string) bool {
	c := strings.Split(s, "")
	n := int(len(c) / 2)
	if len(c)%2 != 0 {
		n -= 1
	}
	for i := n; i > 0; i-- {
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
