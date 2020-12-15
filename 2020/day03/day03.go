package main

import (
	"fmt"
	"log"

	"github.com/dmies/adventOfGo/filehandler"
)

// TreeOnMap checks if there is a tree ("#") on the map on x,y
func TreeOnMap(mapData []string, x int, y int) (bool, error) {
	if y >= len(mapData) {
		return false, fmt.Errorf("TreeOnMap: couldn't read value on (%v,%v)", x, y)
	}
	line := []rune(mapData[y])

	// now calculate correct x: the map is repeated on the x axis so we use modulo to get the correct value
	xCoordinate := x % len(line)
	if xCoordinate == len(line) {
		xCoordinate = 0
	}
	// get item at position
	item := line[xCoordinate]
	// check if item is a tree
	return string(item) == "#", nil

}

// CountTreesOnMapForSlope checks how often a tree on a given map is hit for the given slope
func CountTreesOnMapForSlope(mapData []string, right int, down int) (int, error) {
	x := right
	y := down
	numberOfTrees := 0

	for y < len(mapData) {
		tree, err := TreeOnMap(mapData, x, y)
		if err != nil {
			return 0, fmt.Errorf("CountTreesOnMapForSlope: couldn't solve map with given slope (%v,%v)", right, down)
		}
		if tree {
			numberOfTrees++
		}

		y += down
		x += right
	}

	return numberOfTrees, nil
}

func main() {
	mapData, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution1, err := CountTreesOnMapForSlope(mapData, 3, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day 03, part1 %v\n", solution1)

	slopes := []struct {
		right int
		down  int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	result := 1
	for _, slope := range slopes {
		solution2, err := CountTreesOnMapForSlope(mapData, slope.right, slope.down)
		if err != nil {
			log.Fatal(err)
		}
		result = result * solution2
	}

	fmt.Printf("day 03, part2 %v\n", result)

}
