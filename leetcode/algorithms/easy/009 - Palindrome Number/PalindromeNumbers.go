package PalindromeNumbers

import (
	"fmt"
	"strings"
)

func isPalindrome(x int) bool {
	s := strings.Split(fmt.Sprintf("%d", x), "")
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		if s[i] != s[opp] {
			return false
		}
	}
	return true
}
