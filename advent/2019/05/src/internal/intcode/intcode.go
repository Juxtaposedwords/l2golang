package intcode

import (
	"internal/operations"

	"github.com/google/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
)

// Process steps through the code performing mutations as opcode instruct.
func Process(input []int, inputInstruction int) ([]int, error) {
	logger.Init("LoggerExample", true, false, ioutil.Discard)
	var output []int
	var opLength int
	for i := 0; i < len(input); i += opLength {
		instructions, err := operations.Parse(input[i])
		if err != nil {
			return nil, err
		}
		opLength = operations.CodeLength[instructions.Operation]
		logger.Infof("input[%d] = %d digits: %#v", i, input[i], input[i:])

		logger.Infof("oplength: %d\n", opLength)
		if len(input[i:]) < opLength {
			return nil, status.Error(codes.FailedPrecondition, "incorrectly sized int list")
		}
		switch instructions.Operation {
		case operations.Terminate:
			return output, nil
		case operations.Print:
			output = append(output, input[input[i+1]])
		default:
			if err := mutate(instructions, i, input, inputInstruction); err != nil {
				return nil, err
			}
		}
		//	logger.Infof("digits: %#v instructions: %#v\n", input, instructions)
	}
	return nil, status.Error(codes.FailedPrecondition, "int list was not terminated with a 99 opcode")
}

func mutate(instructions *operations.InstructionSet, index int, input []int, inputInstruction int) error {
	targetIndex := input[index+operations.CodeLength[instructions.Operation]-1]

	first, second := input[index+1], input[index+2]
	if instructions.First == operations.Position {
		first = input[first]
	}
	if instructions.Second == operations.Position {
		second = input[second]
	}
	switch instructions.Operation {
	case operations.Copy:
		input[targetIndex] = inputInstruction
	case operations.Multiply:
		input[targetIndex] = first * second
	case operations.Add:
		//	logger.Infof("adding:  digits %#v\n", input)

		input[targetIndex] = first + second
		//		logger.Infof("adding:  digits %#v\n", input)
	default:
		return status.Errorf(codes.FailedPrecondition, "invalid opcode: %#v", instructions)
	}
	return nil

}
