package amplifier

import (
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestLargestPhaseCombinationResult(t *testing.T) {
	tests := []struct {
		desc              string
		haveSoftware      []int
		havePhaseSettings []int
		haveIntialInput   int
		wantValue         int
		wantCombination   []int
		wantErrorCode     codes.Code
	}{
		{
			desc:              "Day 7: 1/2 pt 1",
			haveIntialInput:   0,
			havePhaseSettings: []int{0, 1, 2, 3, 4},
			haveSoftware:      []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			wantCombination:   []int{4, 3, 2, 1, 0},
			wantValue:         43210,
		},
		{
			desc:              "Day 7: 1/2 pt 2",
			haveIntialInput:   0,
			havePhaseSettings: []int{0, 1, 2, 3, 4},
			haveSoftware:      []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			wantCombination:   []int{0, 1, 2, 3, 4},
			wantValue:         54321,
		},

		{
			desc:              "Day 7: 1/2 pt 3",
			haveIntialInput:   0,
			havePhaseSettings: []int{0, 1, 2, 3, 4},
			haveSoftware:      []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			wantCombination:   []int{1, 0, 4, 3, 2},
			wantValue:         65210,
		},
		{
			desc:              "Day 7: 1/2 pt 3",
			haveIntialInput:   0,
			havePhaseSettings: []int{0, 1, 2, 3, 4},
			haveSoftware:      []int{3,8,1001,8,10,8,105,1,0,0,21,38,55,68,93,118,199,280,361,442,99999,3,9,1002,9,2,9,101,5,9,9,102,4,9,9,4,9,99,3,9,101,3,9,9,1002,9,3,9,1001,9,4,9,4,9,99,3,9,101,4,9,9,102,3,9,9,4,9,99,3,9,102,2,9,9,101,4,9,9,102,2,9,9,1001,9,4,9,102,4,9,9,4,9,99,3,9,1002,9,2,9,1001,9,2,9,1002,9,5,9,1001,9,2,9,1002,9,4,9,4,9,99,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,99},
			wantCombination:   []int{1, 2, 3, 0, 4},
			wantValue:         277328,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			//t.Parallel()
			gotValue, gotCombination, err := LargestPhaseCombinationResult(tc.haveSoftware, tc.havePhaseSettings, tc.haveIntialInput)
			if got, want := status.Code(err), tc.wantErrorCode; got != want {
				t.Errorf("LargestPhaseCombinationResult() unexpected error. want: %s got: %s - %s", want, got, err.Error())
				return
			}
			if tc.wantErrorCode != codes.OK {
				return
			}
			if got, want := gotValue, tc.wantValue; got != want {
				t.Errorf("LargestPhaseCombinationResult()  value mismatch got: %d want: %d", got, want)
			}
			if diff := cmp.Diff(gotCombination, tc.wantCombination); diff != "" {
				t.Errorf("LargestPhaseCombinationResult() combo (-want +got):\n%s\n got: %#v want: %#v", diff, gotCombination, tc.wantCombination)
			}
		})
	}
}
func TestThrusterValue(t *testing.T) {
	tests := []struct {
		desc              string
		haveSoftware      []int
		havePhaseSettings []int
		haveIntialInput   int
		want              int
		wantErrorCode     codes.Code
	}{
		{
			desc:              "Day 7: 1/2 pt 1",
			haveIntialInput:   0,
			havePhaseSettings: []int{4, 3, 2, 1, 0},
			haveSoftware:      []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			want:              43210,
		},
		{
			desc:              "Day 7: 1/2 pt 2",
			haveIntialInput:   0,
			havePhaseSettings: []int{0, 1, 2, 3, 4},
			haveSoftware:      []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			want:              54321,
		},
		{
			desc:              "Day 7: 1/2 pt 3",
			haveIntialInput:   0,
			havePhaseSettings: []int{1, 0, 4, 3, 2},
			haveSoftware:      []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			want:              65210,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			//t.Parallel()
			gotResp, err := thrusterValue(tc.haveSoftware, tc.havePhaseSettings, tc.haveIntialInput)
			if got, want := status.Code(err), tc.wantErrorCode; got != want {
				t.Errorf("thrusterValue() unexpected error. want: %s got: %s - %s", want, got, err.Error())
				return
			}
			if tc.wantErrorCode != codes.OK {
				return
			}
			if diff := cmp.Diff(gotResp, tc.want); diff != "" {
				t.Errorf("thrusterValue() mismatch (-want +got):\n%s\n got: %#v want: %#v", diff, gotResp, tc.want)
			}
		})
	}
}
