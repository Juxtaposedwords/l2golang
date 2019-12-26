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
			want: 6148914691236517203,
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
func TestTotalRequired(t *testing.T) {
	tests := []struct {
		desc string
		have []uint64
		want uint64
	}{
		{
			desc: "happy path",
			have: []uint64{7},
			want: 1,
		},
		{
			desc: "0 valued",
			have: []uint64{0,0,0,1,2,3,4,5,6},
			want: 0,
		},
		{
			desc: "Largest possible value",
			have: []uint64{18446744073709551615,18446744073709551615,18446744073709551615},
			want: 18446744073709551609,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := TotalRequired(tc.have)
			if got != tc.want {
				t.Errorf("TotalRequired(%d) returned %d . Expected: %d", tc.have, got, tc.want)
			}
		})
	}
}
