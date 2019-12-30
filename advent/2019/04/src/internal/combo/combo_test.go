package combo

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestBrute(t *testing.T) {
	tests := []struct {
		desc      string
		haveLower int
		haveUppper  int
		want      int
		wantError error
	}{
				{
		desc: "happyPath",
		haveLower: 11,
		haveUppper: 12,
		want: 1,
		},
						{
		desc: "lower and upper equal",
		haveLower: 22,
		haveUppper: 22,
		want: 1,
		},
		{
		desc: "lower > upper",
		haveLower: 12,
		haveUppper: 11,
		wantError: status.Error(codes.InvalidArgument, "lower > upper"),
		},
		{
		desc: "upper and lower length mismatch",
		haveLower: 120,
		haveUppper: 11,
		wantError: status.Error(codes.InvalidArgument, "length mismatch"),
		},

	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := Brute(tc.haveLower, tc.haveUppper)
			if got, want := status.Code(err), status.Code(tc.wantError); got != want {
				t.Errorf("Brute() unexpected error. want: %s got: %s", want, got)
				return
			}
			if got != tc.want {
				t.Errorf("Brute() mismatch want: %d got: %d", tc.want, got)
			}
		})
	}
}