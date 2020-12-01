package main

import (
	"fmt"
	"log"

	"github.com/dmies/adventOfGo/2019/intcomputer"
	"github.com/dmies/adventOfGo/filehandler"
)

func part1(program []int) int {

	clone := make([]int, len(program))
	copy(clone, program)
	clone[1] = 12
	clone[2] = 2

	computer := intcomputer.Create(clone)

	return computer.Process().Memory[0]
}

func part2(program []int) int {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			clone := make([]int, len(program))
			copy(clone, program)
			clone[1] = noun
			clone[2] = verb
			computer := intcomputer.Create(clone)
			result := computer.Process().Memory[0]
			if result == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return -1
}

func main() {
	program, err := filehandler.ImportNumberList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part1: %d \n", part1(program))
	fmt.Printf("Part2: %d \n", part2(program))

}
