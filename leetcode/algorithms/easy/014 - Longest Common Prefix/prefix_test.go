package prefix

import (
	"reflect"
	"strings"
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
		havea []string
		haveb []string
		want  []string
	}{
		{
			strings.Split("flower", ""),
			strings.Split("flight", ""),
			strings.Split("fl", ""),
		},
		{
			strings.Split("dog", ""),
			strings.Split("racecar", ""),
			strings.Split("", ""),
		},
		{
			strings.Split("flipped", ""),
			strings.Split("flipped", ""),
			strings.Split("flipped", ""),
		},
	}
	for _, v := range tt {
		got := commonPrefix(v.havea, v.haveb)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("commonPrefix: %#v got:%s", v, got)
		}
	}
}
