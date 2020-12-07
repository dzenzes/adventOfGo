package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/dmies/adventOfGo/filehandler"
)

// GroupWithSize wraps the answer for a group and their size
type GroupWithSize struct {
	answers string
	size    int
}

// GetGroupsFromInput gets a []string and joins all lines until an empty line.
func GetGroupsFromInput(input []string) []GroupWithSize {
	result := make([]GroupWithSize, 0)
	actualLine := ""
	size := 0
	for _, line := range input {
		if line == "" {
			result = append(result, GroupWithSize{actualLine, size})
			actualLine = ""
			size = 0
		} else {
			actualLine = actualLine + line
			size++
		}
	}
	result = append(result, GroupWithSize{actualLine, size})
	return result
}

// Unique removes all duplicate letters in a string
func Unique(input string) string {
	keys := make(map[rune]bool)
	var result []rune
	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return string(result)
}

// CountCharsPerLine counts the length of each string in the input and returns it
func CountCharsPerLine(input []string) []int {
	result := make([]int, 0)
	for _, line := range input {
		result = append(result, len(line))
	}
	return result
}

// GetSumOfAnswers checks the sum of answers given in the input
func GetSumOfAnswers(input []string) int {
	result := 0
	groups := GetGroupsFromInput(input)
	filteredGroups := make([]string, 0)
	for _, line := range groups {
		filteredLine := Unique(line.answers)
		filteredGroups = append(filteredGroups, filteredLine)
	}

	countPerGroup := CountCharsPerLine(filteredGroups)

	for _, count := range countPerGroup {
		result += count
	}
	return result

}

// CountCommonAnswers counts how many "answers" where given by all group members
func CountCommonAnswers(input GroupWithSize) int {
	result := 0
	unique := Unique(input.answers)
	for _, letter := range unique {
		count := strings.Count(input.answers, string(letter))
		if count == input.size {
			result++
		}
	}
	return result
}

// GetSumOfCommonAnswers checks the sum of common answers given in the input
func GetSumOfCommonAnswers(input []string) int {
	result := 0
	groups := GetGroupsFromInput(input)
	for _, line := range groups {
		result += CountCommonAnswers(line)

	}

	return result

}

func main() {
	votes, err := filehandler.ImportSringList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution1 := GetSumOfAnswers(votes)

	fmt.Printf("day 06, part1 %v\n", solution1)

	solution2 := GetSumOfCommonAnswers(votes)

	fmt.Printf("day 06, part2 %v\n", solution2)

}
