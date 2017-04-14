package bsearch

import (
	"fmt"
)

func BSearch(a []int, b int) (int, error) {
	pos := (len(a) - 1) / 2
	d := (len(a) - 1)
	for {
		d = halfer(d)
		switch {
		case a[pos] == b:
			return pos, nil
		case d < 1:
			return 0, fmt.Errorf("Entry was not found")
		case a[pos] > b:
			pos = pos - d
		case a[pos] < b:
			pos = pos + d
		}
	}
}

func halfer(d int) int {
	h := int(d / 2)
	if d%2 != 0 {
		h = h + 1
	}
	return h
}
