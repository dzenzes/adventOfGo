package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/dmies/adventOfGo/filehandler"
)

type Point struct {
	x int
	y int
}

func (p *Point) isNoDiagonalLine(q Point) bool {
	return p.x == q.x || p.y == q.y
}

func ParseLine(input string) (p Point, q Point, err error) {
	r := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

	parseInt := func(input string) int {
		res, err := strconv.Atoi(input)
		if err != nil {
			panic("couldn't parse int")

		}
		return res
	}

	matches := r.FindStringSubmatch(input)
	p = Point{parseInt(matches[1]), parseInt(matches[2])}
	q = Point{parseInt(matches[3]), parseInt(matches[4])}
	return
}

func PointsOnLine(p Point, q Point, countDiagonal bool) (line []Point) {
	if !countDiagonal && !p.isNoDiagonalLine(q) {
		return []Point{}
	}

	nextPoint := p

	for i := 0; nextPoint.x != q.x || nextPoint.y != q.y; i++ {
		line = append(line, nextPoint)

		if nextPoint.x < q.x {
			nextPoint.x++
		} else if nextPoint.x > q.x {
			nextPoint.x--
		}

		if nextPoint.y < q.y {
			nextPoint.y++
		} else if nextPoint.y > q.y {
			nextPoint.y--
		}
	}

	line = append(line, nextPoint)
	return
}

func CountOverlappingLines(input []string) int {
	grid := map[Point]uint{}
	for _, line := range input {
		p, q, err := ParseLine(line)
		if err != nil {
			panic("couldn't parse input")
		}
		for _, point := range PointsOnLine(p, q, false) {
			grid[point]++
		}
	}
	res := 0
	for _, count := range grid {
		if count >= 2 {
			res++
		}
	}
	return res
}

func CountOverlappingLinesIncludingDiagonals(input []string) int {

	grid := map[Point]uint{}
	for _, line := range input {
		p, q, err := ParseLine(line)
		if err != nil {
			panic("couldn't parse input")
		}
		for _, point := range PointsOnLine(p, q, true) {
			grid[point]++
		}
	}
	res := 0
	for _, count := range grid {
		if count >= 2 {
			res++
		}
	}
	return res
}

func main() {
	input, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatalf("Couldn't read input")
	}
	numberOfOverlappingPoints := CountOverlappingLines(input)
	fmt.Printf("day 05, part1 %v\n", numberOfOverlappingPoints)

	numberOfOverlappingPoints2 := CountOverlappingLinesIncludingDiagonals(input)
	fmt.Printf("day 05, part2 %v\n", numberOfOverlappingPoints2)
}
