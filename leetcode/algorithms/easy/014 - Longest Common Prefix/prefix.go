package prefix

import (
	"strings"
)

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
Note:

All given inputs are in lowercase letters a-z.
*/
func longestCommonPrefix(strs []string) string {

	// Our logic will always be comparing two strings. We'll start off by
	//   returning anything that is too short
	if len(strs) < 2 {
		return strs[0]
	}

	// We begin at first by assumming there is complete overlap
	o := strings.Split(strs[0], "")

	for i := 1; i < len(strs); i++ {
		// Now we got through all the other word and check if there is a match
		x := commonPrefix(o, strings.Split(strs[i], ""))

		// If the prefix is smaller then we'll set that as the new limit
		if len(o) > len(x) {
			o = x
		}
	}
	return strings.Join(o, "")
}

// We break common prefix into a separate function for testing and
//    and for ease of reading.
func commonPrefix(str1 []string, str2 []string) []string {
	//Ensure str1 is always the shortest string, for ease of iterating
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}
	// Create the output array
	o := []string{}

	// Now loop over the array until we find a non-match
	for i := range str1 {
		if str1[i] != str2[i] {
			return o
		}
		o = append(o, str1[i])
	}
	return o
}
