package orbit

import (
	"bufio"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"os"
	"strings"
	"testing"
)

func loadFile(t *testing.T, filepath string) io.Reader {
	f, err := os.Open(filepath)
	if err != nil {
		t.Fatal(err.Error())
	}
	return bufio.NewReader(f)
}

 func TestTotal(t *testing.T) {
	tests := []struct {
		desc          string
		have          io.Reader
		wantResp      int
		wantErrorCode codes.Code
	}{
		{
			desc: "happy path",
			have: strings.NewReader(`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`),
			wantResp: 42,
		},
		{
			desc: "provided example",
			have: loadFile(t, "/workspace/l2golang/advent/2019/06/input.txt"),
			wantResp: 227612,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := total(tc.have)
			if got, want := status.Code(err), tc.wantErrorCode; got != want {
				t.Errorf("total() unexpected error. want: %s got: %s", got, want)
				return
			}
			if got != tc.wantResp {
				t.Errorf("total() returned %d . Expected: %d", got, tc.wantResp)
			}
		})
	}
} 
func TestSanDistance(t *testing.T) {
	tests := []struct {
		desc          string
		have          io.Reader
		wantResp      int
		wantErrorCode codes.Code
	}{
		{
			desc: "happy path",
			have: strings.NewReader(`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`),
			wantResp: 4,
		},
		{
			desc:     "provided example",
			have:     loadFile(t, "/workspace/l2golang/advent/2019/06/input2.txt"),
			wantResp: 454,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := sanDistance(tc.have)
			if got, want := status.Code(err), tc.wantErrorCode; got != want {
				t.Errorf("sanDistance() unexpected error. want: %s got: %s", got, want)
				return
			}
			if got != tc.wantResp {
				t.Errorf("sanDistance() returned %d . Expected: %d", got, tc.wantResp)
			}
		})
	}
}
