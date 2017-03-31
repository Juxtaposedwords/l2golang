package URLify

import (
	"fmt"
	"strings"
)

func URLify(r []rune, l int) (string, error) {
	if len(r) < l {
		return "", fmt.Errorf("Length must be longer than the length of the string")
	}
	last := len(r) - 1
	beforeString, afterString := strings.Count(string(r[0:l]), " "), strings.Count(string(r[l:]), " ")
	if afterString != beforeString*2 {
		return "", fmt.Errorf("Impossible to pad, not enough trailing whitepsace")
	}

	for i := l - 1; i >= 0; i-- {
		switch {
		case string(r[i]) == " " && last != l-1:
			r[last-2], r[last-1], r[last] = rune('%'), rune('2'), rune('0')
			last -= 3
		case string(r[i]) != " ":
			r[last] = r[i]
			last--
		case string(r[i]) == " " && last == l-1:
			continue
		}
		fmt.Printf("i: %d l: %s, String is now: %s\n  last: %d\n", i, string(r[i]), string(r), last)
	}
	return string(r), nil
}
