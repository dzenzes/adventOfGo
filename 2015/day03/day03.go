package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/dmies/adventOfGo/filehandler"
)

type Coord struct{ x, y int }

type Santa struct{ coords Coord }

func (santa *Santa) right() {
	santa.coords = Coord{santa.coords.x + 1, santa.coords.y}
}

func (santa *Santa) top() {
	santa.coords = Coord{santa.coords.x, santa.coords.y + 1}
}

func (santa *Santa) left() {
	santa.coords = Coord{santa.coords.x - 1, santa.coords.y}
}

func (santa *Santa) bottom() {
	santa.coords = Coord{santa.coords.x, santa.coords.y - 1}
}

func (santa *Santa) deliver(homes map[Coord]bool) map[Coord]bool {
	homes[santa.coords] = true
	return homes
}

func VisitedHouses(input string) (int, error) {
	homes := make(map[Coord]bool)
	coord := Coord{0, 0}

	homes[coord] = true
	santa := new(Santa)
	for _, direction := range strings.Split(string(input), "") {
		switch direction {
		case ">":
			santa.right()
		case "<":
			santa.left()
		case "^":
			santa.top()
		case "v":
			santa.bottom()
		}
		santa.deliver(homes)
	}

	return len(homes), nil
}

func VisitedHousesWithRobot(input string) (int, error) {

	santas := []*Santa{new(Santa), new(Santa)}

	homes := make(map[Coord]bool)
	coord := Coord{0, 0}

	homes[coord] = true

	for index, direction := range strings.Split(string(input), "") {
		santa := santas[index%2]
		switch direction {
		case ">":
			santa.right()
		case "<":
			santa.left()
		case "^":
			santa.top()
		case "v":
			santa.bottom()
		}
		santa.deliver(homes)
	}

	return len(homes), nil
}

func main() {
	input, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution1, err := VisitedHouses(input[0])
	if err != nil {
		log.Fatalf("Couldn't count visited houses: %v", err)
	}

	fmt.Printf("day 03, part1 %v\n", solution1)

	solution2, err := VisitedHousesWithRobot(input[0])
	if err != nil {
		log.Fatalf("Couldn't count visited houses: %v", err)
	}

	fmt.Printf("day 03, part2 %v\n", solution2)
}
