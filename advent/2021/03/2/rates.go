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

func LifeSupport(input []string) (int, error) {
	c, err := C02(input)
	if err != nil {
		return 0, err
	}
	o, err := Oxygen(input)
	if err != nil {
		return 0, err
	}
	return c * o, nil
}
func C02(inputCodes []string) (int, error) {
	var validRows [][]string
	for _, entry := range inputCodes {
		validRows = append(validRows, strings.Split(entry, ""))
	}
	for i := 0; i < len(validRows[0]); i++ {
		var oneList, zeroList [][]string
		for _, row := range validRows {
			if row[i] == "1" {
				oneList = append(oneList, row)
			} else {
				zeroList = append(zeroList, row)
			}
		}
		if len(oneList) < len(zeroList) {
			validRows = oneList
		} else {
			validRows = zeroList
		}
		if len(validRows) == 1 {
			break
		}
		if len(validRows) < 1 {
			return 0, status.Error(codes.Internal, "no valid rows")
		}

	}
	if len(validRows) != 1 {
		return 0, status.Errorf(codes.Internal, "Found %d not one number", len(validRows))
	}
	c02, err := strconv.ParseInt(strings.Join(validRows[0], ""), 2, 64)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "failed to convert %q from string binary to decimal with error: %s", validRows[0], err)
	}
	return int(c02), nil
}
func Oxygen(inputCodes []string) (int, error) {
	var validRows [][]string
	for _, entry := range inputCodes {
		validRows = append(validRows, strings.Split(entry, ""))
	}
	for i := 0; i < len(validRows[0]); i++ {
		var oneList, zeroList [][]string
		for _, row := range validRows {
			if row[i] == "1" {
				oneList = append(oneList, row)
			} else {
				zeroList = append(zeroList, row)
			}
		}
		if len(oneList) < len(zeroList) {
			validRows = zeroList
		} else {
			validRows = oneList
		}
		if len(validRows) == 1 {
			break
		}
		if len(validRows) < 1 {
			return 0, status.Error(codes.Internal, "no valid rows")
		}

	}
	if len(validRows) != 1 {
		return 0, status.Errorf(codes.Internal, "Found %d not one number", len(validRows))
	}
	c02, err := strconv.ParseInt(strings.Join(validRows[0], ""), 2, 64)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "failed to convert %q from string binary to decimal with error: %s", validRows[0], err)
	}
	return int(c02), nil
}
