package main

import (
	"errors"
	"fmt"
	"log"
	"sort"

	"github.com/dmies/adventOfGo/filehandler"
)

// FindSum checks if the given numberToCheck is a sum of two numbers in previousNumbers
func FindSum(numberToCheck int, previousNumbers []int) bool {
	for i, x := range previousNumbers {
		for _, y := range previousNumbers[i+1:] {
			if x+y == numberToCheck {
				return true
			}
		}
	}
	return false
}

// FindFirstWrongNumber finds the first number in input that is no sum of two numbers of the previous preambleSize numbers
func FindFirstWrongNumber(input []int, preambleSize int) int {
	for i, value := range input[preambleSize:] {
		preamble := input[i : preambleSize+i]
		if !FindSum(value, preamble) {
			return value
		}
	}
	return -1
}

// FindContiguousSetOfNumbersThatSumUpTo looks for a contiguous set of numbers in input that sum up to searched and returns them
func FindContiguousSetOfNumbersThatSumUpTo(searched int, input []int) ([]int, error) {
	var result []int
	for i, value := range input {
		sum := value
		j := i + 1
		result = []int{value}
		for sum < searched && j < len(input) {
			sum += input[j]
			result = append(result, input[j])
			j++
		}
		if sum == searched {
			return result, nil
		}
	}
	return nil, errors.New("FindContiguousSetOfNumbersThatSumUpTo() couldn't find suitable set")
}

// GetMinAndMaxFromList returns the min and max in the given list
func GetMinAndMaxFromList(input []int) (int, int) {
	copy := append([]int(nil), input...)
	sort.Ints(copy)
	return copy[0], copy[len(copy)-1]
}

// FindEncryptionWeakness finds the set of contiguous numbers that sum up to the illegalNumber, gets the min and max off this list and returns the sum of them
func FindEncryptionWeakness(illegalNumber int, input []int) (int, error) {
	listToCheck, err := FindContiguousSetOfNumbersThatSumUpTo(illegalNumber, input)
	if err != nil {
		return -1, err
	}
	min, max := GetMinAndMaxFromList(listToCheck)
	return min + max, nil
}

func main() {
	numbers, err := filehandler.ImportNumberList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution1 := FindFirstWrongNumber(numbers, 25)

	fmt.Printf("day 09, part1 %v\n", solution1)

	solution2, err := FindEncryptionWeakness(solution1, numbers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day 09, part2 %v\n", solution2)
}
