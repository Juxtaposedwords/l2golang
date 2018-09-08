package search

import (
//	"fmt"
)

/*
In order to get the most performant time we can use the binary search to ensure
    nlogn time.
*/
func searchInsert(nums []int, target int) int {

	// Store the bounds of what we're search the first and last
	start, end := 0, len(nums)-1

	// the exit case is when our window collapses on itself
	for start <= end {
		//find the middle of our search parameter
		m := (start + end) / 2
		if nums[m] < target {
			//now close the window up
			start = m + 1
		} else {
			// or close the window down
			end = m - 1
		}
	}
	return start
}
