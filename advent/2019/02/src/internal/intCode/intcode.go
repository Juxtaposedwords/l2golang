package intcode

import (
	"github.com/google/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
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
	answer := make(chan *pair, 1)
	errs := make(chan error, 1)
	var wg sync.WaitGroup

	for i := 0; i < nounCeiling; i++ {
		for j := 0; j < verbCeiling; j++ {
			wg.Add(1)
			output := make([]int, len(input))
			copy(output, input)
			go func(noun, verb int, output []int) {
				output[1], output[2] = noun, verb
				resp, err := List(output)
				if err != nil {
				} else if resp[0] == target {
					answer <- &pair{noun: noun, verb: verb}
				}
				wg.Done()
				//	logger.Infof("%d %d output: %d",verb, noun, resp[0])
			}(i, j, output)
		}
	}
	wg.Wait()
	select {
	case err := <-errs:
		return 0, 0, err
	case r := <-answer:
		return r.noun, r.verb, nil
	default:
	}

	return 0, 0, status.Error(codes.NotFound, "unable to find a posible pair")

}

// List steps through the code performing mutations where necessary.
func List(input []int) ([]int, error) {
	output := make([]int, len(input))
	copy(output, input)
	for i := 0; i < len(output); i += 4 {
		operation, ok  := opMap[output[i]]
		switch {
		case operation == terminate:
			return output, nil
		case !ok:
			return nil, status.Error(codes.InvalidArgument, "incorrectly shaped")
		case len(input[i:]) < 4:
			logger.Infof("** i: %d  digit: %d opcode: %#v", i,output[i], operation)
			return nil, status.Error(codes.FailedPrecondition, "incorrect number of items to the right of operator")
		}

		first, second, target := output[i+1], output[i+2], output[i+3]
		if first > len(input) || second > len(input) || target > len(input){
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
