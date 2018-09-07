package search

import (
//	"fmt"
)

/*
In order to get the most performant time we can use the binary search to ensure
    nlogn time.
*/
func searchInsert(nums []int, target int) int {

	// Now we also store that as the location, as we'll be halving that while
	//    we search
	start, end := 0, len(nums)-1
	for start <= end {
		m := (start + end) / 2
		if nums[m] < target {
			start = m + 1
		} else {
			end = m - 1
		}
	}
	return start
}
