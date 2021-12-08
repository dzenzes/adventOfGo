package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dmies/adventOfGo/filehandler"
)

func FinalScore(input []string) (int, error) {
	numbers, boards, err := parseInput(input)
	if err != nil {
		return -1, err
	}

	for _, number := range numbers {
		for _, board := range boards {
			board.mark(number)

			if board.solved() {
				return board.sum() * number, nil
			}
		}
	}

	return -1, errors.New("no board solved")
}

func ScoreOfLastWinningBoard(input []string) (int, error) {
	numbers, boards, err := parseInput(input)
	if err != nil {
		return -1, err
	}
	solvedBoards := 0
	for _, number := range numbers {
		for _, board := range boards {

			if !board.solved() && board.mark(number) {
				if board.solved() {
					solvedBoards++
					if len(boards) == solvedBoards {
						return board.sum() * number, nil
					}
				}
			}
		}
	}

	return -1, errors.New("no last board found")
}

func ParseNumbers(input string) ([]int, error) {
	numbers := []int{}
	if strings.TrimSpace(input) == "" {
		return numbers, nil
	}

	for _, numberString := range strings.Split(input, ",") {
		num, err := strconv.Atoi(numberString)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

type Cell struct {
	value  int
	marked bool
}

type Board struct {
	cells [][]Cell
}

func (board *Board) mark(num int) bool {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if board.cells[r][c].value == num {
				board.cells[r][c].marked = true
				return true
			}
		}
	}
	return false
}

func (board *Board) solved() bool {
	checkColumn := func(column int) bool {
		for row := 0; row < 5; row++ {
			if !board.cells[row][column].marked {
				return false
			}
		}
		return true
	}

	checkRow := func(row int) bool {
		for column := 0; column < 5; column++ {
			if !board.cells[row][column].marked {
				return false
			}
		}
		return true
	}
	for row := 0; row < 5; row++ {
		if checkRow(row) {
			return true
		}
	}
	for column := 0; column < 5; column++ {
		if checkColumn(column) {
			return true
		}
	}
	return false
}

func (board *Board) sum() (sum int) {
	for _, row := range board.cells {
		for _, column := range row {
			if !column.marked {
				sum += column.value
			}
		}
	}
	return
}

func NewCell(input string) (Cell, error) {
	value, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return Cell{}, err
	}
	return Cell{value, false}, nil
}

func NewBoard(input []string) (board Board, err error) {
	if len(input) == 0 {
		err = errors.New("Board needs valid input")
		return
	}

	for r := 0; r < 5; r++ {
		board.cells = append(board.cells, make([]Cell, 5))
	}

	for y, line := range input {
		for x, value := range strings.Fields(line) {
			board.cells[y][x], err = NewCell(value)
			if err != nil {
				return board, err
			}
		}
	}
	return
}

func ReadBoards(input []string) ([]Board, error) {

	boards := []Board{}
	nextBoard := []string{}

	for _, line := range input {
		if strings.TrimSpace(line) == "" {
			b, boardError := NewBoard(nextBoard)
			if boardError != nil {
				return nil, boardError
			}

			boards = append(boards, b)
			nextBoard = []string{}
		} else {
			nextBoard = append(nextBoard, line)
		}
	}
	b, boardError := NewBoard(nextBoard)
	if boardError != nil {
		return nil, boardError
	}

	boards = append(boards, b)
	return boards, nil
}

func parseInput(input []string) (numbers []int, boards []Board, err error) {
	numbers, err = ParseNumbers(input[0])
	if err != nil {
		return
	}

	boards, err = ReadBoards(input[2:])
	if err != nil {
		return
	}
	return
}

func main() {
	input, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatalf("Couldn't read input")
	}
	score, _ := FinalScore(input)
	fmt.Printf("day 04, part1 %v\n", score)
	lastScore, _ := ScoreOfLastWinningBoard(input)
	fmt.Printf("day 04, part2 %v\n", lastScore)

}
