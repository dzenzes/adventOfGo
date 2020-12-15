package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/dmies/adventOfGo/filehandler"
)

// PackageDimensions defines the length, width and height of an package
type PackageDimensions struct {
	length int
	width  int
	height int
}

func (p PackageDimensions) getSurface() int {
	length := p.length * p.width
	width := p.width * p.height
	height := p.height * p.length

	// get smallest side
	sides := []int{length, width, height}
	sort.Ints(sides)
	return 2*length + 2*width + 2*height + sides[0]
}

func (p PackageDimensions) getWrap() int {
	// get smallest side
	sides := []int{p.length, p.width, p.height}
	sort.Ints(sides)
	return (sides[0] + sides[1]) * 2
}

func (p PackageDimensions) getBow() int {
	return p.length * p.width * p.height
}

// Parse parses a String and returns the matching PackageDimensions
func Parse(input string) (PackageDimensions, error) {
	dimensions := strings.Split(input, "x")
	length, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return PackageDimensions{}, err
	}
	width, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return PackageDimensions{}, err
	}
	height, err := strconv.Atoi(dimensions[2])
	if err != nil {
		return PackageDimensions{}, err
	}
	return PackageDimensions{length, width, height}, nil
}

// GetTotalSquareFeetOfWrappingPaper parsses the input and returns the total square feet of wrapping paper for the packages
func GetTotalSquareFeetOfWrappingPaper(input []string) (int, error) {
	result := 0
	for _, line := range input {
		dimensions, err := Parse(line)
		if err != nil {
			return -1, err
		}
		result += dimensions.getSurface()
	}
	return result, nil
}

// GetTotalFeetOfRibbon parsses the input and returns the total feet of ribbon for the packages
func GetTotalFeetOfRibbon(input []string) (int, error) {
	result := 0
	for _, line := range input {
		dimensions, err := Parse(line)
		if err != nil {
			return -1, err
		}
		result += dimensions.getWrap() + dimensions.getBow()
	}
	return result, nil
}

func main() {
	input, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution1, err := GetTotalSquareFeetOfWrappingPaper(input)
	if err != nil {
		log.Fatalf("Couldn't calculate total square feet of wrapping paper: %v", err)
	}

	fmt.Printf("day 02, part1 %v\n", solution1)

	solution2, err := GetTotalFeetOfRibbon(input)
	if err != nil {
		log.Fatalf("Couldn't calculate total feet of ribbon: %v", err)
	}

	fmt.Printf("day 02, part2 %v\n", solution2)

}
