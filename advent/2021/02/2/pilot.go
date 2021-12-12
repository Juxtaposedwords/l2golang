package pilot

import (
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Navigate(instructions []string) (int, error) {
	horizontal, depth := 0, 0
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
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		default:
			return 0, status.Errorf(codes.InvalidArgument, "%q is not a recognized command. Recognized commands are forward, down, and up. ", operator)
		}
	}
	return horizontal * depth, nil
}
