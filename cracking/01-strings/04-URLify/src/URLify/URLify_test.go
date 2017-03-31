package URLify

import (
	"testing"
)

func TestURLify(t *testing.T) {
	tt := []struct {
		haveString string
		haveLength int
		want       string
	}{
		{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
		{"Muad'Dib   The  Desert Mouse            ", 28, "Muad'Dib%20%20%20The%20%20Desert%20Mouse"},
		{"  way To The Golden Path            ", 24, "%20%20way%20To%20The%20Golden%20Path"},
		{"PaulAtreides", 12, "PaulAtreides"},
	}

	for _, e := range tt {
		got, err := URLify([]rune(e.haveString), e.haveLength)
		want := e.want
		if got != want {
			t.Errorf("URLify(%q, %d): got %q, want %q error: %s", e.haveString, e.haveLength, got, want, err)
		}
	}
}
