package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetSpokenNumber(t *testing.T) {
	tests := []struct {
		input          []int
		searchedNumber int
		want           int
	}{

		{[]int{0, 3, 6}, 1, 0},
		{[]int{0, 3, 6}, 2, 3},
		{[]int{0, 3, 6}, 4, 0},
		{[]int{0, 3, 6}, 5, 3},
		{[]int{0, 3, 6}, 6, 3},
		{[]int{0, 3, 6}, 7, 1},
		{[]int{0, 3, 6}, 8, 0},
		{[]int{0, 3, 6}, 9, 4},
		{[]int{0, 3, 6}, 10, 0},
		{[]int{0, 3, 6}, 2020, 436},
	}

	for _, test := range tests {
		t.Run("GetSpokenNumber("+fmt.Sprint(test.input)+", "+fmt.Sprint(test.searchedNumber)+")", func(t *testing.T) {
			fmt.Println("GetSpokenNumber(" + fmt.Sprint(test.input) + ", " + fmt.Sprint(test.searchedNumber) + ")")
			got := GetSpokenNumber(test.input, test.searchedNumber)

			assert.Equal(t, test.want, got)
		})
	}
}

func Test_WriteToCache(t *testing.T) {
	tests := []struct {
		key      int
		value    int
		oldCache map[int][]int
		want     map[int][]int
	}{
		{0, 4, map[int][]int{
			0: {0},
			3: {1},
			4: {0},
			5: {0},
			6: {2},
		}, map[int][]int{
			0: {0, 4},
			3: {1},
			4: {0},
			5: {0},
			6: {2},
		},
		},
	}

	for _, test := range tests {
		t.Run("WriteToCache", func(t *testing.T) {
			got := WriteToCache(test.key, test.value, test.oldCache)
			assert.EqualValues(t, test.want, got)
		})
	}
}
