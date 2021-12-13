package bingo

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type board struct {
	entries [5][5]string
	
}

func WinningScore(input *bufio.Reader, t *testing.T) (int, error) {
	line, _, err := input.ReadLine()
	if err != nil {
		return 0, status.Errorf(codes.InvalidArgument, "failed to read first line of buffer with: %s", err)
	}
	called := strings.Split(string(line), ",")
	for i := range called {
		if len(called[i]) == 1 {
			called[i] = fmt.Sprintf(" %s", called[i])
		}
	}
	t.Logf("numbers: %#v", called)

	b, err := boards(input, t)
	if err != nil {
		t.Logf("%s", err)

		return 0, err
	}
	t.Logf("board: %#v", b)

	return 0, nil
}
func winningValue(board [5][5]string, called map[string]bool) (int, bool) {

	for i := 0; i < 5; i++ {
		var win bool
		for j := 0; j < 5; j++ {

		}
	}
	return 0, false
}
func boards(input *bufio.Reader, t *testing.T) ([][5][5]string, error) {
	output := [][5][5]string{}
	for i := 0; ; i++ {
		_, _, err := input.ReadLine()
		if errors.Is(err, io.EOF) {
			return output, nil
		}
		if err != nil {
			return output, status.Errorf(codes.Internal, "failed to read a line of buffer: %s", err)
		}
		var board [5][5]string
		for j := 0; j < 5; j++ {
			rawLine, _, err := input.ReadLine()
			if err != nil {
				return output, status.Errorf(codes.Internal, "failed to read a row %d from board %d", j, i)
			}
			board[i], err = row(string(rawLine))
			if err != nil {
				return output, err
			}
			t.Logf("board %d row %d :%#v", i, j, board[i])
		}

	}
}
func row(s string) ([5]string, error) {
	if len(s) != 14 {
		return [5]string{}, status.Errorf(codes.FailedPrecondition, "invalid length of a row found. Expected 14 got %d for %s", len(s), s)
	}
	row := [5]string{
		string(s[0:2]), string(s[3:5]), string(s[6:8]), string(s[9:11]), string(s[12:14]),
	}
	return row, nil
}
