package hydro

import (
	"bufio"
	"bytes"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestOverlap(t *testing.T) {
	tests := []struct {
		desc       string
		haveReader *bufio.Reader
		haveLevel  int
		wantResult int
		wantCode   codes.Code
	}{
		{
			desc: "Provided example",
			haveReader: bufio.NewReader(bytes.NewBufferString(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`)),
			haveLevel:  2,
			wantResult: 12,
		},
	}

	for _, tc := range tests {
		tc := tc // Without this t.Parallel() will break)
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := Overlap(tc.haveLevel, tc.haveReader)
			if got, want := status.Code(err), tc.wantCode; got != want {
				t.Fatalf("Overlap() unexpected status code. want: %s got: %s err: %s", want, got, err)
			}
			if got != tc.wantResult {
				t.Errorf("Overlap() mismatch want: %d got: %d", tc.wantResult, got)
			}
		})
	}
}
