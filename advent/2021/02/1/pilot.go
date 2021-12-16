package pilot

import (
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Navigate(instructions []string) (int, error) {
	x, y := 0, 0
	for _, instruction := range instructions {
		pieces := strings.Split(instruction, " ")
		if len(pieces) != 2 {
			return 0, status.Errorf(codes.FailedPrecondition, "%q does not contain a space seperated {command} {value} statement.", instruction)
		}
		operator, rawValue := pieces[0], pieces[1]
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			return 0, status.Errorf(codes.InvalidArgument, "%s is not a valid number in %q", rawValue, instruction)
		}
		switch operator {
		case "forward":
			x += value
		case "down":
			y += value
		case "up":
			y -= value
		default:
			return 0, status.Errorf(codes.InvalidArgument, "%q is not a recognized command. Recognized commands are forward, down, and up. ", operator)
		}
	}
	return x * y, nil
}
