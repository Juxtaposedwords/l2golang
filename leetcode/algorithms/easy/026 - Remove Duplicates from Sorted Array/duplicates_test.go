package duplicates

import (
	"reflect"
	"testing"
)

var RemoveDuplicates = removeDuplicates

func TestRemoveDuplicates(t *testing.T) {
	tt := []struct {
		have          []int
		want          int
		wantTransform []int
	}{
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{1, 1, 2}, 2, []int{1, 2}},
	}
	for _, v := range tt {
		got := removeDuplicates(v.have)
		if v.want != got ||
			!reflect.DeepEqual(
				v.have[:len(v.wantTransform)],
				v.wantTransform) {
			t.Errorf("removeDuplicates %#v got: %d\n", v, got)
		}
	}
}
