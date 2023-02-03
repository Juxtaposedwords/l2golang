package anagram

import (
	"testing"
)

func TestRemover(t *testing.T) {
	tt := []struct {
		desc      string
		haveLeft  string
		haveRight string
		want      bool
	}{
		{
			desc:      "case 1: basic example of match",
			haveLeft:  "anagram",
			haveRight: "nagaram",
			want:      true,
		},
		{
			desc:      "case 2: basic example of mismatch",
			haveLeft:  "rat",
			haveRight: "car",
			want:      false,
		},
		{
			desc:      "case 2: rude length of left",
			haveLeft:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			haveRight: "car",
			want:      false,
		},
	}
	for _, tc := range tt {
		tc := tc // Required for parallel tests: https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			if got := IsAnagram(tc.haveLeft, tc.haveRight); got != tc.want {
				t.Errorf("IsAnagaram(%s,%s):Failed  to get expected output. Got: %t Want: %t", tc.haveLeft, tc.haveRight, got, tc.want)

			}
		})
	}
}
