package intcode

import (
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestList(t *testing.T) {
	tests := []struct {
		desc          string
		have          []int
		want          []int
		wantErrorCode codes.Code
	}{
		{
			desc: "opcode 1 happy path",
			have: []int{1, 0, 1, 2, 99},
			want: []int{1, 0, 1, 2, 99},
		},
		{
			desc: "opcode 2 happy path",
			have: []int{2, 1, 2, 3, 99},
			want: []int{2, 1, 2, 2, 99},
		},
		{
			desc: "Changes to same value",
			have: []int{1, 1, 2, 3, 99},
			want: []int{1, 1, 2, 3, 99},
		},
		{
			desc: "Provided example: 1",
			have: []int{1, 0, 0, 0, 99},
			want: []int{2, 0, 0, 0, 99},
		},
		{
			desc: "Provided example: 2",
			have: []int{2, 3, 0, 3, 99},
			want: []int{2, 3, 0, 6, 99},
		},
		{
			desc: "Provided example: 3",
			have: []int{2, 4, 4, 5, 99, 0},
			want: []int{2, 4, 4, 5, 99, 9801},
		},
		{
			desc: "Provided example: 4",
			have: []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			want: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
		{
			desc:          "invalid opcode",
			have:          []int{4, 1, 0, 2, 3},
			wantErrorCode: codes.InvalidArgument,
		},
		{
			desc:          "not terminated by 99",
			have:          []int{2, 0, 1, 2},
			wantErrorCode: codes.FailedPrecondition,
		},
		{
			desc:          "invalid opcode created",
			have:          []int{1, 1, 2, 4, 99},
			wantErrorCode: codes.InvalidArgument,
		},
		{
			desc:          "out of bounds",
			have:          []int{1, 16, 2, 4, 99},
			wantErrorCode: codes.InvalidArgument,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			//t.Parallel()
			gotResp, err := list(tc.have)
			if got, want := status.Code(err), tc.wantErrorCode; got != want {
				t.Errorf("List() unexpected error. want: %s got: %s", want, got)
				return
			}
			if tc.wantErrorCode != codes.OK {
				return
			}
			if diff := cmp.Diff(gotResp, tc.want); diff != "" {
				t.Errorf("List() mismatch (-want +got):\n%s\n got: %#v want: %#v", diff, gotResp, tc.want)
			}
		})
	}
}
