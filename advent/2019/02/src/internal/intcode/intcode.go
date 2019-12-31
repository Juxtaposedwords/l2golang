package intcode

import (
	"github.com/google/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	nounCeiling = 100
	verbCeiling = 100
)

type pair struct {
	verb int
	noun int
}
type opCode int

const (
	unknown opCode = iota
	add
	multiply
	terminate opCode = 99
)

var opMap = map[int]opCode{
	1:  add,
	2:  multiply,
	99: terminate,
}

// BrutePair uses brute force to find all possible combinations of noun and verbs.
func BrutePair(input []int, target int) (int, int, error) {

	for i := 0; i < nounCeiling; i++ {
		for j := 0; j < verbCeiling; j++ {
			input[1], input[2] = j,i
			resp, err := list(input)
			if err != nil {
				return 0, 0, err
			} else if resp[0] == target {
				return j,i, nil
			}
		}
	}

	return 0, 0, status.Error(codes.NotFound, "unable to find a posible pair")

}

// list steps through the code performing mutations as opcode instruct.
func list(input []int) ([]int, error) {
	output := make([]int, len(input))
	copy(output, input)
	for i := 0; i < len(output); i += 4 {
		operation, ok := opMap[output[i]]
		switch {
		case operation == terminate:
			return output, nil
		case !ok:
			return nil, status.Error(codes.InvalidArgument, "incorrectly shaped")
		case len(input[i:]) < 4:
			logger.Infof("** i: %d  digit: %d opcode: %#v", i, output[i], operation)
			return nil, status.Error(codes.FailedPrecondition, "incorrect number of items to the right of operator")
		}

		first, second, target := output[i+1], output[i+2], output[i+3]
		if first > len(input) || second > len(input) || target > len(input) {
			return nil, status.Error(codes.InvalidArgument, "opcode target out of bounds")
		}
		switch operation {
		case add:
			output[target] = output[first] + output[second]
		case multiply:
			output[target] = output[first] * output[second]
		}
	}
	return nil, status.Error(codes.FailedPrecondition, "int list was not terminated with a 99 opcode")
}
