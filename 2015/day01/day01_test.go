package main

import (
	"fmt"
	"testing"
)

func Test_FindFloor(t *testing.T) {
	tests := []struct {
		instruction string
		want        int
	}{

		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, test := range tests {
		t.Run("FindFloor("+test.instruction+") => "+fmt.Sprint(test.want), func(t *testing.T) {
			input := test.instruction
			want := test.want
			got := FindFloor(input)
			if got != want {
				t.Errorf("FindFloor() returned %v but %v was wanted", got, want)
			}
		})
	}
}
func Test_FirstTimeInBasement(t *testing.T) {
	tests := []struct {
		instruction string
		want        int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, test := range tests {
		t.Run("FirstTimeInBasement("+test.instruction+") => "+fmt.Sprint(test.want), func(t *testing.T) {
			input := test.instruction
			want := test.want
			got := FirstTimeInBasement(input)
			if got != want {
				t.Errorf("FindFloor() returned %v but %v was wanted", got, want)
			}
		})
	}
}
