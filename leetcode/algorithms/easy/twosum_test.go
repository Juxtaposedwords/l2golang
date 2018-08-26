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
		{[]int{2, 7, 11, 15}, 17, []int{0, 3}},
		{[]int{2, 7, 11, 15}, 0, []int{}},
	}

	for _, v := range have {
		got := twoSum(v.given, v.target)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("Error with %#v, got %v", v, got)
		}
	}
}
