package pilot

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNavigate(t *testing.T) {
	tests := []struct {
		desc       string
		have       []string
		wantResult int
		wantCode   codes.Code
	}{
		{
			desc:       "Provided example",
			have:       []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			wantResult: 900,
		},

		{
			desc:     `Command is not in "{operator} {value}" format`,
			have:     []string{"forward up and onward yet"},
			wantCode: codes.FailedPrecondition,
		},
		{
			desc:     "Value is not a number",
			have:     []string{"forward up"},
			wantCode: codes.InvalidArgument,
		},
		{
			desc:     "invalid operator",
			have:     []string{"forward 2", "down 4", "updog 3"},
			wantCode: codes.InvalidArgument,
		},
	}

	for _, tc := range tests {
		tc := tc // Without this t.Parallel() will break)
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := Navigate(tc.have)
			if got, want := status.Code(err), tc.wantCode; got != want {
				t.Fatalf("Navigate() unexpected status code. want: %s got: %s", want, got)
			}
			if got != tc.wantResult {
				t.Errorf("Navigate() mismatch want: %d got: %d", tc.wantResult, got)
			}
		})
	}
}
