package accurateSorting

import (
	"testing"
)

func TestIsAbsSortable(t *testing.T) {
	tt := []struct {
		haveCount int
		haveSlice []int
		want      string
	}{
		{4, []int{1, 0, 3, 2}, "Yes"},
		{3, []int{2, 1, 0}, "No"},
		{5, []int{12, 1, 0, 2, 4}, "No"},
		{7, []int{12, 1, 4, 3, 0, 2, 4}, "No"},
	}
	for _, test := range tt {
		got := IsAbsSortable(test.haveCount, test.haveSlice)
		if got != test.want {
			t.Errorf("Error with: %+v, got: %s\n", test, got)
		}
	}
}
