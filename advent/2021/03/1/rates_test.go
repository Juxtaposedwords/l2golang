package rates

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCompute(t *testing.T) {
	tests := []struct {
		desc       string
		have       []string
		wantResult int
		wantCode   codes.Code
	}{
		{
			desc:       "Provided example",
			have:       []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			wantResult: 198,
		},
	}

	for _, tc := range tests {
		tc := tc // Without this t.Parallel() will break)
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := Compute(tc.have)
			if got, want := status.Code(err), tc.wantCode; got != want {
				t.Fatalf("Compute() unexpected status code. want: %s got: %s", want, got)
			}
			if got != tc.wantResult {
				t.Errorf("Compute() mismatch want: %d got: %d", tc.wantResult, got)
			}
		})
	}
}
