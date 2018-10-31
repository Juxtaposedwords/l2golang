package binary

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	short, long := strings.Split(a, ""), strings.Split(b, "")
	if len(short) > len(long) {
		short, long = long, short
	}
	carry := false
	output := make([]string, len(long))

	for i := 0; i < len(long); i++ {
		s, l := len(short)-i-1, len(long)-i-1
		se, le := "0", long[l]
		if s >= 0 {
			se = short[s]
		}

		sum := getBinaryInt(se) + getBinaryInt(le)
		if carry {
			sum += 1
			carry = false
		}

		if sum > 1 {
			carry = true
		}
		output[l] = fmt.Sprintf("%d", sum%2)

	}
	if carry {
		output = append([]string{"1"}, output...)
	}
	return strings.Join(output, "")
}

func getBinaryInt(input string) int {
	switch input {
	case "1":
		return 1
	default:
		return 0
	}
}
