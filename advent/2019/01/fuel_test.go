package fuel

import (
	"testing"
	"fmt"
	"io"
	"strings"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
			have: 8,
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
		{
			desc: "input too large input",
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
func TestReadInts(t *testing.T) {
	tests := []struct {
		desc string
		have io.Reader
		wantResp uint64
		wantError error 
	}{
		{
			desc: "happy path",
			have: strings.NewReader("9"),
			wantResp: 1,
		},
		{
			desc: "0 valued",
			have:  strings.NewReader("0 0 0 1 2 3 4 5 6 7 8"),
						wantResp: 0,
		},
		{
			desc: "Largest possible value",
			have: strings.NewReader(fmt.Sprintf("%d",^uint64(0))),
			wantResp:  uint64(^uint64(0)/3)-2,
		},
		{
			desc: "uint overflow",
			have: strings.NewReader("18446744073709551615 18446744073709551615 18446744073709551615 18446744073709551615 18446744073709551615 18446744073709551615 18446744073709551615 18446744073709551615"),
			wantError: status.Error(codes.Internal, "uint64 overflow"),
		},
		{
			desc: "uint too big",
			have: strings.NewReader("1844674407370955161518446744073709551615184467440737095516151844674407370955161518446744073709551615184467440737095516151844674407370955161518446744073709551615"),
			wantError: status.Error(codes.InvalidArgument, "uint64 too large"),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := readInts(tc.have)
			if got, want := status.Code(err), status.Code(tc.wantError); got != want {
				t.Errorf("readints() unexpected error. want: %s got: %s",got,want )
				return
			}
			if got != tc.wantResp {
				t.Errorf("TotalRequired() returned %d . Expected: %d",  got, tc.wantResp)
			}
		})
	}
}
