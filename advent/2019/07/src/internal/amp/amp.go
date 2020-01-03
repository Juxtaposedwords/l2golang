package intcode

import (
	"github.com/google/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"internal/operations"
)

// State of the running machine
type State int

// Unstarted, Running, and Finished represent the four states of the machine.
const (
	Unstarted State = iota + 1
	Running
	Broken
	Finished
)

type Machine struct {
	State         State
	software      []int
	softwareIndex int
	input         chan (int)
	output        chan (int)
	err           chan (error)
	finished      chan (bool)
	last          int
	status        string
	verbose       bool
}

func (m Machine) oneParam(input *operations.InstructionSet) (int, error) {
	if len(m.software[m.softwareIndex:]) < 1 {
		return 0, status.Error(codes.Internal, "out of bounds acccess requested")
	}
	resp := m.software[m.softwareIndex+1]
	if input.First == operations.Position {
		resp = m.software[resp]
	}
	return resp, nil
}
func (m Machine) twoParams(input *operations.InstructionSet) (int, int, error) {
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
func (m Machine) threeParams(input *operations.InstructionSet) (int, int, int, error) {
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
func (m *Machine) copy(input *operations.InstructionSet) error {
	// we only want the exact value as this is the target
	target, err := m.oneParam(input)
	if err != nil {
		return err
	}
	if target < 0 || target >= len(m.software) {
		return status.Error(codes.Internal, "copy() out of bounds error")
	}
	m.software[target] = <-m.input
	m.softwareIndex += 2
	return nil
}
func (m *Machine) print(input *operations.InstructionSet) error {
	val, err := m.oneParam(input)
	if err != nil {
		return err
	}
	m.output <- val
	m.last = val
	m.softwareIndex += 2
	return nil
}

func (m *Machine) compute(input *operations.InstructionSet) error {
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
func (m *Machine) multiply(input *operations.InstructionSet) error {
	first, second, target, err := m.threeParams(input)
	if err != nil {
		return err
	}
	m.software[target] = first * second
	m.softwareIndex += 4
	return nil
}
func (m *Machine) jump(input *operations.InstructionSet) error {
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
func (m *Machine) compare(input *operations.InstructionSet) error {
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

func (m *Machine) advance() (bool, error) {
	instruction, err := operations.Parse(m.software[m.softwareIndex])
	if err != nil {
		return false, err
	}
	if m.verbose {
		logger.Infof("\nadvance(%#v): %#v->%#v", instruction.Operation, m.software[:m.softwareIndex], m.software[m.softwareIndex:])
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

// Outputs connects the output of machine m to the input of tm.
func (m *Machine) Outputs(tm *Machine) {
	tm.input = m.output
}

// Start causes an amp machine to go from unstarted to running.
func (m *Machine) Start() error {
	if m.State != Unstarted {
		return status.Error(codes.FailedPrecondition, "machined already started")
	}
	m.State = Running
	var finished bool
	go func(m *Machine) {
		for !finished {
			var err error
			finished, err = m.advance()
			if err != nil {
				m.State = Broken
				m.status = err.Error()
				m.err <- err
				return
			}
		}
		m.finished <- true

	}(m)

	return nil
}

// Create makes a copy of the tape and returns the machine.
func create(software []int, finished chan bool, errChan chan error) *Machine {
	softwareCopy := make([]int, len(software))
	copy(softwareCopy, software)
	return &Machine{
		State: Unstarted,
		input: make(chan(int), 100),
		output: make(chan(int), 10),
		software: softwareCopy,
		finished: finished,
		err:      errChan,
	}
}

// CreateChained connects a series of amps who have connected output and input leading in a circle.
func createChained(software []int, amount int) ([]*Machine, chan bool, chan error) {
	if amount <= 0 {
		return nil, nil,nil
	}
	finishChan, errChan := make(chan bool, amount), make(chan error, amount)
	machines := []*Machine{create(software, finishChan,errChan)}
	for i := 1; i < amount; i++ {
		m := create(software, finishChan,errChan)
		m.input = machines[i-1].output
		machines = append(machines, m)
	}
	machines[0].input = machines[amount-1].output
	return machines, finishChan, errChan
}

// ChainedProcess creates an Amp machine for each phase and connects the input and output of each. The  last value issued by the machine associated with the last phase code is returned.
func ChainedProcess(software []int, phases []int, intialInput int) (int, error) {
	machines, finishChan, errChan := createChained(software, len(phases))
	for i, code := range phases {
		machines[i].input <- code
		if  i == 0 {
			machines[i].input <-intialInput
		}
		if err := machines[i].Start(); err != nil {
			return 0, err
		}
	}
	var finished int
	for finished < len(phases) {
		select {
		case <-finishChan:
			finished++
		case err := <- errChan:
			return 0, err
		}
	}
	return machines[len(machines)-1].last, nil
}
