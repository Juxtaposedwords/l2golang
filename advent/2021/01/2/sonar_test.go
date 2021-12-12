package sonar

import (
	"testing"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"
)

func TestIncreased(t *testing.T) {
	tests := []struct {
		desc string
		have []int
		want int
	}{
		{
			desc: "Happy path: proivded case",
			have: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			want: 7,
		},
		{
			desc: "one length item",
			have: []int{24601},
			want: 0,
		},
		{
			desc: "one increase and all even",
			have: []int{1, 2, 2, 2, 2, 2, 2, 2},
			want: 1,
		},
	}
	for _, tc := range tests {
		tc := tc // Without this t.Parallel() will break)
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			if got := Increased(tc.have); got != tc.want {
				t.Errorf("Increased() mismatch want: %d got: %d", tc.want, got)
			}
		})
	}
}
func TestWindowIncreased(t *testing.T) {
	tests := []struct {
		desc       string
		haveDepths []int
		haveWindow int
		want       int
		wantCode   codes.Code
	}{
		{
			desc:       "Happy path: provided case",
			haveDepths: []int{601, 618, 618, 617, 647, 716, 769, 792},
			haveWindow: 3,
			want:       5,
		},
		{
			desc:       "invalid window length",
			haveWindow: 0,
			wantCode:   codes.InvalidArgument,
			want:       0,
		},
	}
	for _, tc := range tests {
		tc := tc // Without this t.Parallel() will break)
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := WindowedIncreased(tc.haveDepths, tc.haveWindow)
			if err := status.Code(err); err != tc.wantCode {
				t.Fatalf("WindowedIncreased() unexpected error. want: %s got: %d", tc.wantCode, got)

			}
			if got != tc.want {
				t.Errorf("Increased() mismatch want: %d got: %d", tc.want, got)
			}
		})
	}
}
