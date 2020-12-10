package main

import (
	"fmt"
	"reflect"
	"testing"
)

func assertErrorIsNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
}

func assertIntListIsEqual(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertEqualInts(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func Test_PrepareInput(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1}, []int{0, 1, 4}},
		{[]int{1, 2, 3}, []int{0, 1, 2, 3, 6}},
		{[]int{4, 2, 1}, []int{0, 1, 2, 4, 7}},
	}

	for _, test := range tests {
		t.Run("PrepareInput("+fmt.Sprint(test.input)+")", func(t *testing.T) {
			got := PrepareInput(test.input)
			assertIntListIsEqual(t, got, test.want)
		})
	}
}

func Test_GetNextAdapter(t *testing.T) {
	adapters := PrepareInput([]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4})

	tests := []struct {
		input int
		want  int
	}{
		{0, 1},
		{1, 4},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 10},
		{10, 11},
		{11, 12},
		{12, 15},
		{15, 16},
		{16, 19},
	}

	for _, test := range tests {
		t.Run("GetNextAdapter("+fmt.Sprint(test.input)+")", func(t *testing.T) {
			got, err := GetNextAdapter(test.input, adapters)
			assertErrorIsNil(t, err)

			assertEqualInts(t, got, test.want)
		})
	}

	t.Run("GetNextAdapter returns an error if no adapter matches", func(t *testing.T) {
		_, err := GetNextAdapter(adapters[len(adapters)-1], adapters)
		if err == nil {
			t.Errorf("GetNextAdapter() should return an error if there is no bigger adapter")
		}
	})

	t.Run("GetNextAdapter returns an error if no adapter can be found (difference of jolts > 3)", func(t *testing.T) {
		brokenAdapters := []int{1, 5}
		_, err := GetNextAdapter(1, brokenAdapters)
		if err == nil {
			t.Errorf("GetNextAdapter() should return an error if there is no bigger adapter")
		}
	})
}

func Test_GetDifferences(t *testing.T) {

	tests := []struct {
		input                []int
		wantedDifferencesOf1 int
		wantedDifferencesOf3 int
	}{
		{[]int{1}, 1, 1},
		{[]int{1, 2}, 2, 1},
		{[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}, 7, 5},
		{[]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}, 22, 10},
	}

	for _, test := range tests {
		t.Run("GetDifferences("+fmt.Sprint(test.input)+")", func(t *testing.T) {
			input := PrepareInput(test.input)
			got1, got3, err := GetDifferences(input)
			assertErrorIsNil(t, err)
			assertEqualInts(t, got1, test.wantedDifferencesOf1)
			assertEqualInts(t, got3, test.wantedDifferencesOf3)
		})
	}
}

func Test_CountPossibleCombinations(t *testing.T) {
	tests := []struct {
		adapters      []int
		startingPoint int
		cache         map[int]int
		want          int
	}{
		{[]int{1}, 0, map[int]int{}, 1},
		{[]int{1, 2, 3, 4}, 2, map[int]int{2: 99}, 99},
		{[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}, 0, map[int]int{}, 8},
		{[]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}, 0, map[int]int{}, 19208},
	}

	for _, test := range tests {
		t.Run("CountPossibleCombinations("+fmt.Sprint(test.adapters)+", "+fmt.Sprint(test.startingPoint)+", "+fmt.Sprint(test.cache)+")", func(t *testing.T) {
			adapters := PrepareInput(test.adapters)
			got := CountPossibleCombinations(adapters, test.startingPoint, test.cache)
			assertEqualInts(t, got, test.want)
		})
	}
}
