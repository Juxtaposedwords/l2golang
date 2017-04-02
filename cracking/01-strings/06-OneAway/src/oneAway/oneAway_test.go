package oneAway

import (
	"testing"
)

func TestOneAway(t *testing.T) {
	tt := []struct {
		have [2]string
		want bool
	}{
		{[2]string{"pale", "ple"}, true},
		{[2]string{"pales", "pale"}, true},
		{[2]string{"pale", "bale"}, true},
		{[2]string{"pale", "bake"}, false},
	}
	for _, e := range tt {
		got := OneAway(e.have[0], e.have[1])
		if e.want != got {
			t.Errorf("There was an error with: %+v got: %v\n", e, got)
		}
	}
}
