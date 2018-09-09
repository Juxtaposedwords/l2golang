package palindrome

import (
	"fmt"
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
	var longestPal string
	fmt.Println(s)
	n := int(len(s) / 2)
	l, r := n, n
	if len(s)%2 == 0 {
		r++
	}
	for l >= 0 && r < len(s) {
		lv, rv := getPalindrome(s, l), getPalindrome(s, r)
		fmt.Printf("s: %s l:%d r: %d lv: %s rv: %s\n", s, l, r, lv, rv)
		switch {
		case len(lv) < len(longestPal) || len(rv) < len(longestPal):
			break
		case len(lv) >= len(rv):
			longestPal = lv
		case len(lv) < len(rv):
			longestPal = rv
		}
		l--
		r++
	}

	return longestPal

}

func getPalindrome(s string, index int) string {
	c := strings.Split(s, "")
	l, r := index, index
	switch {

	case (len(c)-1)-index > 0 && c[index] == c[index+1]:
		r++
	}
	//is this to the left or in the middle?
	var bit bool
	// let's search on even!
	for l >= 0 && r < len(c) {
		if c[l] != c[r] {
			break
		}
		bit = true
		l--
		r++
	}
	if bit {
		l++
		r--
	}
	return strings.Join(c[l:r+1], "")

}
