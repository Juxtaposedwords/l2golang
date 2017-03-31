package palinperm

import (
	"fmt"
	"strings"
)

func PalinPerm(s string) bool {
	charCount := make(map[string]int)
	for _, e := range []rune(s) {
		if e == ' ' {
			continue
		}
		c := strings.ToLower(string(e))
		if _, ok := charCount[c]; ok {
			charCount[c]++
		} else {
			charCount[c] = 1
		}
	}
	oddCount := 0
	for _, v := range charCount {
		if v%2 != 0 {
			oddCount++
		}
		if oddCount > 1 {
			return false
		}
	}
	return true
}
