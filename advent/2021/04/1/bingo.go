package bingo

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WinningScore(input *bufio.Reader) (int, error) {
	line, _, err := input.ReadLine()
	if err != nil {
		return 0, status.Errorf(codes.InvalidArgument, "failed to read first line of buffer with: %s", err)
	}
	called := strings.Split(string(line), ",")
	var calledNumbers []int
	for i, number := range called {
		calledNumber, err := strconv.Atoi(strings.TrimSpace(number))
		if err != nil {
			return 0, status.Errorf(codes.Internal, "failed to parse called called number %d (zero based) with: %s ", i, err)
		}
		calledNumbers = append(calledNumbers, calledNumber)

	}

	allBoards, err := parse(input)

	if err != nil {
		return 0, err
	}

	for _, calledNumber := range calledNumbers {
		var winning []*board

		for _, singleBoard := range allBoards {
			ok, err := singleBoard.check(calledNumber)
			if err != nil {
				return 0, err
			}

			if ok {
				winning = append(winning, singleBoard)
			}
		}
		var winner int
		for _, b := range winning {
			score, err := b.score()
			if err != nil {
				return 0, err
			}
			if score > winner {
				winner = score
			}
		}
		if winner != 0 {
			return winner, nil
		}
	}
	return 0, nil
}

func parse(input *bufio.Reader) ([]*board, error) {
	var output []*board
	_, _, err := input.ReadLine()
	for !errors.Is(err, io.EOF) {
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to read a line of buffer: %s", err)
		}
		var grid [5][5]int
		// Load one board
		for i := 0; i < 5; i++ {
			rowInput, _, err := input.ReadLine()
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to read a row %d , %#v", i, rowInput)
			}
			grid[i], err = row(string(rowInput))
			if err != nil {
				return nil, err
			}
		}
		output = append(output, new(grid))
		_, _, err = input.ReadLine()
	}
	return output, nil
}
func row(s string, t *testing.T) ([5]int, error) {
	if len(s) != 14 {
		return [5]int{}, status.Errorf(codes.FailedPrecondition, "invalid length of a row found. Expected 14 got %d for %s", len(s), s)
	}
	var row []string
	for _, entry := range strings.Split(s, " ") {
		if entry != "" {
			row = append(row, strings.TrimSpace(entry))
		}
	}

	var output [5]int

	for i := 0; i < 5; i++ {
		var err error
		output[i], err = strconv.Atoi(strings.TrimSpace(row[i]))
		if err != nil {
			return output, status.Errorf(codes.Internal, "failed to parse row(%#v) entry: %q at column %d entry with: %s", s, row[i], i, err)
		}
	}
	return output, nil
}
