package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dmies/adventOfGo/filehandler"
)

type Position struct {
	x int
	y int
}

type submarine struct {
	position Position
	aim      bool
	aimValue int
}

func NewSubmarine() submarine {
	return submarine{position: Position{x: 0, y: 0}, aim: false, aimValue: 0}
}

func NewSubmarineWithAim() submarine {
	return submarine{position: Position{x: 0, y: 0}, aim: true, aimValue: 0}

}
func (submarine *submarine) forward(distance int) {
	submarine.position.x += distance
	if submarine.aim {
		submarine.position.y += submarine.aimValue * distance
	}
}
func (submarine *submarine) down(distance int) {
	if submarine.aim {
		submarine.aimValue += distance
	} else {

		submarine.position.y += distance
	}
}
func (submarine *submarine) up(distance int) {
	if submarine.aim {
		submarine.aimValue -= distance
	} else {

		submarine.position.y -= distance
	}
}

func ParseInput(line string) (command string, distance int, err error) {
	tokens := strings.Split(line, " ")
	command = tokens[0]
	distance, err = strconv.Atoi(tokens[1])
	return
}

func (submarine *submarine) calculatePosition(input []string) (position Position, err error) {

	for _, line := range input {

		command, distance, err2 := ParseInput(line)

		if err2 != nil {
			err = err2
			return
		}
		switch command {
		case "forward":
			submarine.forward(distance)
		case "down":
			submarine.down(distance)
		case "up":
			submarine.up(distance)
		}
	}

	return submarine.position, nil
}

func main() {
	input, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatalf("Couldn't read input")
	}
	submarine := NewSubmarine()
	position, _ := submarine.calculatePosition(input)
	fmt.Printf("day 02, part1 %v\n", position.x*position.y)
	submarine2 := NewSubmarineWithAim()
	position2, _ := submarine2.calculatePosition(input)
	fmt.Printf("day 02, part2 %v\n", position2.x*position2.y)
}
