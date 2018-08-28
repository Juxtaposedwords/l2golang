package zigzag

import (
	"strings"
)

/*for each rune in string:

put in next space down, if available
if not move over one and up
once you get to the top, go down

*/
func convert(s string, numRows int) string {
	goingDown := false
	data := make([][]string, numRows)
	// Set the index operator, as computer science
	n := numRows - 1
	// If for some reason they only gave us 1 row, return the answer
	if n == 0 {
		return s
	}
	for i, e := range strings.Split(s, "") {
		// if we're at the top or bottom invert the direction we're going
		if i%n == 0 {
			goingDown = !goingDown
		}
		// set the variable we'll use for selecting the slice
		var x int

		// if we're going down, we'll use the moduolous of the index
		if goingDown {
			x = i % n
			// otherwise we're going up, then we'll use the inverse
		} else {
			x = n - (i % n)
		}
		// now we append the string to the slice
		data[x] = append(data[x], e)
	}
	var o []string
	// build the slice of string characters we'll join fro output
	for _, e := range data {
		o = append(o, e...)
	}
	// now join the glorious output
	return strings.Join(o, "")
}
