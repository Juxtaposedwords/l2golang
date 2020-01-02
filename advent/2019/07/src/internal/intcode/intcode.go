package intcode

import (
	"github.com/google/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"internal/operations"
)

type machine struct {
	software   []int
	softwareIndex  int
	inputs  []int
	inputIndex int
	output []int
	verbose bool
}

func (m machine) oneParam(input *operations.InstructionSet) (int, error) {
	if len(m.software[m.softwareIndex:]) < 1 {
		return 0, status.Error(codes.Internal, "out of bounds acccess requested")
	}
	resp := m.software[m.softwareIndex+1]
	if input.First == operations.Position {
		resp = m.software[resp]
	}
	return resp, nil
}
func (m machine) twoParams(input *operations.InstructionSet) (int, int, error) {
	if len(m.software[m.softwareIndex:]) < 3 {
		return 0, 0, status.Error(codes.Internal, "out of bounds acccess requested")
	}
	first, second := m.software[m.softwareIndex+1], m.software[m.softwareIndex+2]
	if input.First == operations.Position {
		first = m.software[first]
	}
	if input.Second == operations.Position {
		second = m.software[second]
	}
	return first, second, nil
}
func (m machine) threeParams(input *operations.InstructionSet) (int, int, int, error) {
	if len(m.software[m.softwareIndex:]) < 4 {
		return 0, 0, 0, status.Error(codes.Internal, "out of bounds acccess requested")
	}
	first, second, third := m.software[m.softwareIndex+1], m.software[m.softwareIndex+2], m.software[m.softwareIndex+3]
	if input.First == operations.Position {
		first = m.software[first]
	}
	if input.Second == operations.Position {
		second = m.software[second]
	}

	return first, second, third, nil
}
func (m *machine) copy(input *operations.InstructionSet) error {
	// we only want the exact value as this is the target
	input.First = operations.Immediate
	target, err := m.oneParam(input) 
	if err != nil {
		return err
	}
	if m.inputIndex >= len(m.inputs) {
		return status.Error(codes.Internal, "copy() out of bounds error")
	} 
	m.software[target] = m.inputs[m.inputIndex]
	m.inputIndex++
	m.softwareIndex += 2
	return nil
}
func (m *machine) print(input *operations.InstructionSet) error {
	val, err := m.oneParam(input)
	if err != nil {
		return err
	}
	m.output = append(m.output, val)
	m.softwareIndex += 2
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
		m.software[target] = first * second
	case operations.Add:
		m.software[target] = first + second
	}
	m.softwareIndex += 4
	return nil
}
func (m *machine) multiply(input *operations.InstructionSet) error {
	first, second, target, err := m.threeParams(input)
	if err != nil {
		return err
	}
	m.software[target] = first * second
	m.softwareIndex += 4
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
		m.softwareIndex = target
	default:
		m.softwareIndex += 3
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
	m.software[target] = output
	m.softwareIndex += 4
	return nil
}

func (m *machine) advance() (bool, error) {
	instruction, err := operations.Parse(m.software[m.softwareIndex])
	if err != nil {
		return false, err
	}
	if m.verbose {
		logger.Infof("\nadvance(%#v): %#v->%#v",instruction.Operation, m.software[:m.softwareIndex ], m.software[m.softwareIndex :])
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
func process(software []int, inputs []int) ([]int, error) {
	mach := &machine{
		software:  software,
		inputs: inputs,
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

// DiagnosticCode steps through the code performing mutations as opcode instruct.
func DiagnosticCode(software []int, inputInstruction []int) (int, error) {
	machineOutput, err := process(software, inputInstruction)
	if err != nil {
		return 0, err
	}
	checks, diagnosticCode := machineOutput[:len(machineOutput)-1], machineOutput[len(machineOutput)-1]
	for _, e := range checks {
		if e != 0 {
			return 0, status.Errorf(codes.Internal, "non 0 value before execution")
		}
	}
	return diagnosticCode, nil
}