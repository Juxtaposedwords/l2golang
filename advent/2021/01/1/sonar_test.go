package sonar

import (
	"testing"
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
