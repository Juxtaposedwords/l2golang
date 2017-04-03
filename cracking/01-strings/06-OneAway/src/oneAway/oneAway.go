package oneAway

import (
	"fmt"
)

func OneAway(a, b string) bool {
	c, d := []rune(a), []rune(b)
	var maxLen int
	if len(c) < len(d) {
		maxLen = len(d)
	} else {
		maxLen = len(c)
	}
	j, k, diff := 0, 0, 0
	fmt.Printf("\n")
	for j < maxLen-1 && k < maxLen-1 {
		switch {
		// check if the lengths are different
		case len(c) == j+1 && len(d)-1 > k:
			k++
		case len(d) == k+1 && len(c)-1 > j:
			j++
		// see if they are the same
		case c[j] == d[k]:
			j++
			k++
			continue
		// check to see if one was inserted or deleted in the second
		case len(d) >= k+1 && c[j] == d[k+1]:
			k++
		// check to see if one was inserted or deleted in the first
		case len(c) >= j+1 && c[j+1] == d[k]:
			j++
		default:
			j++
			k++
		}
		diff++
		if diff >= 2 {
			return false
		}
	}
	return true
}
