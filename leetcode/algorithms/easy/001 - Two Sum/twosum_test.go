package TwoSum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	have := []struct {
		given  []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 26, []int{2, 3}},
		{[]int{2, 7, 11, 15}, 0, []int{}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, v := range have {
		got := twoSum(v.given, v.target)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("Error with %#v, got %v", v, got)
		}
	}
}
