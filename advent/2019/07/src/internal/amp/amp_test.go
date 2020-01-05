package intcode

import (
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestChainedProcess(t *testing.T) {
	tests := []struct {
		desc              string
		haveSoftware      []int
		havePhaseSettings []int
		haveIntialInput   int
		want              int
		wantErrorCode     codes.Code
	}{
		{
			desc:              "Day 7: 2/2 pt 1",
			haveIntialInput:   0,
			havePhaseSettings: []int{9, 8, 7, 6, 5},
			haveSoftware:      []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
			want:              139629729,
		},
	 	{
			desc:              "Day 7: 2/2 pt 2",
			haveIntialInput:   0,
			havePhaseSettings: []int{9, 7, 8, 5, 6},
			haveSoftware:      []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
			want:              18216,
		}, 
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			//t.Parallel()
			gotResp, err := ChainedProcess(tc.haveSoftware, tc.havePhaseSettings, tc.haveIntialInput)
			if got, want := status.Code(err), tc.wantErrorCode; got != want {
				t.Errorf("ChainedProcess() unexpected error. want: %s got: %s - %s", want, got, err.Error())
				return
			}
			if tc.wantErrorCode != codes.OK {
				return
			}
			if diff := cmp.Diff(tc.want, gotResp); diff != "" {
				t.Errorf("ChainedProcess() mismatch (-want +got):\n%s\n got: %#v want: %#v", diff, gotResp, tc.want)
			}
		})
	}
}



func TestPossiblePermutations(t *testing.T) {
	tests := []struct {
		desc              string
		haveSoftware      []int
		havePhaseSettings []int
		haveIntialInput   int
		wantVal              int
		wantPhases []int
		wantErrorCode     codes.Code
	}{
		{
			desc:              "Day 7: 2/2 pt 1",
			haveIntialInput:   0,
			havePhaseSettings: []int{5,6,7,8,9},
			haveSoftware:      []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
			wantPhases: []int{9, 8, 7, 6, 5},
			wantVal:              139629729,
		},
	 	{
			desc:              "Day 7: 2/2 pt 2",
			haveIntialInput:   0,
			havePhaseSettings: []int{5,6,7,8,9},
			haveSoftware:      []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
			wantPhases: []int{9, 8, 7, 6, 5},
			wantVal:              18216,
		}, 
	}
	
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			//t.Parallel()
			gotVal, gotPhases, err := PossiblePermutations(tc.haveSoftware, tc.havePhaseSettings, tc.haveIntialInput)
			if got, want := status.Code(err), tc.wantErrorCode; got != want {
				t.Errorf("ChainedProcess() unexpected error. want: %s got: %s - %s", want, got, err.Error())
				return
			}
			if tc.wantErrorCode != codes.OK {
				return
			}
			if diff := cmp.Diff(tc.wantVal, gotVal); diff != "" {
				t.Errorf("ChainedProcess() mismatch (-want +got):\n%s\n got: %#v want: %#v", diff, gotVal, tc.wantVal)
			}
			if diff := cmp.Diff(gotPhases, tc.wantPhases); diff != "" {
				t.Errorf("ChainedProcess() combo (-want +got):\n%s\n got: %#v want: %#v", diff, gotPhases, tc.wantPhases)
			}
		})
	}
}
