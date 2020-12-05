package main

import (
	"errors"
	"fmt"
	"github.com/dmies/adventOfGo/filehandler"
	"log"
	"math"
	"sort"
)

// CalculateRange takes the current min and max and calculates the new ones by the given letter
func CalculateRange(letter string, min int, max int) (int, int, error) {
	difference := int(math.Trunc(float64(max-min) / 2))
	if letter == "F" || letter == "L" {
		return min, min + difference, nil
	} else if letter == "B" || letter == "R" {
		return max - difference, max, nil
	}
	return -1, -1, errors.New("CalculateRange(): unsupported letter " + letter)
}

// GetRow takes the boarding pass and returns the row encoded in it
func GetRow(boardingPass string) (int, error) {
	rowData := boardingPass[0:7]
	min := 0
	max := 127
	for _, letter := range rowData {
		minimum, maximum, err := CalculateRange(string(letter), min, max)
		if err != nil {
			return -1, err
		}
		min = minimum
		max = maximum
	}
	return min, nil
}

// GetSeat takes the boarding pass and returns the seat encoded in it
func GetSeat(boardingPass string) (int, error) {
	rowData := boardingPass[len(boardingPass)-3:]
	min := 0
	max := 7
	for _, letter := range rowData {
		minimum, maximum, err := CalculateRange(string(letter), min, max)
		if err != nil {
			return -1, err
		}
		min = minimum
		max = maximum
	}
	return min, nil
}

func GetSeatId(boardingPass string) (int, error) {
	row, err := GetRow(boardingPass)
	if err != nil {
		return -1, err
	}
	seat, err := GetSeat(boardingPass)
	if err != nil {
		return -1, err
	}
	return row*8 + seat, nil
}

// GetHighestSeatId finds the biggest seatId in the given list of boarding passes
func GetHighestSeatId(boardingPasses []string) (int, error) {
	seatId := -1
	for _, boardingPass := range boardingPasses {
		id, err := GetSeatId(boardingPass)
		if err != nil {
			return -1, err
		}
		if id > seatId {
			seatId = id
		}
	}
	return seatId, nil
}

func GetSortedSeatIds(boardingPasses []string) ([]int, error) {
	result := make([]int, 0)
	for _, boardingPass := range boardingPasses {
		id, err := GetSeatId(boardingPass)
		if err != nil {
			return []int{}, err
		}
		result = append(result, id)
	}
	sort.Ints(result)
	return result, nil
}

func GetMissingIds(ids []int) ([]int, error) {
	missing := make([]int, 0)
	lastId := ids[0]
	for _, id := range ids[1:] {
		if lastId+1 != id {
			missing = append(missing, lastId+1)
		}
		lastId = id
	}
	return missing, nil
}

func FindMySeat(boardingPasses []string) ([]int, error) {
	ids, err := GetSortedSeatIds(boardingPasses)
	if err != nil {
		return []int{}, err
	}
	return GetMissingIds(ids)
}

func main() {
	boardingPasses, err := filehandler.ImportSringList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution1, err := GetHighestSeatId(boardingPasses)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day 05, part1 %v\n", solution1)

	solution2, err := FindMySeat(boardingPasses)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day 05, part2 %v\n", solution2)

}
