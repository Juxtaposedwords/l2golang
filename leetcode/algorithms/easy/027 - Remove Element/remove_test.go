package remove

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemover(t *testing.T) {
	tt := []struct {
		desc     string
		haveNums []int
		haveVal  int
		wantResp int
		wantNums []int
	}{
		{
			desc:     "Provided example",
			haveVal:  3,
			wantResp: 2,
			haveNums: []int{3, 2, 2, 3},
			wantNums: []int{2, 2},
		},
		{
			desc:     "Provided example",
			haveVal:  2,
			wantResp: 5,
			haveNums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			wantNums: []int{0, 1, 4, 0, 3},
		},
	}
	for _, tc := range tt {
		tc := tc // Required for parallel tests: https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			if got := Remover(tc.haveNums, tc.haveVal); got != tc.wantResp {
				t.Fatalf("Remover(%#v,%d): Failed to get expected output. Got: %d Want: %d", tc.haveNums, tc.haveVal, got, tc.wantResp)
			}
			trans := cmp.Transformer("Sort", func(in []int) []int {
				out := append([]int(nil), in...) // Copy input to avoid mutating it
				sort.Ints(out)
				return out
			})
			if diff := cmp.Diff(tc.haveNums[:tc.wantResp], tc.wantNums, trans); diff != "" {
				t.Errorf("Remover(%#v,%d) input array mutation mismatch (-want +got):\n%s", tc.haveNums, tc.haveVal, diff)
			}
		})
	}
}
