package intcomputer

import "log"

const (
	add      = 1
	multiply = 2
	input    = 3
	output   = 4
	stop     = 99
)

// IntComputer is a struct that holds all information needed in the IntComputer
type IntComputer struct {
	Memory  []int
	Input   []int
	Output  []int
	Pointer int
}

// OpCode holds the actual operation and all modes
type OpCode struct {
	Operation int
	modes     []int
}

func (computer IntComputer) getNextOpCode() OpCode {
	var parameterModes []int
	opCode := computer.Memory[computer.Pointer]
	remaining := opCode
	operation := opCode % 100
	remaining -= operation

	thirdMode := opCode / 10000
	remaining -= thirdMode * 10000
	secondMode := remaining / 1000
	remaining -= secondMode * 1000
	firstMode := remaining / 100

	switch opCode {
	case 3, 4: //only single parameter mode
		parameterModes = []int{firstMode}
	case 5, 6: //two parameter modes
		parameterModes = []int{firstMode, secondMode}
	case 1, 2, 7, 8: //three parameter modes
		parameterModes = []int{firstMode, secondMode, thirdMode}
	}

	return OpCode{Operation: operation, modes: parameterModes}
}

func (computer IntComputer) getValue(position int, modes []int) int {
	mode := modes[position-1]

	// position mode
	if mode == 0 {
		pos := computer.Memory[computer.Pointer+position]
		return computer.Memory[pos]
	} else if mode == 1 {
		// immediate mode
		return computer.Memory[computer.Pointer+position]
	} else {
		log.Fatal("Unsupported mode")
	}

}

// Process to calculate programs
func (computer IntComputer) Process() IntComputer {
	isRunning := true
	for isRunning {
		opcode := computer.getNextOpCode()
		switch opcode.Operation {
		case add:

			storePosition := computer.Memory[computer.Pointer+3]
			computer.Memory[storePosition] = computer.getValue(1, opcode.modes) + computer.getValue(1, opcode.modes)
			computer.Pointer += 4
		case multiply:
			pos1 := computer.Memory[computer.Pointer+1]
			pos2 := computer.Memory[computer.Pointer+2]
			storePosition := computer.Memory[computer.Pointer+3]
			computer.Memory[storePosition] = computer.Memory[pos1] * computer.Memory[pos2]
			computer.Pointer += 4
		case input:
			storePosition := computer.Memory[computer.Pointer+1]
			computer.Memory[storePosition] = computer.Input[0]
			computer.Input = computer.Input[1:]
			computer.Pointer += 2
		case output:
			position := computer.Memory[computer.Pointer+1]
			computer.Output = append(computer.Output, computer.Memory[position])
		case stop:
			isRunning = false
		}
	}
	return computer
}

// Create creates an IntComputer for a given program
func Create(program []int) IntComputer {
	return IntComputer{Memory: program, Input: make([]int, 0), Output: make([]int, 0), Pointer: 0}
}
