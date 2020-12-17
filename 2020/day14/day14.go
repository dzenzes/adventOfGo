package main

import (
	"fmt"
	"github.com/dmies/adventOfGo/filehandler"
	"log"
	"strconv"
	"strings"
)

type Program struct {
	mask    string
	address int
	value   int
}

func (p *Program) calculate() (result int, err error) {

	valueAsBinary := strconv.FormatInt(int64(p.value), 2)

	valueAsBinary = addZerosToFitSize(valueAsBinary, len(p.mask))
	resultString := ""
	for i, mask := range p.mask {
		if string(mask) == "X" {
			resultString += string(valueAsBinary[i])
		} else {
			resultString += string(mask)
		}
	}
	r, err := strconv.ParseInt(resultString, 2, 64)
	if err != nil {
		return result, err
	}
	return int(r), nil
}

func (p *Program) calculateAddresses() (result []int, err error) {
	addressAsBinary := strconv.FormatInt(int64(p.address), 2)
	mask := p.mask

	newAddressWithFloating := AddBitmasks(addressAsBinary, mask)
	addresses := GetAddressesFromFloating(newAddressWithFloating)
	for _, address := range addresses {
		r, err := strconv.ParseInt(address, 2, 64)
		if err != nil {
			return result, err
		}
		result = append(result, int(r))
	}

	return
}

func addZerosToFitSize(input string, length int) (result string) {
	result = input
	for len(result) < length {
		result = "0" + result
	}
	return
}

func ParseMask(input string) (mask string) {
	splittedInput := strings.Split(input, " = ")
	mask = splittedInput[1]
	return
}

func ParseMemory(input string) (address int, value int, err error) {
	splittedInput := strings.Split(input, " = ")
	value, err = strconv.Atoi(splittedInput[1])
	if err != nil {
		return
	}
	addressAsString := splittedInput[0][4 : len(splittedInput[0])-1]
	address, err = strconv.Atoi(addressAsString)
	return
}

func Parse(input []string) (programs []Program, err error) {
	mask := ""

	for _, line := range input {
		if strings.Index(line, "mask") > -1 {
			mask = ParseMask(line)
		} else {
			address, value, err := ParseMemory(line)
			if err != nil {
				return nil, err
			}
			program := Program{mask, address, value}
			programs = append(programs, program)
		}

	}
	return
}

func AddBitmasks(input string, mask string) (result string) {

	bits := addZerosToFitSize(input, len(mask))

	for i, m := range mask {
		if string(m) == "0" {
			result += string(bits[i])
		} else {
			result += string(m)
		}
	}

	return
}

func GetAddressesFromFloating(mask string) (masks []string) {

	for _, char := range mask {
		newMasks := make([]string, 0)
		if len(masks) == 0 {
			if string(char) == "X" {
				newMasks = []string{"0", "1"}
			} else {
				newMasks = []string{string(char)}
			}
		} else {
			for _, currentMask := range masks {
				if string(char) == "X" {
					newMasks = append(newMasks, currentMask+"0")
					newMasks = append(newMasks, currentMask+"1")
				} else {
					newMasks = append(newMasks, currentMask+string(char))
				}
			}
		}
		masks = newMasks
	}

	return
}

func GetSumOfAllPrograms(input []string) (result int, err error) {
	programs, err := Parse(input)
	if err != nil {
		return result, err
	}
	resultMap := make(map[int]int, 0)

	for _, program := range programs {
		resultMap[program.address], err = program.calculate()
		if err != nil {
			return result, err
		}
	}

	for _, value := range resultMap {
		result += value
	}
	return
}

func GetSumOfAllProgramsV2(input []string) (result int, err error) {
	programs, err := Parse(input)
	if err != nil {
		return result, err
	}
	resultMap := make(map[int]int, 0)

	for _, program := range programs {
		addresses, err := program.calculateAddresses()
		if err != nil {
			return result, err
		}
		for _, address := range addresses {
			resultMap[address] = program.value
		}

	}

	for _, value := range resultMap {
		result += value
	}
	return
}

func main() {
	input, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatal("Couldn't read input")
	}

	solution1, err := GetSumOfAllPrograms(input)
	if err != nil {
		log.Fatalf("GetSumOfAllPrograms returned an error %v", err)
	}

	fmt.Printf("day 14, part1 %v\n", solution1)

	solution2, err := GetSumOfAllProgramsV2(input)
	if err != nil {
		log.Fatalf("GetSumOfAllProgramsV2 returned an error %v", err)
	}

	fmt.Printf("day 14, part2 %v\n", solution2)
}
