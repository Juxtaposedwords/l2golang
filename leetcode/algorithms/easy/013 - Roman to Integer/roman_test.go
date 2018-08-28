package roman

import (
	"testing"
)

func TestRoman(t *testing.T) {

	tt := []struct {
		have string
		want int
	}{
		{"III", 3},
		{"IV", 4},
		{"DCCXCVIII", 798},
		{"LXXXIII", 83},
	}
	for _, v := range tt {
		got, err := Roman(v.have)
		if err != nil {
			t.Errorf("Error for %v: %s", v, err)
		}
		if got != v.want {
			t.Errorf("Test: %v Got: %d want: %d ", v, got, v.want)
		}
	}
}
