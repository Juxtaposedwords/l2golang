package beautifulWord

import (
	"testing"
)

func TestIsBeautiful(t *testing.T) {
	tt := []struct {
		have string
		want string
	}{
		{"batman", "Yes"},
		{"apple", "No"},
		{"beauty", "No"},
		{"abacaba", "Yes"},
		{"badd", "No"},
		{"yes", "No"},
	}

	for _, e := range tt {
		if IsBeautiful(e.have) != e.want {
			t.Errorf("There was a problem with: %+v got: %s\n", e, IsBeautiful(e.have))
		}
	}
}
