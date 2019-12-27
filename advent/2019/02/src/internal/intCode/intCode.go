package intCode

import "google.golang.org/grpc/status"

import "google.golang.org/grpc/codes"



// List steps through the code performing mutations where necessary.
func List(input []int) ([]int, error){
	output := make([]int, len(input))
	copy(output, input)
	opcodes := map[int]bool{
		1:true,
		2:true,
		99:true}

	for i := 0; i < len(output) ;i= i+4 {
		opcode, ok := output[i], opcodes[output[i]]
		switch {
		case opcode == 99 :
			return output, nil
		case !ok:
			return nil, status.Error(codes.InvalidArgument, "incorrectly shaped")
		case len(output[i:]) <= 4 :
			return nil, status.Error(codes.InvalidArgument, "incorrectly shaped")
		}

		first, second, target  := output[i+1],output[i+2], output[i+3]
		switch {
		case first >= len(output):
			return nil, status.Errorf(codes.FailedPrecondition, "index %d is greater than the length of input(length: %d", first,len(input))
		case second >= len(output):
			return nil, status.Errorf(codes.FailedPrecondition, "index %d is greater than the length of input(length: %d", second, len(input))
		case target >= len(output):
			return nil, status.Errorf(codes.FailedPrecondition, "index %d is greater than the length of input(length: %d", second, len(input))
		}

		switch opcode {
		case 1:
			output[target] = output[first] + output[second]
		case 2:
			output[target] = output[first] * output[second]

		}


	}

	return nil, status.Error(codes.Internal, "int list was not terminated with a 99 opcode")
}