package palinperm

import (
	"testing"
)

func TestPalinPerm(t *testing.T) {
	tt := []struct {
		have string
		want bool
	}{
		{"Tact Coa", true},
		{"ttoo", true},
		{"easea", true},
		{"ea e", true},
	}
	for _, e := range tt {
		if e.want != PalinPerm(e.have) {
			t.Errorf("There was an error with: %+v\n", e)
		}
	}
}
