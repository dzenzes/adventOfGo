package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/dmies/adventOfGo/filehandler"
)

func CountOneInBinaries(input []string, position int) (count int) {

	for _, line := range input {
		if string(line[position]) == "1" {
			count++
		}
	}
	return
}

func GetMostCommonBit(input []string, position int) string {
	count := CountOneInBinaries(input, position)
	if count > len(input)-count {
		return "1"
	}
	return "0"
}

func CalculateGammaAndEpsilonRate(input []string) (int, int, error) {
	gammaBinary := ""
	epsilonBinary := ""
	for i := 0; i < len(input[0]); i++ {
		mostCommonBit := GetMostCommonBit(input, i)
		gammaBinary += mostCommonBit
		if mostCommonBit == "1" {
			epsilonBinary += "0"
		} else {
			epsilonBinary += "1"
		}
	}

	gamma, err := strconv.ParseInt(gammaBinary, 2, 0)
	if err != nil {
		return 0, 0, err
	}
	epsilon, err2 := strconv.ParseInt(epsilonBinary, 2, 0)
	if err2 != nil {
		return 0, 0, err2
	}
	return int(gamma), int(epsilon), nil
}

func CalculatePowerConsumption(input []string) (int, error) {
	gamma, epsilon, err := CalculateGammaAndEpsilonRate(input)
	if err != nil {
		return 0, err
	}
	return gamma * epsilon, nil
}

func FilterByBit(input []string, bit string, position int) (result []string) {
	for _, line := range input {
		if string(line[position]) == bit {
			result = append(result, line)
		}
	}
	return
}

type filterFunction func(input []string, numberOfOnes int) string

func CalculateRating(input []string, fn filterFunction) (int, error) {
	for i := 0; i < len(input[0]) && len(input) > 1; i++ {
		numberOfOnes := CountOneInBinaries(input, i)
		filterBy := fn(input, numberOfOnes)
		input = FilterByBit(input, filterBy, i)
	}
	if len(input) != 1 {
		return 0, errors.New("couldn't find single binary")
	}
	oxygen, err := strconv.ParseInt(input[0], 2, 0)
	if err != nil {
		return 0, err
	}
	return int(oxygen), nil
}

func CalculateOxygenGeneratorRating(input []string) (int, error) {
	return CalculateRating(input, func(input []string, numberOfOnes int) string {
		filterBy := "0"
		if numberOfOnes >= len(input)-numberOfOnes {
			filterBy = "1"
		}
		return filterBy
	})
}

func CalculateCO2ScrubberRating(input []string) (int, error) {
	return CalculateRating(input, func(input []string, numberOfOnes int) string {
		filterBy := "1"
		if numberOfOnes >= len(input)-numberOfOnes {
			filterBy = "0"
		}
		return filterBy
	})
}

func CalculateLifeSupport(input []string) (int, error) {
	oxygen, err := CalculateOxygenGeneratorRating(input)
	if err != nil {
		return 0, err
	}
	co2, err2 := CalculateCO2ScrubberRating(input)
	if err2 != nil {
		return 0, err2
	}

	return oxygen * co2, nil
}

func main() {
	input, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatalf("Couldn't read input")
	}
	power, _ := CalculatePowerConsumption(input)
	fmt.Printf("day 03, part1 %v\n", power)
	lifeSupport, _ := CalculateLifeSupport(input)
	fmt.Printf("day 03, part2 %v\n", lifeSupport)
}
