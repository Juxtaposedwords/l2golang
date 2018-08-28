package longest

import (
	//	"fmt"
	"strings"
)

// Problem: Given a string, find the length of the longest substring
//  without repeating characters.

//  Solution Time: O(n) *  1,114,112
//         the runtime is deceptive as it can get quite large, as any unicode
//         rune can be used as a "character."
//
//  Approach:
//  Create the list you want to compare against as you go, and store the
//       highest list
// 1. Step over rune in the slice/string
//     a. If we we've seen it before truncate the previous list to include
//          everything until right after it.
//     b. Append the new item to the list
//     c. check if this new list is the largest one yet

func lengthOfLongestSubstring(s string) int {
	z := []string{}
	o := 0
	//  we use a 1 based loop to help keep better track of the distance
	//  step over every item in the list
	for _, e := range strings.Split(s, "") {
		for j, f := range z {
			if f == e {
				z = z[j+1:]
			}
		}
		z = append(z, e)
		if len(z) > o {
			o = len(z)
		}
	}

	return o
}
