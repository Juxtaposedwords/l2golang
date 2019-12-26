package fuel

import (
	"testing"
)

func TestRequired(t *testing.T) {
	tests := []struct {
		desc string
		have uint64
		want uint64
	}{
		{
			desc: "happy path",
			have: 14,
			want: 2,
		},
		{
			desc: "big Endian",
			have: 1,
			want: 0,
		},
		{desc: "provide example #1",
			have: 12,
			want: 2,
		},
		{
			desc: "free fuel upper limit",
			have: 6,
			want: 0,
		},
		{
			desc: "free fuel lower limit",
			have: 0,
			want: 0,
		},
		{
			desc: "upper limit of input",
			have: 18446744073709551615,
			want: 0,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := required(tc.have)
			if got != tc.want {
				t.Errorf("Required(%d) returned %d. Expected: %d", tc.have, got, tc.want)
			}
		})
	}
}
