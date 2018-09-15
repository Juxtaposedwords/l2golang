package last

import (
	"strings"
)

/*
Notably this only works if the string can be put into memory. to get around that
I would need to use a read buffer
*/

func lengthOfLastWord(s string) int {
	c := strings.Split(s, " ")
	if len(c) == 1 {
		return len(c[0])
	}
	for i := len(c) - 1; i >= 0; i-- {
		if len(strings.Replace(c[i], " ", "", -1)) != 0 {
			return len(c[i])
		}
	}
	return 0
}
