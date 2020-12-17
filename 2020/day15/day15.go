package main

import "fmt"

func numberWasNotSpokenBefore(number int, cache map[int][]int) (numberWasNotSpoken bool) {

	if turnsNumberWasSpoken, ok := cache[number]; ok {
		if len(turnsNumberWasSpoken) == 1 {
			numberWasNotSpoken = true
		}
	}
	return
}

func WriteToCache(key int, value int, cache map[int][]int) map[int][]int {
	if values, ok := cache[key]; ok {
		lastValueInCache := values[len(values)-1]
		cache[key] = []int{lastValueInCache, value}
	} else {
		cache[key] = []int{value}
	}

	return cache
}

// GetSpokenNumber returns the spokenNumber for a specified turn
func GetSpokenNumber(numbers []int, wantedNumberOfTurns int) (spokenNumber int) {

	cache := make(map[int][]int, 0)

	for actualTurn := 1; actualTurn <= wantedNumberOfTurns; actualTurn++ {
		if actualTurn <= len(numbers) {
			spokenNumber = numbers[actualTurn-1]
			cache = WriteToCache(spokenNumber, actualTurn, cache)

		} else {

			if numberWasNotSpokenBefore(spokenNumber, cache) {
				spokenNumber = 0
			} else {
				turns := cache[spokenNumber]
				spokenNumber = turns[1] - turns[0]
			}
			cache = WriteToCache(spokenNumber, actualTurn, cache)
		}

	}

	return
}

func main() {
	input := []int{0, 13, 1, 8, 6, 15}
	solution1 := GetSpokenNumber(input, 2020)
	fmt.Printf("day 15, part1 %v\n", solution1)
	solution2 := GetSpokenNumber(input, 30000000)
	fmt.Printf("day 15, part2 %v\n", solution2)
}
