package operations

import (
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		desc          string
		have          int
		wantResp      *InstructionSet
		wantErrorCode codes.Code
	}{
		{
			desc: "create add insturctions with padded 0's",
			have: 1,
			wantResp: &InstructionSet{
				Operation: Add,
				First:     Position,
				Second:    Position,
				Third:     Position,
			},
		},
		{
			desc: "create multiple instructions all with immediate",
			have: 11101,
			wantResp: &InstructionSet{
				Operation: Add,
				First:     Immediate,
				Second:    Immediate,
				Third:     Immediate,
			},
		},

		{
			desc: "int too long",
			have: 8675309,
			wantErrorCode: codes.FailedPrecondition,
		},
		{
			desc: "invalid operation code",
			have: 87,
			wantErrorCode: codes.InvalidArgument,
		},
		{
			desc: "invalid first mode",
			have: 201,
			wantErrorCode: codes.InvalidArgument,
		},
		{
			desc: "invalid second mode",
			have: 2001,
			wantErrorCode: codes.InvalidArgument,
		},
		{
			desc: "invalid third mode",
			have: 20001,
			wantErrorCode: codes.InvalidArgument,
		},

	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			gotResp, err := Parse(tc.have)
			if got, want := status.Code(err), tc.wantErrorCode; got != want {
				t.Errorf("Parse() unexpected error. want: %s got: %s", want, got)
				return
			}
			if tc.wantErrorCode != codes.OK {
				return
			}
			if diff := cmp.Diff(gotResp, tc.wantResp); diff != "" {
				t.Errorf("Parse() mismatch (-want +got):\n%s\n got: %#v want: %#v", diff, gotResp, tc.wantResp)
			}
		})
	}
}
