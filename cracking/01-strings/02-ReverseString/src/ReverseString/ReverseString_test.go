package ReverseString

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		have string
		want string
	}{
		{"Tacocat", "tacocaT"},
		{"Amber", "rebmA"},
		{"otto", "otto"},
		{"yes", "sey"},
	}
	for _, e := range tt {
		if e.want != Reverse(e.have) {
			t.Errorf("There was an issue with: %+V", e)
		}
	}
}
