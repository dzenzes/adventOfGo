package main

import "strconv"

func GetMostCommonBit(input []string, position int) string {
	count := 0
	for _, line := range input {
		if string(line[position]) == "1" {
			count++
		}
	}
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
