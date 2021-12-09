package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dmies/adventOfGo/filehandler"
)

func HowManyLanternfishAfterDays(input string, days int) (result int) {
	fishes := map[int]int{}
	for _, ageString := range strings.Split(input, ",") {

		age, err := strconv.Atoi(ageString)
		if err != nil {
			panic("couldn't parse input")
		}

		fishes[age]++

	}

	for day := 0; day < days; day++ {
		nextRound := map[int]int{}
		for age := 0; age <= 8; age++ {
			nextRound[age] += fishes[(age+1)%9]
		}
		nextRound[6] += fishes[0]
		fishes = nextRound
	}

	for _, fishes := range fishes {
		result += fishes
	}

	return result

}
func main() {
	input, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatalf("Couldn't read input")
	}
	numberOfFishes := HowManyLanternfishAfterDays(input[0], 80)
	fmt.Printf("day 06, part1 %v\n", numberOfFishes)

	numberOfFishes = HowManyLanternfishAfterDays(input[0], 256)
	fmt.Printf("day 06, part2 %v\n", numberOfFishes)
}
