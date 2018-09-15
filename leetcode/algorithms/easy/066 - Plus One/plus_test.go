package plus

import (
	"reflect"
	"testing"
)

var PlusOne = plusOne

func TestPlusOne(t *testing.T) {

	tt := []struct {
		have []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{1, 9, 9}, []int{2, 0, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{4, 6, 9}, []int{4, 7, 0}},
	}

	for _, v := range tt {
		got := plusOne(v.have)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("\nlongestPalindrome failed with\n  test: %#v\n  got: %#v\n", v, got)
		}
	}
}
