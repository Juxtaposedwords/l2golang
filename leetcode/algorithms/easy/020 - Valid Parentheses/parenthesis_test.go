package parenthesis

import (
	"testing"
)

var IsValid = isValid

func TestIsValid(t *testing.T) {

	tt := []struct {
		have string
		want bool
	}{
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"]", false},
	}
	for _, v := range tt {
		got := isValid(v.have)
		if got != v.want {
			t.Errorf("isValid %#v got:%b", v, got)
		}
	}
}
