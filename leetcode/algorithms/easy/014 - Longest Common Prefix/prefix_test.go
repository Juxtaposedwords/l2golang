package prefix

import (
	"reflect"
	"testing"
)

var (
	LongestCommonPrefix = longestCommonPrefix
	CommonPrefix        = commonPrefix
)

func TestLongestCommonPrefix(t *testing.T) {
	tt := []struct {
		have []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"", "", "   "}, ""},
	}
	for _, v := range tt {
		got := longestCommonPrefix(v.have)
		if got != v.want {
			t.Errorf("longestCommonPrefix: %#v , got %#v\n", v, got)
		}
	}
}

func TestCommmonPrefix(t *testing.T) {
	tt := []struct {
		havea string
		haveb string
		want  string
	}{
		{
			"flower",
			"flight",
			"fl",
		},
		{
			"dog",
			"racecar",
			"",
		},
		{
			"flipped",
			"flipped",
			"flipped",
		},
	}
	for _, v := range tt {
		got := commonPrefix(v.havea, v.haveb)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("commonPrefix: %#v got:%s", v, got)
		}
	}
}
