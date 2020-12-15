package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dmies/adventOfGo/filehandler"
)

/*
Supported operations for the handheld
*/
const (
	Accumulate  = "acc"
	Jump        = "jmp"
	NoOperation = "nop"
)

// Instruction has an operation and an argument
type Instruction struct {
	operation string
	argument  int
}

// HandheldState holds accumulator and pointer of the Handheld
type HandheldState struct {
	accumulator       int
	pointer           int
	processedPointers []int
}

// contains checks if a string is present in a slice
func (h HandheldState) contains(pointer int) bool {
	for _, visited := range h.processedPointers {
		if visited == pointer {
			return true
		}
	}

	return false
}

// Parse takes a list of instructions and parses them to []Instruction
func Parse(input []string) ([]Instruction, error) {
	result := make([]Instruction, 0)
	for _, line := range input {
		instruction := strings.Split(line, " ")
		operation := instruction[0]
		argument, err := strconv.Atoi(instruction[1])
		if err != nil {
			return nil, err
		}
		result = append(result, Instruction{operation, argument})
	}
	return result, nil
}

// ProcessInstruction takes an instruction and HandheldState and calculats the next HandheldState if the operation is supported. Returns an error otherwise
func ProcessInstruction(instruction Instruction, state HandheldState) (HandheldState, error) {
	accumulator := state.accumulator
	pointer := state.pointer
	processedPointers := append(state.processedPointers, state.pointer)
	switch instruction.operation {
	case Accumulate:
		accumulator += instruction.argument
		pointer++
	case Jump:
		pointer += instruction.argument
	case NoOperation:
		pointer++
	default:
		return HandheldState{}, errors.New("ProcessInstruction() cannot handle instruction " + instruction.operation)
	}

	return HandheldState{accumulator, pointer, processedPointers}, nil
}

// Process processes the given Instructions and returns the final HandheldState. If there is an infinite loop detected an error is returned
func Process(instructions []Instruction) (HandheldState, error) {
	state := HandheldState{0, 0, []int{}}
	for state.pointer <= len(instructions)-1 {
		if state.contains(state.pointer) {
			return state, errors.New("instructions got stuck in infinite loop")
		}
		instruction := instructions[state.pointer]
		newState, err := ProcessInstruction(instruction, state)
		if err != nil {
			return HandheldState{}, err
		}
		state = newState
	}
	return state, nil
}

// GetInstructionsToSwitch returns all indices of Jump and NoOperation instructions
func GetInstructionsToSwitch(instructions []Instruction) []int {
	result := make([]int, 0)
	for i, instruction := range instructions {
		if instruction.operation == NoOperation || instruction.operation == Jump {
			result = append(result, i)
		}
	}
	return result
}

// CreateInstructionsCandidate clones the given instructions and switches the instruction at the given position from jmp to nop / nop to jmp
func CreateInstructionsCandidate(pointerToUpdate int, instructions []Instruction) []Instruction {
	result := append(instructions[:0:0], instructions...)
	instruction := instructions[pointerToUpdate]
	if instruction.operation == NoOperation {
		result[pointerToUpdate] = Instruction{Jump, instruction.argument}
	} else if instruction.operation == Jump {
		result[pointerToUpdate] = Instruction{NoOperation, instruction.argument}
	}
	return result
}

// FindCorrectInstructionsAndProcess finds all possible instructions and checks if the Handheld can process it without loop
func FindCorrectInstructionsAndProcess(instructions []Instruction) (HandheldState, error) {
	possibleInstructionsToSwitch := GetInstructionsToSwitch(instructions)

	for _, candidate := range possibleInstructionsToSwitch {
		instructionsCandidate := CreateInstructionsCandidate(candidate, instructions)
		result, err := Process(instructionsCandidate)
		if err == nil {
			// no need to handle errors. Just try with next instruction
			return result, nil
		}
	}

	return HandheldState{}, errors.New("couldn't find working instructions")
}

func main() {
	instructionList, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	instructions, err := Parse(instructionList)
	if err != nil {
		log.Fatal(err)
	}
	solution1, err := Process(instructions)
	if err == nil {
		log.Fatal("process should thrown an error because it detected an infinite loop")
	}
	fmt.Printf("day 08, part1 %v\n", solution1.accumulator)

	solution2, err := FindCorrectInstructionsAndProcess(instructions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day 08, part2 %v\n", solution2.accumulator)
}
