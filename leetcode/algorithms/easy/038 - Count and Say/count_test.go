package count

import (
	"testing"
)

var (
	CountAndSay     = countAndSay
	NextSaySequence = nextSaySequence
)

func TestCountAndSay(t *testing.T) {
	tt := []struct {
		have int
		want string
	}{
		{1, "1"},
		{4, "1211"},
	}

	for _, v := range tt {
		got := countAndSay(v.have)
		if got != v.want {
			t.Errorf("countAndSay failed with\n  have: %#v\n  want: %d\n  got:  %#v\n", v.have, v.want, got)
		}
	}

}

func TestNextSaySequence(t *testing.T) {
	tt := []struct {
		have string
		want string
	}{
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
	}
	for _, v := range tt {
		got := nextSaySequence(v.have)
		if got != v.want {
			t.Errorf("nextSaySequence failed with\n  have: %#v\n  want: %d\n  got:  %#v\n", v.have, v.want, got)
		}
	}
}
