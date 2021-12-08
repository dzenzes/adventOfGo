package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var emptyBoard = Board{
	cells: [][]Cell{
		{
			{22, false},
			{13, false},
			{17, false},
			{11, false},
			{0, false},
		},
		{
			{8, false},
			{2, false},
			{23, false},
			{4, false},
			{24, false},
		},
		{
			{21, false},
			{9, false},
			{14, false},
			{16, false},
			{7, false},
		},
		{
			{6, false},
			{10, false},
			{3, false},
			{18, false},
			{5, false},
		},
		{
			{1, false},
			{12, false},
			{20, false},
			{15, false},
			{19, false},
		},
	},
}

func Test_FinalScore(t *testing.T) {

	input := []string{
		"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
		"",
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
		"",
		" 3 15  0  2 22",
		" 9 18 13 17  5",
		"19  8  7 25 23",
		"20 11 10 24  4",
		"14 21 16 12  6",
		"",
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	}
	want := 4512
	got, _ := FinalScore(input)
	assert.Equal(t, want, got)
}
func Test_ScoreOfLastWinningBoard(t *testing.T) {

	input := []string{
		"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
		"",
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
		"",
		" 3 15  0  2 22",
		" 9 18 13 17  5",
		"19  8  7 25 23",
		"20 11 10 24  4",
		"14 21 16 12  6",
		"",
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	}
	want := 1924
	got, _ := ScoreOfLastWinningBoard(input)
	assert.Equal(t, want, got)
}

func Test_ParseNumbers(t *testing.T) {

	tests := []struct {
		input string
		want  []int
	}{
		{
			"",
			[]int{},
		},
		{
			"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
			[]int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
		},
	}

	for _, test := range tests {
		t.Run("ParseNumbers "+fmt.Sprint(test.input), func(t *testing.T) {
			got, err := ParseNumbers(test.input)
			if err != nil {
				t.Errorf("Error parsing numbers")
			}
			assert.Equal(t, test.want, got)
		})
	}

}

func Test_NewBoard(t *testing.T) {
	t.Run("NewBoard needs some valid input to work", func(t *testing.T) {
		input := []string{}
		_, err := NewBoard(input)
		assert.Error(t, err, "Board needs valid input")
	})

	tests := []struct {
		input []string
		want  Board
	}{
		{
			[]string{
				"22 13 17 11  0",
				"8  2 23  4 24",
				"21  9 14 16  7",
				"6 10  3 18  5",
				"1 12 20 15 19",
			},
			emptyBoard,
		},
	}

	for _, test := range tests {
		t.Run("NewBoard "+fmt.Sprint(test.input), func(t *testing.T) {
			got, err := NewBoard(test.input)
			if err != nil {
				t.Errorf("Error parsing board")
			}
			assert.Equal(t, test.want, got)
		})

	}
}

func Test_MarkNumber(t *testing.T) {
	input := emptyBoard

	want := Board{
		cells: [][]Cell{
			{
				{22, false},
				{13, false},
				{17, false},
				{11, false},
				{0, false},
			},
			{
				{8, false},
				{2, false},
				{23, true},
				{4, false},
				{24, false},
			},
			{
				{21, false},
				{9, false},
				{14, false},
				{16, false},
				{7, false},
			},
			{
				{6, false},
				{10, false},
				{3, false},
				{18, false},
				{5, false},
			},
			{
				{1, false},
				{12, false},
				{20, false},
				{15, false},
				{19, false},
			},
		},
	}

	got := input.mark(23)
	assert.True(t, got)
	assert.Equal(t, want, input)
}
func Test_BoardSolved(t *testing.T) {
	boardInput := []string{
		"22 13 17 11  0",
		"8  2 23  4 24",
		"21  9 14 16  7",
		"6 10  3 18  5",
		"1 12 20 15 19",
	}

	tests := []struct {
		markedNumbers []int
		solved        bool
	}{
		{[]int{6, 10, 3, 18, 50}, false},
		{[]int{6, 10, 3, 18, 5}, true},
		{[]int{1, 12, 20, 15, 19}, true},
		{[]int{17, 23, 14, 3, 20}, true},
	}
	for _, test := range tests {

		t.Run("board.solved ("+fmt.Sprint(test.markedNumbers)+")", func(t *testing.T) {
			board, _ := NewBoard(boardInput)
			for _, number := range test.markedNumbers {
				board.mark(number)
			}
			got := board.solved()
			assert.Equal(t, test.solved, got)
		})
	}
}

func Test_BoardSumOfUnmarked(t *testing.T) {
	boardInput := []string{
		"22 13 17 11  0",
		"8  2 23  4 24",
		"21  9 14 16  7",
		"6 10  3 18  5",
		"1 12 20 15 19",
	}

	tests := []struct {
		markedNumbers []int
		want          int
	}{
		{[]int{}, 300},
		{[]int{22}, 278},
		{[]int{6, 10, 3, 18, 5}, 258},
	}
	for _, test := range tests {

		t.Run("board.sum ("+fmt.Sprint(test.markedNumbers)+")", func(t *testing.T) {
			board, _ := NewBoard(boardInput)
			for _, number := range test.markedNumbers {
				board.mark(number)
			}
			got := board.sum()
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_ReadBoards(t *testing.T) {
	input := []string{
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
		"",
		" 3 15  0  2 22",
		" 9 18 13 17  5",
		"19  8  7 25 23",
		"20 11 10 24  4",
		"14 21 16 12  6",
		"",
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	}
	board1, _ := NewBoard([]string{
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
	})
	board2, _ := NewBoard([]string{
		" 3 15  0  2 22",
		" 9 18 13 17  5",
		"19  8  7 25 23",
		"20 11 10 24  4",
		"14 21 16 12  6",
	})
	board3, _ := NewBoard([]string{
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	})
	want := []Board{
		board1, board2, board3,
	}

	got, err := ReadBoards(input)
	if err != nil {
		t.Errorf("ReadBoards couldn't parse given boards")
	}
	assert.Equal(t, want, got)
}
