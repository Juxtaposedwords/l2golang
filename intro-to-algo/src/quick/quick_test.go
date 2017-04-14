package quick

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	tt := []struct {
		have []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 2, 3, 4, 5}, []int{1, 2, 2, 3, 3, 4, 4, 5}},
		{[]int{5, 4, 3, 3, 2, 2, 1}, []int{1, 2, 2, 3, 3, 4, 5}},
		{[]int{5, 2, 4, 6, 1, 3, 1}, []int{1, 1, 2, 3, 4, 5, 6}},
		{[]int{5, 2, 2584, 6, 1, 3}, []int{1, 2, 3, 5, 6, 2584}},
	}
	var got []int
	for _, e := range tt {
		got = quickSort(e.have)
		if !reflect.DeepEqual(e.want, got) {
			t.Errorf("insertSort():  have %v, want %v, got: %v\n", e.have, e.want, got)
		}
	}
}
