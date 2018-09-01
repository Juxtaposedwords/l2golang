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
		{[]string{}, ""},
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
		have [2]string
		want string
	}{
		{[2]string{"flower", "flight"}, "fl"},
		{[2]string{"dog", "racecar"}, ""},
		{[2]string{"flipped", "flipped"}, "flipped"},
	}
	for _, v := range tt {
		got := commonPrefix(v.have[0], v.have[1])
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("commonPrefix: %#v got:%s", v, got)
		}
	}
}
