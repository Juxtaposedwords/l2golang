package uniqueString

import (
	"testing"
)

func TestUnique(t *testing.T) {
	tt := []struct {
		have string
		want bool
	}{
		{"Clotho", false},
		{"Lachesi", true},
		{"Ll", true},
		{"Clothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothoothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothoothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothoothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothoothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothoothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothotho", false},
	}
	for _, e := range tt {
		if e.want != Unique(e.have) {
			t.Errorf("There was an error with: %+V", e)
		}
	}
}
