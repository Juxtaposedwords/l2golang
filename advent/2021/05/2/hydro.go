package hydro

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Overlap(level int, input *bufio.Reader) (int, error) {
	coordinates, err := parse(input)
	if err != nil {
		return 0, err
	}

	coordFrequency := map[string]int{}
	for _, entry := range coordinates {
		if _, ok := coordFrequency[entry]; !ok {
			coordFrequency[entry] = 0
		}
		coordFrequency[entry]++
	}

	var output int
	for _, v := range coordFrequency {
		if v >= level {
			output += 1
		}
	}

	return output, nil
}

func parse(input *bufio.Reader) ([]string, error) {
	var output []string
	var rowInput []byte

	rowInput, _, err := input.ReadLine()
	for !errors.Is(err, io.EOF) {

		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to read a line of buffer: %s", err)
		}
		var points []string
		points, err = find(string(rowInput))
		if err != nil {
			return nil, err
		}

		output = append(output, points...)
		rowInput, _, err = input.ReadLine()
	}
	return output, nil
}

func find(input string) ([]string, error) {
	pieces := strings.Split(strings.TrimSpace(input), " -> ")
	if len(pieces) != 2 {
		return nil, status.Errorf(codes.Internal, "invalid format found on row, expecting `0,0 -> 0,0`, but got %s", input)
	}
	start, err := coords(pieces[0])
	if err != nil {
		return nil, err
	}

	end, err := coords(pieces[1])
	if err != nil {
		return nil, err
	}
	var output []string
	//	return output, nil

	for !(start[0] == end[0] && start[1] == end[1]) {
		output = append(output, fmt.Sprintf("%d|%#v", start[0], start[1]))
		if start[0] != end[0] {
			start[0] += diff(start[0], end[0])
		}
		if start[1] != end[1] {
			start[1] += diff(start[1], end[1])
		}
	}
	output = append(output, fmt.Sprintf("%d|%#v", end[0], end[1]))

	return output, nil
}
func diff(a, b int) int {
	if a > b {
		return -1
	} else {
		return 1
	}
}

func coords(input string) ([2]int, error) {
	pieces := strings.Split(input, ",")
	var output [2]int
	if len(pieces) != 2 {
		return output, status.Errorf(codes.Internal, "invalid format found for coordinate, expecting `0,0` format but got %s", input)
	}
	var err error
	output[0], err = strconv.Atoi(pieces[0])

	if err != nil {
		return output, status.Errorf(codes.Internal, "unable to parse x coord %s with :%s", pieces[0], err)
	}
	output[1], err = strconv.Atoi(pieces[1])
	if err != nil {
		return output, status.Errorf(codes.Internal, "unable to parse y coord %s with :%s", pieces[1], err)
	}

	return output, nil

}
