package bingo

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type board struct {
	grid    [5][5]int
	called  map[int]bool
	winning bool
	last    int
}

func new(input [5][5]int) *board {
	return &board{
		grid:   input,
		called: map[int]bool{},
	}
}

func (b *board) check(input int) (bool, error) {
	b.called[input] = true
	b.last = input

	// could be two nested for loops, but this way allows for dianogal supprot in the future(TM), is easier to troubleshoot, and is funnier.
	switch {
	// Columns
	case b.called[b.grid[0][0]] && b.called[b.grid[1][0]] && b.called[b.grid[2][0]] && b.called[b.grid[3][0]] && b.called[b.grid[4][0]]:
	case b.called[b.grid[0][1]] && b.called[b.grid[1][1]] && b.called[b.grid[2][1]] && b.called[b.grid[3][1]] && b.called[b.grid[4][1]]:
	case b.called[b.grid[0][2]] && b.called[b.grid[1][2]] && b.called[b.grid[2][2]] && b.called[b.grid[3][2]] && b.called[b.grid[4][2]]:
	case b.called[b.grid[0][3]] && b.called[b.grid[1][3]] && b.called[b.grid[2][3]] && b.called[b.grid[3][3]] && b.called[b.grid[4][3]]:
	case b.called[b.grid[0][4]] && b.called[b.grid[1][4]] && b.called[b.grid[2][4]] && b.called[b.grid[3][4]] && b.called[b.grid[4][4]]:
	// Rows
	case b.called[b.grid[0][0]] && b.called[b.grid[0][1]] && b.called[b.grid[0][2]] && b.called[b.grid[0][3]] && b.called[b.grid[0][4]]:
	case b.called[b.grid[1][0]] && b.called[b.grid[1][1]] && b.called[b.grid[1][2]] && b.called[b.grid[1][3]] && b.called[b.grid[1][4]]:
	case b.called[b.grid[2][0]] && b.called[b.grid[2][1]] && b.called[b.grid[2][2]] && b.called[b.grid[2][3]] && b.called[b.grid[2][4]]:
	case b.called[b.grid[3][0]] && b.called[b.grid[3][1]] && b.called[b.grid[3][2]] && b.called[b.grid[3][3]] && b.called[b.grid[3][4]]:
	case b.called[b.grid[4][0]] && b.called[b.grid[4][1]] && b.called[b.grid[4][2]] && b.called[b.grid[4][3]] && b.called[b.grid[4][4]]:
	// Diagonal top left to bottom right.
	//	case b.called[b.grid[0][0]] && b.called[b.grid[1][1]] && b.called[b.grid[2][2]] && b.called[b.grid[3][3]] && b.called[b.grid[4][4]]:
	// Dianogal bottom left to top right.
	//case b.called[b.grid[4][0]] && b.called[b.grid[3][1]] && b.called[b.grid[2][2]] && b.called[b.grid[1][3]] && b.called[b.grid[0][4]]:
	default:
		return false, nil
	}
	b.winning = true

	return true, nil
}
func (b *board) score() (int, error) {
	if !b.winning {
		return 0, status.Errorf(codes.Internal, "%#v is not a winning board and so can't be scored", b)
	}

	var unmarkedSum int
	for _, entry := range sliceify(b.grid) {
		if !b.called[entry] {
			unmarkedSum += entry
		}

	}
	return unmarkedSum * b.last, nil
}

func sliceify(board [5][5]int) []int {
	var output []int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			output = append(output, board[i][j])
		}
	}
	return output
}
