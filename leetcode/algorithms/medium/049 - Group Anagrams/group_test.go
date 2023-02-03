package group

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGroupAnagarams(t *testing.T) {

	tests := []struct {
		desc string
		have []string
		want [][]string
	}{
		{
			desc: "Case 1: Happy Path",
			have: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				[]string{"bat"},
				[]string{"nat", "tan"},
				[]string{"ate", "eat", "tea"},
			},
		},
		{
			desc: "Case 2: Happy Path",
			have: []string{""},
			want: [][]string{
				[]string{""},
			},
		},

		{
			desc: "Case 3: Happy Path",
			have: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				[]string{},
			},
		},
	}
	for _, tc := range tests {
		tc := tc // for concurrency
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := GroupAnagrams(tc.have)
			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("IsAnagram(...) output different from expected(-want +got):\n%s", diff)
			}

		})
	}

}
