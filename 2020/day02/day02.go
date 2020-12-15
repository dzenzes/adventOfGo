package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dmies/adventOfGo/filehandler"
)

// Parse takes a line from input and parses its contents
func Parse(line string) (int, int, string, string, error) {
	configAndPassword := strings.Split(line, ":")
	config := strings.Trim(configAndPassword[0], " ")
	password := strings.Trim(configAndPassword[1], " ")
	minMaxAndLetter := strings.Split(config, " ")
	minMax := strings.Split(minMaxAndLetter[0], "-")
	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		return 0, 0, "", "", fmt.Errorf("Parse: couldn't parse minimum %s", minMax[0])
	}
	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		return 0, 0, "", "", fmt.Errorf("Parse: couldn't parse maximum %s", minMax[1])
	}
	letter := minMaxAndLetter[1]
	return min, max, letter, password, nil
}

// ValidatePassword checks if the given passwords matches the given policy
func ValidatePassword(min int, max int, letter string, password string) bool {

	numberOfOccurrences := strings.Count(password, letter)
	return numberOfOccurrences >= min && numberOfOccurrences <= max
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

// ValidatePasswordForNewShop checks if the given passwords matches the given policy (for the new shop)
func ValidatePasswordForNewShop(first int, second int, letter string, password string) bool {
	firstIndex := first - 1
	secondIndex := second - 1

	letterAtFirstPosition := substr(password, firstIndex, 1)
	letterAtSecondPosition := substr(password, secondIndex, 1)
	if letter == letterAtFirstPosition || letter == letterAtSecondPosition {
		if letterAtFirstPosition != letterAtSecondPosition {
			return true
		}
	}

	return false

}

// Part1 is yet to be implemented
func Part1(passwordsWithPolicy []string) (int, error) {
	var count = 0
	for _, line := range passwordsWithPolicy {
		min, max, letter, password, err := Parse(line)
		if err != nil {
			return 0, err
		}
		isValid := ValidatePassword(min, max, letter, password)
		if isValid {
			count++
		}
	}
	return count, nil
}

// Part2 is yet to be implemented
func Part2(passwordsWithPolicy []string) (int, error) {
	var count = 0
	for _, line := range passwordsWithPolicy {
		min, max, letter, password, err := Parse(line)
		if err != nil {
			return 0, err
		}
		isValid := ValidatePasswordForNewShop(min, max, letter, password)
		if isValid {
			count++
		}
	}
	return count, nil
}

func main() {
	passwordsWithPolicy, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution1, err := Part1(passwordsWithPolicy)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day 02, part1 %v\n", solution1)
	solution2, err := Part2(passwordsWithPolicy)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day 02, part2 %v\n", solution2)

}
