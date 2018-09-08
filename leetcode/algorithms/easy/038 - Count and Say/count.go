package count

import (
	"fmt"
	"strings"
)

/*
The count-and-say sequence is the sequence of integers with the first five terms as following:

1.            1
2.           11
3.           21
4.         1211
5.       111221
6.       312211
7.     13112221
8.   1113213211

1. make a stack of character's your tracking
2. go through each character
3. when you hit a new character write len,number, and reset stack
for each character
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.

Given an integer n, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.



Input Constraints:

1 <= n <= 30


First thoughts:
This reads similiar to the roman numeral problem. So i'm inclinde ot work towards
 genreating a dictionary as we only have 30 items.
*/
func countAndSay(n int) string {
	x := sequenceBuilder(30)
	return x[n]
}

func sequenceBuilder(n int) map[int]string {
	input := "1"
	output := map[int]string{1: "1"}
	for i := 2; i <= n; i++ {
		x := nextSaySequence(input)
		input = x
		output[i] = x
	}
	return output
}

func nextSaySequence(s string) string {
	var o string
	type counter struct {
		i   int
		val string
	}
	//	fmt.Printf("doing X: %d\n", n)
	c := &counter{i: 0, val: ""}
	for _, e := range strings.Split(s, "") {
		switch {
		case c.i == 0:
			c.val = e
			c.i++
		case c.val != e:
			o = fmt.Sprintf("%s%d%s", o, c.i, c.val)
			c.i = 1
			c.val = e
		case c.val == e:
			c.i++
		}
	}
	if c.i > 0 {
		o = fmt.Sprintf("%s%d%s", o, c.i, c.val)
	}
	return o

}
