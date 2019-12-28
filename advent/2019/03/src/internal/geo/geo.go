package geo

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"strconv"
	"strings"
)

type location struct {
	x    int
	y    int
	step int
}
type intersection struct {
	left  *location
	right *location
}
type cardinality int

const (
	up cardinality = iota + 1
	down
	left
	right
)

var (
	directionRE = regexp.MustCompile(`^[UDRL]\d+$`)
)

// QuickestIntersection gives the closest intersection of the two list of coordinates.
func QuickestIntersection(left, right []string) (int, error) {
	leftLocations, err := getLocations(left)
	if err != nil {
		return 0, err
	}
	rightLocations, err := getLocations(right)
	if err != nil {
		return 0, err
	}
	leftVisits := map[string]*location{}
	for _, loc := range leftLocations {
		leftVisits[fmt.Sprintf("%03d%03d", loc.x, loc.y)] = loc
	}

	var intersections []*intersection
	for _, rightLoc := range rightLocations {
		leftLoc, ok := leftVisits[fmt.Sprintf("%03d%03d", rightLoc.x, rightLoc.y)]
		if !ok {
			continue
		}
		intersections = append(intersections,
			&intersection{
				left:  leftLoc,
				right: rightLoc,
			})
	}

	var steps int
	for _, inst := range intersections {
		instanceSteps := inst.left.step + inst.right.step
		if steps == 0 || instanceSteps < steps {
			steps = instanceSteps
		}

	}
	if steps == 0 {
		return 0, status.Error(codes.NotFound, "no intersections found")
	}

	return steps, nil
}

// ClosestIntersection gives the closest intersection of the two list of coordinates.
func ClosestIntersection(left, right []string) (int, error) {
	leftLocations, err := getLocations(left)
	if err != nil {
		return 0, err
	}
	rightLocations, err := getLocations(right)
	if err != nil {
		return 0, err
	}
	leftVisits := map[string]bool{}
	for _, loc := range leftLocations {
		leftVisits[fmt.Sprintf("%03d%03d", loc.x, loc.y)] = true
	}

	var lowestDistance int
	for _, loc := range rightLocations {
		if !leftVisits[fmt.Sprintf("%03d%03d", loc.x, loc.y)] {
			continue
		}
		if lowestDistance == 0 || taxiDistance(loc) < lowestDistance {
			lowestDistance = taxiDistance(loc)
		}
	}
	if lowestDistance == 0 {
		return 0, status.Error(codes.NotFound, "no intersections found")
	}

	return lowestDistance, nil
}
func taxiDistance(loc *location) int {
	return absVal(loc.x) + absVal(loc.y)
}

func absVal(input int) int {
	if input < 0 {
		return input * -1
	}
	return input
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}

func getLocations(directions []string) ([]*location, error) {
	if len(directions) == 0 {
		return nil, status.Error(codes.FailedPrecondition, "empty slice provided")
	}
	locs := map[string]*location{}
	step, x, y := 0, 0, 0
	for _, instruction := range directions {
		dir, dist, err := parse(instruction)
		if err != nil {
			return nil, err
		}
		for i := 0; i < dist; i++ {
			step++
			switch dir {
			case up:
				y++
			case down:
				y--
			case left:
				x--
			case right:
				x++
			}

			if _, ok := locs[fmt.Sprintf("%03d%03d", x, y)]; !ok {
				locs[fmt.Sprintf("%03d%03d", x, y)] = &location{x: x, y: y, step: step}
			}
		}
	}
	var output []*location
	for _, value := range locs {
		output = append(output, value)
	}
	return output, nil
}

func parse(input string) (cardinality, int, error) {
	if !directionRE.MatchString(input) {
		return 0, 0, status.Error(codes.InvalidArgument, "input does not match a valid direction")
	}
	characters := strings.Split(input, "")
	distance, _ := strconv.Atoi(strings.Join(characters[1:], ""))
	switch characters[0] {
	case "U":
		return up, distance, nil
	case "D":
		return down, distance, nil
	case "L":
		return left, distance, nil
	case "R":
		return right, distance, nil
	}
	return 0, 0, nil
}
