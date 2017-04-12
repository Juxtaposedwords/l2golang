package beautifulWord

import (
	"testing"
)

func TestIsBeautiful(t *testing.T) {
	tt := []struct {
		have string
		want string
	}{
		{"batman", "yes"},
		{"apple", "no"},
		{"beauty", "no"},
	}

	for _, e := range tt {
		if IsBeautiful(e.have) != e.want {
			t.Errorf("There was a problem with: %+v\n", e)
		}
	}
}
