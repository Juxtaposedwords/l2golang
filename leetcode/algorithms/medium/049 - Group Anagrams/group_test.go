package group

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
			have: []string{"a"},
			want: [][]string{
				[]string{"a"},
			},
		},
	}
	for _, tc := range tests {
		tc := tc // for concurrency
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := GroupAnagrams(tc.have)
			lessSlice := func(l, r []string) bool { return len(l) < len(r) }
			lessAlpha := func(l, r string) bool { return l < r }

			if diff := cmp.Diff(tc.want, got, cmpopts.SortSlices(lessAlpha), cmpopts.SortSlices(lessSlice)); diff != "" {
				t.Errorf("IsAnagram(...) output different from expected(-want +got):\n%s", diff)
			}

		})
	}

}
