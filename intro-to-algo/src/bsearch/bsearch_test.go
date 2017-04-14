package bsearch

import (
	"testing"
)

func TestBSearch(t *testing.T) {
    tc := []struct {
        have_slice []int
        have_int int
        want int
    }{
        { []int{0,1,2,3,4,5,6}, 4, 4 },
        { []int{0,0,0,1,2,2,2,2,2,2,3,4,6}, 3, 10},
    }
    for _, e := range(tc) {
        got, err := BSearch(e.have_slice, e.have_int)
        if err != nil {
            t.Errorf("Error: %s\n", err)
            continue
        }
        if got != e.want {
            t.Errorf("Slice: %s Got: %s Want: %s\n", e.have_slice, got, e.want)
        }
    }
}
