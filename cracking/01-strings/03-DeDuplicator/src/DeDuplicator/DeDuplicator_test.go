package DeDuplicatorString

import (
	"testing"
)

func TestDeDuplicator(t *testing.T) {
	tt := []struct {
		have string
		want string
	}{
		{"Forsooth", "Forsth"},
		{"Clothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothoothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothoothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothoothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothoothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothoothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothothotho", "Cloth"},
		{"Forsake", "Forsake"},
	}
	for _, e := range tt {
		if e.want != DeDuplicator(e.have) {
			t.Errorf("There was a problem with: %+V", e)
		}
	}
}
