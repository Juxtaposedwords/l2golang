package operations

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

// Code tells us the operations which can be taken.
type Code int

/*	Add (3 parameters) - param c = param a + param b
    Multiply (3 parameters) - param c = param a * param b
    Copy(1 parameters) - write the value of param b to the item in index of param b
 	Print(1 parameter) - print/"output" the parameter given
	Terminate - required ending Code */
const (
	Add Code = iota + 1
	Multiply
	Copy
	Print
	JumpIfTrue
	JumpIfFalse
	LessThan
	Equals
	Terminate Code = 99
)

var codeMap = map[string]Code{
	"01": Add,
	"02": Multiply,
	"03": Copy,
	"04": Print,
	"05": JumpIfTrue,
	"06": JumpIfFalse,
	"07": LessThan,
	"08": Equals,
	"99": Terminate,
}

// CodeLength is a way to store some metadata about instructions.
var CodeLength = map[Code]int{
	Add:         4,
	Multiply:    4,
	Copy:        2,
	Print:       2,
	Terminate:   0,
	JumpIfTrue:  4,
	JumpIfFalse: 4,
	Equals:      4,
}

// Mode represents the two ways to write.
type Mode int

// Position - use  the associated parameter value as index
// Immediate - use the associated paramter value as a value
const (
	Position Mode = iota + 1
	Immediate
)

var modeMap = map[string]Mode{
	"0": Position,
	"1": Immediate,
}

//InstructionSet contains all the mappings of paramters input to modes.
type InstructionSet struct {
	Operation Code
	First     Mode
	Second    Mode
	Third     Mode
}

// Parse turns an operation Code into an output of an instruction set.
func Parse(input int) (*InstructionSet, error) {

	paddedInput := fmt.Sprintf("%05d", input)
	if len(paddedInput) > 5 {
		return nil, status.Error(codes.FailedPrecondition, "invalid integer length")
	}

	digits := strings.Split(paddedInput, "")
	opCode, ok := codeMap[fmt.Sprintf("%s%s", digits[3], digits[4])]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "invalid operation Code ('%s') provided", fmt.Sprintf("%s%s", digits[3], digits[4]))
	}

	first, firstOK := modeMap[digits[2]]
	second, secondOK := modeMap[digits[1]]
	third, thirdOK := modeMap[digits[0]]
	if !firstOK || !secondOK || !thirdOK {
		return nil, status.Error(codes.InvalidArgument, "invalid operation Mode provided")
	}

	resp := &InstructionSet{
		Operation: opCode,
		First:     first,
		Second:    second,
		Third:     third,
	}

	switch {
	case CodeLength[resp.Operation] == 4 && resp.Third == Immediate:
		fallthrough
	case resp.Operation == Copy && resp.First == Immediate:
		return nil, status.Error(codes.InvalidArgument, "target cannot be immediate")
	}
	return resp, nil

}
