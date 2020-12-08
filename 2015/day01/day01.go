package main

import (
	"fmt"
	"github.com/dmies/adventOfGo/filehandler"
	"log"
)

func FindFloor(instruction string) int {
	result := 0
	for _, char := range instruction {
		if string(char) == "(" {
			result++
		} else if string(char) == ")" {
			result--
		}

	}
	return result
}

func FirstTimeInBasement(instruction string) int {
	floor := 0
	for i, char := range instruction {
		if string(char) == "(" {
			floor++
		} else if string(char) == ")" {
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return -1
}

func main() {
	instructions, err := filehandler.ImportSringList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	instruction := instructions[0]

	solution1 := FindFloor(instruction)

	fmt.Printf("day 01, part1 %v\n", solution1)

	solution2 := FirstTimeInBasement(instruction)

	fmt.Printf("day 01, part2 %v\n", solution2)
}
