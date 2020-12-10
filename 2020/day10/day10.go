package main

import (
	"errors"
	"fmt"
	"log"
	"sort"

	"github.com/dmies/adventOfGo/filehandler"
)

// PrepareInput sorts the given adapters and adds the first adapter (0) and the last adapter (max adapter +3)
func PrepareInput(input []int) (adapters []int) {
	input = append(input, 0)
	adaptersSorted := (sort.IntSlice)(input)
	adaptersSorted.Sort()
	adapters = adaptersSorted
	builtInAdapter := adapters[len(adapters)-1] + 3
	adapters = append(adapters, builtInAdapter)
	return adapters
}

// GetNextAdapter checks the (sorted!) adapters for the next adapter that is bigger than the given inputJolts
func GetNextAdapter(inputJolts int, adapters []int) (int, error) {

	for _, jolt := range adapters {
		if jolt > inputJolts {
			if jolt-inputJolts > 3 {
				return -1, errors.New("no matching next adapter found")
			}

			return jolt, nil
		}
	}
	return -1, errors.New("no matching adapter found")
}

// GetDifferences checks the adapters and returns the difference in jolts by 1 and 3
func GetDifferences(adapters []int) (int, int, error) {
	jolts := adapters[0]
	idx := 1
	difference1 := 0
	difference3 := 0
	for idx < len(adapters) {
		nextJolts, err := GetNextAdapter(jolts, adapters)
		if err != nil {
			return -1, -1, err
		}

		if nextJolts-jolts == 1 {
			difference1++
		} else if nextJolts-jolts == 3 {
			difference3++
		}

		idx++
		jolts = nextJolts

	}
	return difference1, difference3, nil
}

// CountPossibleCombinations counts all possible adapter combinations. It takes the sorted (!) adapters as input
func CountPossibleCombinations(adapters []int, idx int, cache map[int]int) (arrangements int) {
	if idx == len(adapters)-1 {
		return 1
	}
	if cachedResult, ok := cache[idx]; ok {
		return cachedResult
	}
	for i, adapter := range adapters[idx+1:] {
		difference := adapter - adapters[idx]
		if difference <= 3 {
			arrangements += CountPossibleCombinations(adapters, idx+i+1, cache)
		} else {
			break
		}
	}
	cache[idx] = arrangements
	return arrangements

}

func main() {
	input, err := filehandler.ImportNumberList("./input.txt")
	if err != nil {
		log.Fatal("Couldn't read input")
	}
	adapters := PrepareInput(input)
	diff1, diff3, err := GetDifferences(adapters)
	if err != nil {
		log.Fatalf("GetDifferences returned an error %v", err)
	}
	fmt.Printf("day 10, part1 %v\n", diff1*diff3)

	solution2 := CountPossibleCombinations(adapters, 0, map[int]int{})

	fmt.Printf("day 10, part2 %v\n", solution2)
}
