package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_FindSum(t *testing.T) {
	input := 5
	inputList := []int{1, 2, 3, 4, 5}
	want := true

	got := FindSum(input, inputList)
	if got != want {
		t.Errorf("FindSum() returned %v but %v was wanted", got, want)
	}
}

func Test_FindFirstWrongNumber(t *testing.T) {
	input := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	preambleSize := 5
	want := 127

	got := FindFirstWrongNumber(input, preambleSize)
	if got != want {
		t.Errorf("FindWrongNumber() returned %v but %v was wanted", got, want)
	}
}

func Test_FindContiguousSetOfNumbersThatSumUpTo(t *testing.T) {
	input := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	searchedNumber := 127
	want := []int{15, 25, 47, 40}

	got, err := FindContiguousSetOfNumbersThatSumUpTo(searchedNumber, input)
	if err != nil {
		t.Errorf("FindContiguousSetOfNumbersThatSumUpTo() throws error %v", err)
	} else if !reflect.DeepEqual(got, want) {
		t.Errorf("FindContiguousSetOfNumbersThatSumUpTo() returned %v but %v was wanted", got, want)
	}
}

func Test_GetMinAndMaxFromList(t *testing.T) {
	tests := []struct {
		input []int
		min   int
		max   int
	}{
		{
			[]int{6, 4, 3, 7, 8, 1}, 1, 8,
		}, {
			[]int{1, 1}, 1, 1,
		}, {
			[]int{100}, 100, 100,
		},
	}

	for i, test := range tests {
		t.Run("GetMinAndMaxFromList() "+fmt.Sprint(i), func(t *testing.T) {
			gotMin, gotMax := GetMinAndMaxFromList(test.input)

			if test.min != gotMin {
				t.Errorf("GetMinAndMaxFromList() returned min %v but %v was wanted", gotMin, test.min)
			}
			if test.max != gotMax {
				t.Errorf("GetMinAndMaxFromList() returned max %v but %v was wanted", gotMax, test.max)
			}

		})
	}
}

func Test_FindEncryptionWeakness(t *testing.T) {
	input := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	illegalNumber := 127
	want := 62

	got, err := FindEncryptionWeakness(illegalNumber, input)
	if err != nil {
		t.Errorf("FindEncryptionWeakness() throws error %v", err)
	} else if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEncryptionWeakness() returned %v but %v was wanted", got, want)
	}
}
