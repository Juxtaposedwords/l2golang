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
var (
	visited = make(map[int]bool)
)

func longestPalindrome(s string) string {
	n := int(len(s) / 2)
	var longestPal string
	for i := 0; i < n; i++ {
		//	fmt.Printf("s: %s i:%d n: %d", s, i, n)
		var x, y string
		switch {
		//ensure we don't hit ourselves
		case n-i < 0:
			x = getPalindrome(s, n+i)
		case n+i > len(s)-1:
			x = getPalindrome(s, n-i)
		default:
			x, y = getPalindrome(s, n-i), getPalindrome(s, n+i)
		}
		//	fmt.Printf(" x:%s y:%s longestPal: %s\n", x, y, longestPal)
		switch {
		case len(x) >= len(y) && len(x) > len(longestPal):
			longestPal = x
		case len(x) < len(y) && len(y) > len(longestPal):
			longestPal = y
		}
	}

	return longestPal

}

func getPalindrome(s string, index int) string {
	c := strings.Split(s, "")
	l, r := index, index

	//is this to the left or in the middle?

	for l >= 0 && r < len(c) {
		if c[l] != c[r] {
			break
		}
		l--
		r++
	}
	l++
	r--
	fl, fr := l, r

	// let's search on even!
	l, r = index, index+1
	for l >= 0 && r < len(c) {
		if c[l] != c[r] {
			break
		}
		l--
		r++
	}
	l++
	r--
	if r-l > fr-fl {
		return strings.Join(c[l:r+1], "")

	} else {
		return strings.Join(c[fl:fr+1], "")

	}

}

func absVal(input int) int {
	if input >= 0 {
		return input
	} else {
		return input * -1
	}
}

func maxIter(length, index int) int {
	var o int
	if 0-(index) <= length-(index) {
		o = absVal(0 - index)
	} else {
		o = absVal(length - index)
	}
	return o
}
