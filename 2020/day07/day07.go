package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dmies/adventOfGo/filehandler"
)

// BagAndCount is a struct that wraps a bag color and the number of bags of this kind, that can be contained within an other bag
type BagAndCount struct {
	bag   string
	count int
}

// ParseAll parses the given rules and creates a map of color -> possible bags
func ParseAll(input []string) map[string][]BagAndCount {
	result := make(map[string][]BagAndCount, 0)

	for _, line := range input {
		lineSplit := strings.Split(line, " bags contain ")

		var contents []BagAndCount
		if !strings.HasSuffix(line, "contain no other bags.") {
			for _, bag := range strings.Split(lineSplit[1], ", ") {
				bagName := strings.Join(strings.Split(bag, " ")[1:], " ")
				bagName = strings.TrimSuffix(bagName, ".")
				bagName = strings.TrimSuffix(bagName, " bag")
				bagName = strings.TrimSuffix(bagName, " bags")
				bagCount, _ := strconv.Atoi(strings.Split(bag, " ")[0])
				bagStruct := BagAndCount{
					bag:   bagName,
					count: bagCount,
				}
				contents = append(contents, bagStruct)
			}
		}

		result[lineSplit[0]] = contents
	}

	return result

}

// FindBag checks if a bag of the given color (searched) can be found direct/indirect in the rules for the given key
func FindBag(searched string, rules map[string][]BagAndCount, key string) bool {

	for _, bag := range rules[key] {
		if bag.bag == searched {
			return true
		}
		if FindBag(searched, rules, bag.bag) {
			return true
		}
	}
	return false
}

// CountBagsThatContainColor checks the rules and counts how many bag colors can eventually contain at least one bag with the given color?
func CountBagsThatContainColor(color string, rules map[string][]BagAndCount) (int, error) {

	var result int
	for name := range rules {
		if FindBag(color, rules, name) {
			result++
		}
	}

	return result, nil
}

func countContainedBags(color string, rules map[string][]BagAndCount) int {
	var result int
	for _, bag := range rules[color] {
		result += bag.count * countContainedBags(bag.bag, rules)
	}
	return result + 1
}

// GetNumberOfContainedBags counts how many individual bags are required inside the bag with the given color
func GetNumberOfContainedBags(color string, rules map[string][]BagAndCount) int {
	return countContainedBags(color, rules) - 1
}

func main() {
	ruleList, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rules := ParseAll(ruleList)

	solution1, err := CountBagsThatContainColor("shiny gold", rules)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("day 07, part1 %v\n", solution1)

	solution2 := GetNumberOfContainedBags("shiny gold", rules)

	fmt.Printf("day 07, part2 %v\n", solution2)

}
