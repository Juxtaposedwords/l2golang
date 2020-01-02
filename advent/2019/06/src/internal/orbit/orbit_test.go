package orbit

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"strings"
	"bufio"
	"os"
	"testing"
)
func loadFile(t *testing.T, filepath string) io.Reader{
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
			have: loadFile(t, "/workspace/l2golang/advent/2019/06/src/internal/orbit/input.txt"),
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
