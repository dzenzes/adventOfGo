package main

import (
	"fmt"
	"log"

	"github.com/dmies/adventOfGo/filehandler"
)

func CountNumberofIncreasedMeasurements(numbers []int) (count int) {
	if len(numbers) <= 1 {
		count = 0
		return
	}
	prev := numbers[0]
	for _, measurement := range numbers[1:] {
		if measurement > prev {
			count++
		}
		prev = measurement
	}
	return
}

func ToSlidingWindows(numbers []int) (slidingWindows []int) {
	if len(numbers) < 3 {
		slidingWindows = []int{}
		return
	}

	slidingWindows = append([]int{numbers[0] + numbers[1] + numbers[2]}, ToSlidingWindows(numbers[1:])...)
	return
}

func CountNumberofIncreasedMeasurementsWithSlidingWindow(numbers []int) (count int) {
	slidingWindows := ToSlidingWindows(numbers)
	count = CountNumberofIncreasedMeasurements(slidingWindows)
	return
}

func main() {
	input, err := filehandler.ImportNumberPerLineList("./input.txt")
	if err != nil {
		log.Fatalf("Couldn't read input")
	}

	solution1 := CountNumberofIncreasedMeasurements(input)
	fmt.Printf("day 01, part1 %v\n", solution1)
	solution2 := CountNumberofIncreasedMeasurementsWithSlidingWindow(input)
	fmt.Printf("day 01, part2 %v\n", solution2)
}
