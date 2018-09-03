package strstr

import (
	"strings"
)

/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

Approach 1:
 Dumb simple, sliding window to look ahead and see if we match.

 Runtime:
    c * n
        where c = number of characters in haystack, and n = number of
                            charaters in needle
Considered approach/struct:
    * Trie
*/

func strStr(haystack string, needle string) int {
	switch {
	// Handles the edge of empty strings
	case needle == "":
		return 0
	//This ensure no out of bound issues
	case len(haystack) < len(needle):
		return -1
	}
	h := strings.Split(haystack, "")
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if sliceEquals(strings.Join(h[i:len(needle)+i], ""), needle) {
			return i
		}
	}
	return -1
}

func sliceEquals(a, b string) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
