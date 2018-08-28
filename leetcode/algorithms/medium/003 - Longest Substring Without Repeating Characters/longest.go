package longest

import (
	"fmt"
	"strings"
)

// Problem: Given a string, find the length of the longest substring
//  without repeating characters.

//  Solution Time: O(n)

//  Approach:
// Create empty maps[string]int of lastSeen  and distances
// 1. Step over rune in the slice/string
//     a. If we haven't seen it before(it's not the lastSeen dict):
//          i. add the rune to Last seen with the index
//     b. If we have seen it before
//          i. update the index
//          ii. check if the new distance is greater
// 2. Step throug the dictionary and find the highest value

func lengthOfLongestSubstring(s string) int {
	l := []string{}
	distance := 0
	lastSeen := make(map[string]int)
	l := strings.Split(s, "")
	o := 0
	//  we use a 1 based loop to help keep better track of the distance
	//  step over every item in the list
	for i := 1; i <= len(l); i++ {

		// Get the character in question
		e := l[i-1]

		// Find out if we've seen it before
		val, seen := lastSeen[e]
		if seen {
			distance = i - val
			if distance > o {
				o = distance
			}
			distance = 1
		} else {
			distance++
			lastSeen[e] = i
		}
		fmt.Printf("Distance: %d slice: %v\n", distance, l[i-distance:i])

	}
	if distance > o {
		o = distance
	}

	return o
}
