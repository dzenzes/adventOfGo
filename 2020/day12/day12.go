package main

import (
	"errors"
	"fmt"
	"github.com/dmies/adventOfGo/filehandler"
	"log"
	"math"
	"strconv"
)

type Instruction struct {
	action string
	value  int
}

type Ship struct {
	north     int
	east      int
	direction string
}

func (s *Ship) getManhattanDistance() int {
	result := math.Abs(float64(s.north)) + math.Abs(float64(s.east))
	return int(result)
}

type Waypoint struct {
	north int
	east  int
}

const (
	IllegalInstructionError = "instruction is not supported"
	UnsupportedValueError   = "value of instruction is not supported"
)

func Parse(input string) (Instruction, error) {
	action := input[0:1]
	valueAsString := input[1:len(input)]
	value, err := strconv.Atoi(valueAsString)
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{action, value}, nil
}

func ParseList(input []string) ([]Instruction, error) {
	result := make([]Instruction, 0)
	for _, row := range input {
		instruction, err := Parse(row)
		if err != nil {
			return []Instruction{}, err
		}
		result = append(result, instruction)
	}
	return result, nil
}

func directionToNumber(direction string) (result int, err error) {
	switch direction {
	case "N":
		return 0, nil
	case "E":
		return 1, nil
	case "S":
		return 2, nil
	case "W":
		return 3, nil
	default:
		return -1, errors.New("cannot parse direction")
	}
}

func numberToDirection(input int) (result string, err error) {
	normalizedNumber := input % 4
	switch normalizedNumber {
	case 0:
		return "N", nil
	case 1:
		return "E", nil
	case 2:
		return "S", nil
	case 3:
		return "W", nil
	default:
		return "", errors.New("cannot transform number (" + fmt.Sprint(normalizedNumber) + ") to direction")
	}
}

func CalculateDirection(oldDirection string, instruction Instruction) (newDirection string, err error) {
	if instruction.value%90 != 0 {
		return "", errors.New(UnsupportedValueError)
	}
	movement := instruction.value / 90
	directionAsNumber, err := directionToNumber(oldDirection)
	if err != nil {
		return "", err
	}
	var newDirectionAsNumber int
	if instruction.action == "R" {
		newDirectionAsNumber = directionAsNumber + movement
	} else if instruction.action == "L" {
		newDirectionAsNumber = directionAsNumber - movement
		if newDirectionAsNumber < 1 {
			newDirectionAsNumber = 4 + newDirectionAsNumber
		}
	}

	newDirection, err = numberToDirection(newDirectionAsNumber)
	if err != nil {
		return "", err
	}
	return newDirection, nil
}

func ApplyInstruction(ship Ship, instruction Instruction) (Ship, error) {

	switch instruction.action {
	case "N":
		return Ship{ship.north + instruction.value, ship.east, ship.direction}, nil
	case "S":
		return Ship{ship.north - instruction.value, ship.east, ship.direction}, nil
	case "E":
		return Ship{ship.north, ship.east + instruction.value, ship.direction}, nil
	case "W":
		return Ship{ship.north, ship.east - instruction.value, ship.direction}, nil
	case "L":
		newDirection, err := CalculateDirection(ship.direction, instruction)
		if err != nil {
			return Ship{}, err
		}
		return Ship{ship.north, ship.east, newDirection}, nil
	case "R":
		newDirection, err := CalculateDirection(ship.direction, instruction)
		if err != nil {
			return Ship{}, err
		}
		return Ship{ship.north, ship.east, newDirection}, nil
	case "F":
		switch ship.direction {
		case "N":
			return Ship{ship.north + instruction.value, ship.east, ship.direction}, nil
		case "E":
			return Ship{ship.north, ship.east + instruction.value, ship.direction}, nil
		case "S":
			return Ship{ship.north - instruction.value, ship.east, ship.direction}, nil
		case "W":
			return Ship{ship.north, ship.east - instruction.value, ship.direction}, nil
		default:
			return Ship{}, errors.New(IllegalInstructionError)
		}
	}

	return Ship{}, errors.New(IllegalInstructionError)

}

func ApplyInstructionWithWaypoint(ship Ship, waypoint Waypoint, instruction Instruction) (newShip Ship, newWaypoint Waypoint, err error) {

	switch instruction.action {
	case "F":
		value := instruction.value
		newShip = Ship{ship.north + value*waypoint.north, ship.east + value*waypoint.east, ship.direction}
		return newShip, waypoint, nil
	case "N":
		newWaypoint = Waypoint{waypoint.north + instruction.value, waypoint.east}
		return ship, newWaypoint, nil
	case "E":
		newWaypoint = Waypoint{waypoint.north, waypoint.east + instruction.value}
		return ship, newWaypoint, nil
	case "S":
		newWaypoint = Waypoint{waypoint.north - instruction.value, waypoint.east}
		return ship, newWaypoint, nil
	case "W":
		newWaypoint = Waypoint{waypoint.north, waypoint.east - instruction.value}
		return ship, newWaypoint, nil
	case "R":
		normalizedDegrees := instruction.value / 90
		north := waypoint.north
		east := waypoint.east
		if normalizedDegrees == 1 {
			north = waypoint.east * -1
			east = waypoint.north
		} else if normalizedDegrees == 2 {
			north = waypoint.north * -1
			east = waypoint.east * -1
		} else if normalizedDegrees == 3 {
			north = waypoint.east
			east = waypoint.north * -1
		} else if normalizedDegrees == 4 {
			north = waypoint.north
			east = waypoint.east
		}
		return ship, Waypoint{north, east}, nil
	case "L":
		normalizedDegrees := instruction.value / 90
		north := waypoint.north
		east := waypoint.east
		if normalizedDegrees == 1 {
			north = waypoint.east
			east = waypoint.north * -1
		} else if normalizedDegrees == 2 {
			north = waypoint.north * -1
			east = waypoint.east * -1
		} else if normalizedDegrees == 3 {
			north = waypoint.east * -1
			east = waypoint.north
		} else if normalizedDegrees == 4 {
			north = waypoint.north
			east = waypoint.east
		}
		return ship, Waypoint{north, east}, nil
	}

	return Ship{}, Waypoint{}, errors.New(IllegalInstructionError)
}

func Process(input []string) (ship Ship, err error) {
	instructions, err := ParseList(input)
	if err != nil {
		return Ship{}, err
	}

	ship = Ship{0, 0, "E"}

	for _, instruction := range instructions {
		ship, err = ApplyInstruction(ship, instruction)
		if err != nil {
			return Ship{}, err
		}
	}

	return ship, nil
}

func ProcessWithWaypoint(input []string) (ship Ship, err error) {
	instructions, err := ParseList(input)
	if err != nil {
		return Ship{}, err
	}

	ship = Ship{0, 0, "E"}
	waypoint := Waypoint{1, 10}

	for _, instruction := range instructions {
		ship, waypoint, err = ApplyInstructionWithWaypoint(ship, waypoint, instruction)
		if err != nil {
			return Ship{}, err
		}
	}

	return ship, nil
}

func main() {
	input, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatal("Couldn't read input")
	}

	ship, err := Process(input)
	if err != nil {
		log.Fatalf("part1 returned an error %v", err)
	}
	solution1 := ship.getManhattanDistance()

	fmt.Printf("day 12, part1 %v %v\n", solution1, ship)

	ship2, err := ProcessWithWaypoint(input)
	if err != nil {
		log.Fatalf("part2 returned an error %v", err)
	}
	solution2 := ship2.getManhattanDistance()

	fmt.Printf("day 12, part2 %v %v\n", solution2, ship)
}
