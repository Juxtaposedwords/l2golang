package img

import (
	"bufio"
	"github.com/google/go-cmp/cmp"
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
func TestLayerMatrix(t *testing.T) {
	tests := []struct {
		desc          string
		haveInput     io.Reader
		haveColumns   int
		haveRows      int
		wantLayers    [][][]int
		wantErrorCode codes.Code
	}{
		{
			desc:        "Day 8: layer composition",
			haveInput:   strings.NewReader("123456789012"),
			haveColumns: 3,
			haveRows:    2,
			wantLayers: [][][]int{
				[][]int{
					[]int{1, 2, 3},
					[]int{4, 5, 6},
				},
				[][]int{
					[]int{7, 8, 9},
					[]int{0, 1, 2},
				},
			},
		},
		{
			desc:          "invalid digit length",
			haveInput:     strings.NewReader("12345678901"),
			haveColumns:   3,
			haveRows:      2,
			wantErrorCode: codes.FailedPrecondition,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			//t.Parallel()
			got, err := layerMatrix(tc.haveInput, tc.haveColumns, tc.haveRows)
			if got, want := status.Code(err), tc.wantErrorCode; got != want {
				t.Errorf("layerMatrix() unexpected error. want: %s got: %s - %s", want, got, err.Error())
				return
			}
			if tc.wantErrorCode != codes.OK {
				return
			}
			if diff := cmp.Diff(tc.wantLayers, got); diff != "" {
				t.Errorf("layerMatrix() mismatch (-want +got):\n%s\n got: %#v want: %#v", diff, got, tc.wantLayers)
			}
		})
	}
}
func TestLayerDigitSums(t *testing.T) {
	tests := []struct {
		desc          string
		haveInput     io.Reader
		haveColumns   int
		haveRows      int
		wantValue     int
		wantErrorCode codes.Code
	}{
		{
			desc:        "Day 8: layer composition",
			haveInput:   strings.NewReader("00225781012"),
			haveColumns: 3,
			haveRows:    2,
			wantValue:   2,
		},
		{
			desc:        "Day 8:actual problem",
			haveInput:   loadFile(t, "/workspace/l2golang/advent/2019/08/src/input.txt"),
			haveColumns: 25,
			haveRows:    6,
			wantValue:   1848,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
		//	t.Parallel()
			got, err := layerDigitSums(tc.haveInput, tc.haveColumns, tc.haveRows)
			if got, want := status.Code(err), tc.wantErrorCode; got != want {
				t.Errorf("layerDigitSums() unexpected error. want: %s got: %s - %s", want, got, err)
				return
			}
			if tc.wantErrorCode != codes.OK {
				return
			}
			if got != tc.wantValue {
				t.Errorf("layerDigitSums() mismatch got: %#v want: %#v", got, tc.wantValue)
			}
		})
	}
}
