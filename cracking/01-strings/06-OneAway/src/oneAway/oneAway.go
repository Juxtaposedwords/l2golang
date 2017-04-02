package oneAway

// One Away: There are three types of edits that can be performed on strings: insert a character,
// remove a character, or replace a character. Given two strings, write a function to check if they are
// one edit (or zero edits) away.
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
	for j < maxLen && k < maxLen {
		fmt.Printf("c: %s d: %s  j: %d k: %d  ", string(c[j]), string(d[k]), j, k)
		switch {
		case len(c) <= j+1:
			k++
			fmt.Printf("first\n")
		case len(d) <= k+1:
			j++
			fmt.Printf("second\n")
		// see if they are the same
		case c[j] == d[k]:
			j++
			k++
			fmt.Printf("third\n")
			continue
		// check to see if one was inserted or deleted in the second
		case len(d) >= k+1 && c[j] == d[k+1]:
			fmt.Printf("fourth\n")
			k++
		// check to see if one was inserted or deleted in the first
		case len(c) >= j+1 && c[j+1] == d[k]:
			fmt.Printf("fifth\n")
			j++
		default:
			fmt.Printf("default\n")
			j++
			k++
		}
		diff++
		fmt.Printf(" diff: %d\n", diff)
		if diff >= 2 {
			return false
		}
	}
	return true
}
