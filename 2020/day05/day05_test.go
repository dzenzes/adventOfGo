package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_CalculateRange_only_supports_F_and_B(t *testing.T) {
	tests := []struct {
		comment       string
		letter        string
		errorExpected bool
	}{
		{"valid case", "F", false},
		{"valid case", "B", false},
		{"invalid case", "X", true},
	}

	for _, test := range tests {
		t.Run("CalculateRange() checks letter "+test.comment+" ("+test.letter+")", func(t *testing.T) {
			_, _, err := CalculateRange(test.letter, 1, 2)
			if (err != nil) != test.errorExpected {
				t.Errorf("CalculateRange() error expected: %v but was %v", test.errorExpected, !test.errorExpected)
			}
		})
	}
}

func Test_CalculateRange(t *testing.T) {
	tests := []struct {
		comment   string
		letter    string
		min       int
		max       int
		wantedMin int
		wantedMax int
	}{
		{"get lower half", "F", 0, 127, 0, 63},
		{"get upper half", "B", 0, 63, 32, 63},
		{"get lower half", "F", 32, 63, 32, 47},
		{"get upper half", "B", 32, 47, 40, 47},
		{"get upper half", "B", 40, 47, 44, 47},
		{"get lower half", "F", 44, 47, 44, 45},
		{"get lower half", "F", 44, 45, 44, 44},
	}

	for _, test := range tests {
		t.Run("CalculateRange() "+test.comment, func(t *testing.T) {
			min, max, err := CalculateRange(test.letter, test.min, test.max)
			if err != nil {
				t.Errorf("CalculateRange() throwed an error: %v", err)
			}
			if min != test.wantedMin {
				t.Errorf("CalculateRange() got wrong min. Wanted %v got %v", test.wantedMin, min)
			}
			if max != test.wantedMax {
				t.Errorf("CalculateRange() got wrong max. Wanted %v got %v", test.wantedMax, max)
			}
		})
	}
}
func Test_GetRow(t *testing.T) {
	tests := []struct {
		boardingPass string
		row          int
	}{
		{"BFFFBBFRRR", 70},
		{"FFFBBBFRRR", 14},
		{"BBFFBBFRLL", 102},
	}

	for _, test := range tests {
		t.Run("GetRow for boarding pass "+test.boardingPass+" should point to row "+fmt.Sprint(test.row), func(t *testing.T) {
			input := test.boardingPass
			want := test.row
			got, err := GetRow(input)
			if err != nil {
				t.Errorf("GetRow() throwed an error: %v", err)
			}
			if got != want {
				t.Errorf("GetRow() got %v but wanted %v", got, want)
			}
		})
	}

}

func Test_GetSeat(t *testing.T) {
	tests := []struct {
		boardingPass string
		row          int
	}{
		{"BFFFBBFRRR", 7},
		{"FFFBBBFRRR", 7},
		{"BBFFBBFRLL", 4},
	}

	for _, test := range tests {
		t.Run("GetSeat for boarding pass "+test.boardingPass+" should point to row "+fmt.Sprint(test.row), func(t *testing.T) {
			input := test.boardingPass
			want := test.row
			got, err := GetSeat(input)
			if err != nil {
				t.Errorf("GetSeat() throwed an error: %v", err)
			}
			if got != want {
				t.Errorf("GetSeat() got %v but wanted %v", got, want)
			}
		})
	}

}

func Test_GetSeatId(t *testing.T) {
	tests := []struct {
		boardingPass string
		row          int
	}{
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, test := range tests {
		t.Run("GetSeatId for boarding pass "+test.boardingPass+" should point to row "+fmt.Sprint(test.row), func(t *testing.T) {
			input := test.boardingPass
			want := test.row
			got, err := GetSeatId(input)
			if err != nil {
				t.Errorf("GetSeatId() throwed an error: %v", err)
			}
			if got != want {
				t.Errorf("GetSeatId() got %v but wanted %v", got, want)
			}
		})
	}
}

func Test_GetHighestSeatId(t *testing.T) {
	tests := []struct {
		boardingPasses []string
		max            int
	}{
		{[]string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}, 820},
	}

	for _, test := range tests {
		t.Run("GetHighestSeatId should return "+fmt.Sprint(test.max), func(t *testing.T) {
			input := test.boardingPasses
			want := test.max
			got, err := GetHighestSeatId(input)
			if err != nil {
				t.Errorf("GetHighestSeatId() throwed an error: %v", err)
			}
			if got != want {
				t.Errorf("GetHighestSeatId() got %v but wanted %v", got, want)
			}
		})
	}
}

func Test_GetSortedSeatIds(t *testing.T) {
	tests := []struct {
		boardingPasses []string
		sortedIds      []int
	}{
		{[]string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}, []int{119, 567, 820}},
	}
	for i, test := range tests {
		t.Run("GetSortedSeatIds "+fmt.Sprint(i), func(t *testing.T) {
			input := test.boardingPasses
			want := test.sortedIds
			got, err := GetSortedSeatIds(input)
			if err != nil {
				t.Errorf("GetSortedSeatIds() throwed an error: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetSortedSeatIds() got %v but wanted %v", got, want)
			}
		})
	}
}

func Test_GetMissingIds(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 4}, []int{3}},
		{[]int{4, 5, 6, 8, 10}, []int{7, 9}},
	}

	for i, test := range tests {
		t.Run("GetMissingIds "+fmt.Sprint(i), func(t *testing.T) {
			input := test.input
			want := test.expected
			got, err := GetMissingIds(input)
			if err != nil {
				t.Errorf("GetMissingIds() throwed an error: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetMissingIds() got %v but wanted %v", got, want)
			}
		})
	}
}
