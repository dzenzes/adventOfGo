package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SampleIsRunningPart1(t *testing.T) {
	input := []int{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}

	want := 7
	got := CountNumberofIncreasedMeasurements(input)
	assert.Equal(t, want, got)
}

func Test_SomeTestDataIsWorking(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 0},
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2, 4, 3, 2}, 2},
	}

	for _, test := range tests {
		t.Run("CountNumberOfIncreasedMeasurements("+fmt.Sprint(test.input)+")", func(t *testing.T) {
			got := CountNumberofIncreasedMeasurements(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_SampleIsRunningPart2(t *testing.T) {
	input := []int{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}

	want := 5
	got := CountNumberofIncreasedMeasurementsWithSlidingWindow(input)
	assert.Equal(t, want, got)
}

func Test_ToSlidingWindows(t *testing.T) {

	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{}},

		{[]int{1}, []int{}},
		{[]int{1, 2}, []int{}},
		{[]int{1, 2, 3}, []int{6}},
		{[]int{1, 2, 3, 4}, []int{6, 9}},
		{[]int{199,
			200,
			208,
			210,
			200,
			207,
			240,
			269,
			260,
			263}, []int{607, 618, 618, 617, 647, 716, 769, 792}},
	}

	for _, test := range tests {
		t.Run("ToSlidingWindows("+fmt.Sprint(test.input)+")", func(t *testing.T) {
			got := ToSlidingWindows(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
