package intcode

import (
	"github.com/google/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"internal/operations"
	"io/ioutil"
)

type machine struct {
	tape   []int
	index  int
	input  int
	output []int
}

func (m machine) oneParam(input *operations.InstructionSet) (int, error) {
	if len(m.tape[m.index:]) < 1 {
		return 0, status.Error(codes.Internal, "out of bounds acccess requested")
	}
	resp := m.tape[m.index+1]
	if input.First == operations.Position {
		resp = m.tape[resp]
	}
	return resp, nil
}
func (m machine) twoParams(input *operations.InstructionSet) (int, int, error) {
	if len(m.tape[m.index:]) < 3 {
		return 0, 0, status.Error(codes.Internal, "out of bounds acccess requested")
	}
	first, second := m.tape[m.index+1], m.tape[m.index+2]
	if input.First == operations.Position {
		first = m.tape[first]
	}
	if input.Second == operations.Position {
		second = m.tape[second]
	}
	return first, second, nil
}
func (m machine) threeParams(input *operations.InstructionSet) (int, int, int, error) {
	if len(m.tape[m.index:]) < 4 {
		return 0, 0, 0, status.Error(codes.Internal, "out of bounds acccess requested")
	}
	first, second, third := m.tape[m.index+1], m.tape[m.index+2], m.tape[m.index+3]
	if input.First == operations.Position {
		first = m.tape[first]
	}
	if input.Second == operations.Position {
		second = m.tape[second]
	}
	return first, second, third, nil
}
func (m *machine) copy(input *operations.InstructionSet) error {
	target := m.tape[m.index+1]

	m.tape[target] = m.input
	m.index += 2
	return nil
}
func (m *machine) print(input *operations.InstructionSet) error {
	val, err := m.oneParam(input)
	if err != nil {
		return err
	}
	m.output = append(m.output,val)
	m.index += 2
	return nil
}

func (m *machine) compute(input *operations.InstructionSet) error {
	first, second, target, err := m.threeParams(input)
	if err != nil {
		return err
	}
	if !(input.Operation == operations.Multiply || input.Operation == operations.Add) {
		return status.Error(codes.InvalidArgument, "bad instruction set")
	}
	switch input.Operation {
	case operations.Multiply:
		m.tape[target] = first * second
	case operations.Add:
		m.tape[target] = first + second
	}
	m.index += 4
	return nil
}
func (m *machine) multiply(input *operations.InstructionSet) error {
	first, second, target, err := m.threeParams(input)
	if err != nil {
		return err
	}
	m.tape[target] = first * second
	m.index += 4
	return nil
}
func (m *machine) jump(input *operations.InstructionSet) error {
	next, target, err := m.twoParams(input)
	if err != nil {
		return err
	}
	switch {
	case next == 0 && input.Operation == operations.JumpIfFalse:
		fallthrough
	case next != 00 && input.Operation == operations.JumpIfTrue:
		m.index = target
	default:
		m.index += 3
	}
	return nil
}
func (m *machine) compare(input *operations.InstructionSet) error {
	first, second, target, err := m.threeParams(input)
	if err != nil {
		return err
	}
	var output int

	switch {
	case first < second && input.Operation == operations.LessThan:
		fallthrough
	case first == second && input.Operation == operations.Equals:
		output = 1
	}
	m.tape[target] = output
	m.index += 2
	return nil
}

func (m *machine) advance() (bool, error) {
	instruction, err := operations.Parse(m.tape[m.index])
	if err != nil {
		return false, err
	}
	switch instruction.Operation {
	case operations.Add, operations.Multiply:
		return false, m.compute(instruction)
	case operations.Copy:
		return false, m.copy(instruction)
	case operations.Print:
		return false, m.print(instruction)
	case operations.JumpIfFalse, operations.JumpIfTrue:
		return false, m.jump(instruction)
	case operations.LessThan, operations.Equals:
		return false, m.compare(instruction)
	case operations.Terminate:
		return true, nil
	}
	return false, status.Errorf(codes.Internal, "machine not configured for opcode in : %#v", instruction)
}

// Process steps through the code performing mutations as opcode instruct.
func Process(input []int, inputInstruction int) ([]int, error) {
	logger.Init("LoggerExample", true, false, ioutil.Discard)
	mach := &machine{
		tape: input,
		input: inputInstruction,
	}
	var finished bool
	for !finished {
		var err error
		finished, err = mach.advance()
		if err != nil {
			return nil, err
		}
	}
	return mach.output, nil
}

