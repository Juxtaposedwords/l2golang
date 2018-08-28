package ReverseInteger

import (
	"fmt"
	"strconv"
	"strings"
)

func Reverse(x int) int {

	neg := false
	if x < 0 {
		neg = true
		x = Abs(x)
	}
	s := reverseArray(strings.Split(strconv.Itoa(x), ""))

	x = sliceToInt(s)
	if neg {
		x *= -1
	}
	return x

}

func sliceToInt(input []string) int {
	var x int
	fmt.Println("%v", input)
	s := strings.Join(input, "")
	_, err := strconv.ParseInt(s, 10, 32)
	fmt.Println(err)
	if err != nil {
		x = 0

	} else {
		x, _ = strconv.Atoi(s)
	}
	return x

}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func reverseArray(input []string) []string {
	for i := len(input)/2 - 1; i >= 0; i-- {
		opp := len(input) - 1 - i
		input[i], input[opp] = input[opp], input[i]
	}
	return input
}
