package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/dmies/adventOfGo/filehandler"
)

// FindNumbersThatSumTo2020 checks an expense report ([]int) if there are two numbers that sum up to 2020 and returns them
func FindNumbersThatSumTo2020(expenseReport []int) (int, int, error) {
	for i, x := range expenseReport {
		for j, y := range expenseReport {
			if i != j && x+y == 2020 {
				return x, y, nil

			}
		}
	}
	return 0, 0, errors.New("couldn't find numbers that sum up to 2020")
}

// FindThreeNumbersThatSumTo2020 checks an expense report ([]int) if there are three numbers that sum up to 2020 and returns them
func FindThreeNumbersThatSumTo2020(expenseReport []int) (int, int, int, error) {
	for i, x := range expenseReport {
		for j, y := range expenseReport {
			for k, z := range expenseReport {
				if i != j && i != k && j != k && x+y+z == 2020 {
					return x, y, z, nil

				}
			}
		}
	}
	return 0, 0, 0, errors.New("couldn't find three numbers that sum up to 2020")
}

// Part1 uses FindNumbersThatSumTo2020 to find the correct numbers and multiplies them
func Part1(expenseReport []int) (int, error) {
	x1, x2, err := FindNumbersThatSumTo2020(expenseReport)
	if err != nil {
		return 0, err
	}
	return x1 * x2, nil
}

// Part2 uses FindThreeNumbersThatSumTo2020 to find the correct numbers and multiplies them
func Part2(expenseReport []int) (int, error) {
	x1, x2, x3, err := FindThreeNumbersThatSumTo2020(expenseReport)
	if err != nil {
		return 0, err
	}
	return x1 * x2 * x3, nil
}

func main() {
	expenseReport, err := filehandler.ImportNumberPerLineList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution1, err := Part1(expenseReport)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day 01, part1 %v\n", solution1)
	solution2, err := Part2(expenseReport)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day 02, part1 %v\n", solution2)

}
