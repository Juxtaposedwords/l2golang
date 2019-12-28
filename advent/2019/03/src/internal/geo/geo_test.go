package geo

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestClosestIntersection(t *testing.T) {
	tests := []struct {
		desc      string
		haveLeft  []string
		haveRight []string
		want      int
		wantError error
	}{
		{
			desc:      "Provided example: 1",
			haveLeft:  []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			haveRight: []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			want:      159,
		},
		{
			desc:      "Provided example: 2",
			haveLeft:  []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R5"},
			haveRight: []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			want:      135,
		},
		{
			desc:      "No left items",
			haveRight: []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			wantError: status.Error(codes.FailedPrecondition, "no left items"),
		},
		{
			desc:      "No right items",
			haveLeft:  []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R5"},
			wantError: status.Error(codes.FailedPrecondition, "no right items"),
		},
		{
			desc:      "Invalid direction: invalid cardinalty",
			haveLeft:  []string{"R98", "Z9"},
			wantError: status.Error(codes.InvalidArgument, "invalid direction"),
		},
		{
			desc:      "Invalid direction: invalid direction",
			haveLeft:  []string{"R98", "Z9"},
			wantError: status.Error(codes.InvalidArgument, "invalid direction"),
		},
		{
			desc:      "No overlap of intersections",
			haveLeft:  []string{"R98", "U9"},
			haveRight: []string{"L98", "D9"},
			wantError: status.Error(codes.NotFound, "no overlap"),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := ClosestIntersection(tc.haveLeft, tc.haveRight)
			if got, want := status.Code(err), status.Code(tc.wantError); got != want {
				t.Errorf("ClosestIntersection() unexpected error. want: %s got: %s", want, got)
				return
			}
			if got != tc.want {
				t.Errorf("ClosestIntersection() mismatch want: %d got: %d", tc.want, got)
			}
		})
	}
}
func TestQuickestIntersection(t *testing.T) {
	tests := []struct {
		desc      string
		haveLeft  []string
		haveRight []string
		want      int
		wantError error
	}{
		{
			desc:      "Provided example: 1",
			haveLeft:  []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			haveRight: []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			want:      610,
		},
		{
			desc:      "Provided example: 2",
			haveLeft:  []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R5"},
			haveRight: []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			want:      410,
		},
		{
			desc:      "No left items",
			haveRight: []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			wantError: status.Error(codes.FailedPrecondition, "no left items"),
		},
		{
			desc:      "No right items",
			haveLeft:  []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R5"},
			wantError: status.Error(codes.FailedPrecondition, "no right items"),
		},
		{
			desc:      "Invalid direction: invalid cardinalty",
			haveLeft:  []string{"R98", "Z9"},
			wantError: status.Error(codes.InvalidArgument, "invalid direction"),
		},
		{
			desc:      "Invalid direction: invalid direction",
			haveLeft:  []string{"R98", "Z9"},
			wantError: status.Error(codes.InvalidArgument, "invalid direction"),
		},
		{
			desc:      "no overlap of intersections",
			haveLeft:  []string{"R98", "U9"},
			haveRight: []string{"L98", "D9"},
			wantError: status.Error(codes.NotFound, "no overlap"),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := QuickestIntersection(tc.haveLeft, tc.haveRight)
			if got, want := status.Code(err), status.Code(tc.wantError); got != want {
				t.Errorf("QuickestIntersection() unexpected error. want: %s got: %s", want, got)
				return
			}
			if got != tc.want {
				t.Errorf("QuickestIntersection() mismatch want: %d got: %d", tc.want, got)
			}
		})
	}
}
