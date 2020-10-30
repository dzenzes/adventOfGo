package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	var strNums []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strNums = strings.Split(line, ",")
		break
	}

	nums := make([]int, 0)
	for _, thisStrNum := range strNums {
		thisNum, err := strconv.Atoi(strings.TrimSpace(thisStrNum))
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, thisNum)
	}

	return nums
}

const (
	// ADD is the operation that adds the next two params and stores the result at poistion 3
	ADD = 1
	// MULTIPLY is the operation that multiplies the next two params and stores the result at poistion 3
	MULTIPLY = 2
	// STOP is the operation that stops everything
	STOP = 99
)

// IntComputer gets a program...
func IntComputer(program []int) []int {
	pointer := 0
	isRunning := true
	for isRunning {
		opcode := program[pointer]
		switch opcode {
		case ADD:
			pointer1 := program[pointer+1]
			pointer2 := program[pointer+2]
			pointerOut := program[pointer+3]
			program[pointerOut] = program[pointer1] + program[pointer2]
			pointer += 4
		case MULTIPLY:
			pointer1 := program[pointer+1]
			pointer2 := program[pointer+2]
			pointerOut := program[pointer+3]
			program[pointerOut] = program[pointer1] * program[pointer2]
			pointer += 4
		case STOP:
			isRunning = false
		default:
			log.Fatal("Unsupported Instruction")
		}
	}
	return program
}

func clone(program []int) []int {
	x := make([]int, len(program))
	copy(x, program)
	return x
}

func part1(program []int) int {
	clonedProgram := clone(program)
	clonedProgram[1] = 12
	clonedProgram[2] = 2

	return IntComputer(clonedProgram)[0]
}

func part2(program []int) int {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			clonedProgram := clone(program)
			clonedProgram[1] = i
			clonedProgram[2] = j
			finalProgram := IntComputer(clonedProgram)
			if finalProgram[0] == 19690720 {
				return 100*i + j
			}
		}
	}
	return 0
}

func main() {
	program := parse("./input.txt")
	fmt.Printf("Part1: %d \n", part1(program))
	fmt.Printf("Part2: %d \n", part2(program))

}
