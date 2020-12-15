package main

import (
	"errors"
	"fmt"
	"github.com/dmies/adventOfGo/filehandler"
	"log"
	"reflect"
)

type Seat struct {
	isSeat     bool
	isOccupied bool
}

func (s Seat) String() string {
	if s.isSeat && s.isOccupied {
		return "#"
	} else if s.isSeat && !s.isOccupied {
		return "L"
	} else {
		return "."
	}
}

type CountFunction func(int, int, [][]Seat) int

func ParsePosition(input string) (Seat, error) {
	if len(input) == 1 {
		switch input {
		case "L":
			return Seat{true, false}, nil
		case "#":
			return Seat{true, true}, nil
		case ".":
			return Seat{false, false}, nil
		}
	}
	return Seat{}, errors.New("couldn't parse input")
}

func ParseMap(input []string) ([][]Seat, error) {
	result := make([][]Seat, 0)
	for _, row := range input {
		parsedRow := make([]Seat, 0)
		for _, letter := range row {
			seat, err := ParsePosition(string(letter))
			if err != nil {
				return nil, err
			}
			parsedRow = append(parsedRow, seat)
		}
		result = append(result, parsedRow)
	}
	return result, nil
}

func CountOccupiedNeighbors(x int, y int, seats [][]Seat) int {
	var iMin, jMin, iMax, jMax int
	iMin = y - 1
	if iMin < 0 {
		iMin = 0
	}
	jMin = x - 1
	if jMin < 0 {
		jMin = 0
	}
	iMax = y + 1
	if iMax == len(seats) {
		iMax--
	}
	jMax = x + 1
	if jMax == len(seats[0]) {
		jMax--
	}
	var occupiedCounter int
	for i := iMin; i <= iMax; i++ {
		for j := jMin; j <= jMax; j++ {
			if !(i == y && j == x) {
				if seats[i][j].isSeat {
					if seats[i][j].isOccupied {
						occupiedCounter++

					}
				}
			}
		}
	}
	return occupiedCounter
}

func CountVisibleOccupiedSeats(x int, y int, seats [][]Seat) int {
	var occupiedCounter int
	maxY := len(seats)
	maxX := len(seats[maxY-1])

	// west
	for i := 1; x-i >= 0; i++ {
		if seats[y][x-i].isSeat {
			if seats[y][x-i].isOccupied {
				occupiedCounter++
			}
			break
		}
	}
	// north west
	for i := 1; y-i >= 0 && x-i >= 0; i++ {
		if seats[y-i][x-i].isSeat {
			if seats[y-i][x-i].isOccupied {
				occupiedCounter++
			}
			break
		}
	}
	// north
	for i := 1; y-i >= 0; i++ {
		if seats[y-i][x].isSeat {
			if seats[y-i][x].isOccupied {
				occupiedCounter++
			}
			break
		}
	}

	// north east
	for i := 1; y-i >= 0 && x+i < maxX; i++ {
		if seats[y-i][x+i].isSeat {
			if seats[y-i][x+i].isOccupied {
				occupiedCounter++
			}
			break
		}
	}

	//  east
	for i := 1; x+i < maxX; i++ {
		if seats[y][x+i].isSeat {
			if seats[y][x+i].isOccupied {
				occupiedCounter++
			}
			break
		}
	}

	//  south east
	for i := 1; (y+i < maxY) && (x+i < maxX); i++ {
		if seats[y+i][x+i].isSeat {
			if seats[y+i][x+i].isOccupied {
				occupiedCounter++
			}
			break
		}

	}

	//  south
	for i := 1; y+i < maxY; i++ {
		if seats[y+i][x].isSeat {
			if seats[y+i][x].isOccupied {
				occupiedCounter++
			}
			break
		}

	}

	//  south west
	for i := 1; y+i < maxY && x-i >= 0; i++ {
		if seats[y+i][x-i].isSeat {
			if seats[y+i][x-i].isOccupied {
				occupiedCounter++
			}
			break
		}

	}

	return occupiedCounter
}

func UpdateSeat(x int, y int, seats [][]Seat, maxNeighbors int, countFunction CountFunction) Seat {
	currentSeat := seats[y][x]
	numberOfNeighbors := countFunction(x, y, seats)
	// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
	if currentSeat.isSeat {
		if !currentSeat.isOccupied {
			if numberOfNeighbors == 0 {
				return Seat{true, true}
			}
		} else {
			// If a seat is occupied (#) and maxNeighbors or more seats adjacent to it are also occupied, the seat becomes empty.
			if numberOfNeighbors >= maxNeighbors {
				return Seat{true, false}
			}
		}

	}

	return currentSeat
}

func ProcessRound(seats [][]Seat, maxNeighbors int, countFunction CountFunction) [][]Seat {
	result := make([][]Seat, 0)

	for y, row := range seats {
		resultRow := make([]Seat, 0)
		for x, _ := range row {
			updatedSeat := UpdateSeat(x, y, seats, maxNeighbors, countFunction)
			resultRow = append(resultRow, updatedSeat)
		}
		result = append(result, resultRow)

	}
	return result
}

func CountOccupiedSeats(seats [][]Seat) int {
	result := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat.isSeat && seat.isOccupied {
				result++
			}
		}
	}
	return result
}

func ProcessUntilNoChange(input []string, maxNeighbors int, countFunction CountFunction) (int, error) {
	seats, err := ParseMap(input)
	if err != nil {
		return -1, err
	}
	seatsDidChange := true
	for seatsDidChange {
		updatedSeats := ProcessRound(seats, maxNeighbors, countFunction)
		seatsDidChange = !reflect.DeepEqual(updatedSeats, seats)
		seats = updatedSeats
	}
	result := CountOccupiedSeats(seats)
	return result, nil

}

func main() {
	input, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatal("Couldn't read input")
	}

	solution1, err := ProcessUntilNoChange(input, 4, CountOccupiedNeighbors)
	if err != nil {
		log.Fatalf("ProcessUntilNoChange returned an error %v", err)
	}
	fmt.Printf("day 11, part1 %v\n", solution1)

	solution2, err := ProcessUntilNoChange(input, 5, CountVisibleOccupiedSeats)
	if err != nil {
		log.Fatalf("ProcessUntilNoChange returned an error %v", err)
	}
	fmt.Printf("day 11, part2 %v\n", solution2)

}
