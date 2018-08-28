package zigzag

import (
	"testing"
)

var Convert = convert

func TestConvert(t *testing.T) {
	tt := []struct {
		haveString string
		haveRow    int
		want       string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"a", 1, "a"},
	}
	for _, v := range tt {
		got := convert(v.haveString, v.haveRow)
		if got != v.want {
			t.Errorf("Issue with %v, got: %s", v, got)
		}
	}
}
