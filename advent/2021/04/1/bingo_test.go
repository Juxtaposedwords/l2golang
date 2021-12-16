package bingo

import (
	"bufio"
	"bytes"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestWinningScore(t *testing.T) {
	tests := []struct {
		desc       string
		have       *bufio.Reader
		wantResult int
		wantCode   codes.Code
	}{
		{
			desc: "Provided example",
			have: bufio.NewReader(bytes.NewBufferString(`7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`)),
			wantResult: 4512,
		},
	}

	for _, tc := range tests {
		tc := tc // Without this t.Parallel() will break)
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := WinningScore(tc.have)
			if got, want := status.Code(err), tc.wantCode; got != want {
				t.Fatalf("WinningScore() unexpected status code. want: %s got: %s err: %s", want, got, err)
			}
			if got != tc.wantResult {
				t.Errorf("WinningScore() mismatch want: %d got: %d", tc.wantResult, got)
			}
		})
	}
}
