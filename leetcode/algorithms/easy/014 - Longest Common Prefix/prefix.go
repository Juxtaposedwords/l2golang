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
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}

	// We begin at first by assumming there is complete overlap
	o := strs[0]

	for i := 1; i < len(strs); i++ {
		// Now we got through all the other word and check if there is a match
		o = commonPrefix(o, strs[i])
		if len(o) == 0 {
			return ""

		}
	}
	return o
}

// We break common prefix into a separate function for testing and
//    and for ease of reading.
func commonPrefix(str1 string, str2 string) string {
	s1, s2 := strings.Split(str1, ""), strings.Split(str2, "")
	//Ensure str1 is always the shortest string, for ease of iterating
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	// Create the output array
	o := []string{}

	// Now loop over the array until we find a non-match
	for i := range s1 {
		if s1[i] != s2[i] {
			return strings.Join(o, "")
		}
		o = append(o, s1[i])
	}
	return strings.Join(o, "")
}
