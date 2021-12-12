package rates

import (
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type count struct {
	one  int
	zero int
}

// Forgive this terrible function names
func Compute(inputCodes []string) (int, error) {
	// Let's make an array for each digit.
	var digits []*count
	for i := 0; i < len(inputCodes[1]); i++ {
		digits = append(digits, &count{})
	}

	for _, inputCode := range inputCodes {
		for i, digit := range strings.Split(inputCode, "") {
			switch digit {
			case "0":
				digits[i].zero++
			case "1":
				digits[i].one++
			default:
				return 0, status.Errorf(codes.InvalidArgument, "%s is not a 1 or 0 . From code %q", digit, inputCode)
			}
		}
	}

	var gamma, epsilon []string
	for _, entry := range digits {
		if entry.one > entry.zero {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		} else {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		}
	}
	var gammaValue, epsilonValue int64
	var err error
	gammaValue, err = strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "failed to convert %q into a digit", strings.Join(gamma, ""))
	}
	epsilonValue, err = strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "failed to convert %q into a digit", strings.Join(epsilon, ""))
	}

	return int(gammaValue * epsilonValue), nil
}
