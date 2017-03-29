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
		{"Mr John Smith      ", 13, "Mr%20John%20Smith"},
		{"Muad'Dib   The  Desert Mouse       ", 35, "Muad'Dib%20%20%20The%20%20Desert%20Mouse"},
		{"  way To The Golden Path            ", 24, "%20%20way%20To%20The%20Golden%20Path"},
	}

	for _, e := range tt {
		if e.want != URLify(e.haveString, e.haveLength) {
			t.Errorf("There was a problem with %+v\n Got: %s, want: %s ", e, URLify(e.haveString, e.haveLength), e.want)
		}
	}
}
