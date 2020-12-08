package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Parse(t *testing.T) {
	input := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	want := []Instruction{
		{NoOperation, +0},
		{Accumulate, +1},
		{Jump, +4},
		{Accumulate, +3},
		{Jump, -3},
		{Accumulate, -99},
		{Accumulate, +1},
		{Jump, -4},
		{Accumulate, +6},
	}

	got, err := Parse(input)
	if err != nil {
		t.Errorf("Parse() threw an error %v", err)
	} else if !reflect.DeepEqual(got, want) {
		t.Errorf("Parse() should return %v but it returned %v", got, want)
	}
}

func Test_ProcessInstruction(t *testing.T) {
	tests := []struct {
		input       HandheldState
		instruction Instruction
		want        HandheldState
	}{
		{HandheldState{0, 0, []int{}}, Instruction{Accumulate, 7}, HandheldState{7, 1, []int{0}}},
		{HandheldState{3, 5, []int{1, 2, 3}}, Instruction{Accumulate, 3}, HandheldState{6, 6, []int{1, 2, 3, 5}}},
		{HandheldState{13, 0, []int{}}, Instruction{Accumulate, -21}, HandheldState{-8, 1, []int{0}}},
		{HandheldState{42, 50, []int{}}, Instruction{Jump, 50}, HandheldState{42, 100, []int{50}}},
		{HandheldState{42, 13, []int{}}, Instruction{Jump, -5}, HandheldState{42, 8, []int{13}}},
		{HandheldState{42, 13, []int{}}, Instruction{NoOperation, 0}, HandheldState{42, 14, []int{13}}},
	}

	for _, test := range tests {
		t.Run("ProcessInstruction() "+fmt.Sprint(test.instruction), func(t *testing.T) {
			input := test.input
			instruction := test.instruction
			want := test.want

			got, err := ProcessInstruction(instruction, input)
			if err != nil {
				t.Errorf("ProcessInstruction() threw an error %v", err)
			} else if !reflect.DeepEqual(got, want) {
				t.Errorf("ProcessInstruction() should return %v but it returned %v", want, got)
			}
		})
	}
}

func Test_HandheldState_contains(t *testing.T) {
	tests := []struct {
		handheldState HandheldState
		pointer       int
		want          bool
	}{
		{HandheldState{5, 1, []int{0, 1, 2, 3, 4}}, 5, false},
		{HandheldState{5, 1, []int{0, 1, 2, 3, 4}}, 0, true},
		{HandheldState{5, 1, []int{0, 1, 2, 3, 4}}, 4, true},
	}

	for i, test := range tests {
		t.Run("HandheldState.contains() "+fmt.Sprint(i), func(t *testing.T) {
			state := test.handheldState
			input := test.pointer
			want := test.want

			got := state.contains(input)
			if got != want {
				t.Errorf("HandheldState.contains() should return %v but it returned %v", want, got)
			}
		})
	}
}

func Test_Process_with_infinite_loop(t *testing.T) {
	instructions := []Instruction{
		{NoOperation, +0},
		{Accumulate, +1},
		{Jump, +4},
		{Accumulate, +3},
		{Jump, -3},
		{Accumulate, -99},
		{Accumulate, +1},
		{Jump, -4},
		{Accumulate, +6},
	}

	want := HandheldState{5, 1, []int{0, 1, 2, 6, 7, 3, 4}}

	got, err := Process(instructions)
	if err == nil {
		t.Errorf("Process() should throw an error %v because the program doesn't end", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Process() should return %v (even if there was an error) but it returned %v", want, got)
	}
}

func Test_Process(t *testing.T) {
	instructions := []Instruction{
		{NoOperation, +0},
		{Accumulate, +1},
		{Jump, +4},
		{Accumulate, +3},
		{Jump, -3},
		{Accumulate, -99},
		{Accumulate, +1},
		{NoOperation, -4},
		{Accumulate, +6},
	}

	want := HandheldState{8, 9, []int{0, 1, 2, 6, 7, 8}}

	got, err := Process(instructions)
	if err != nil {
		t.Errorf("Process() threw an error %v ", err)
	} else if !reflect.DeepEqual(got, want) {
		t.Errorf("Process() should return %v (even if there was an error) but it returned %v", want, got)
	}
}

func Test_GetInstructionsToSwitch(t *testing.T) {
	instructions := []Instruction{
		{NoOperation, +0},
		{Accumulate, +1},
		{Jump, +4},
		{Accumulate, +3},
		{Jump, -3},
		{Accumulate, -99},
		{Accumulate, +1},
		{Jump, -4},
		{Accumulate, +6},
	}

	want := []int{0, 2, 4, 7}

	got := GetInstructionsToSwitch(instructions)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetInstructionsToSwitch() should return %v (even if there was an error) but it returned %v", want, got)
	}
}

func Test_CreateInstructionsCandidate(t *testing.T) {
	instructions := []Instruction{
		{NoOperation, +0},
		{Accumulate, +1},
		{Jump, +4},
		{Accumulate, +3},
		{Jump, -3},
		{Accumulate, -99},
		{Accumulate, +1},
		{Jump, -4},
		{Accumulate, +6},
	}

	tests := []struct {
		input int
		want  []Instruction
	}{
		{
			0,
			[]Instruction{
				{Jump, +0},
				{Accumulate, +1},
				{Jump, +4},
				{Accumulate, +3},
				{Jump, -3},
				{Accumulate, -99},
				{Accumulate, +1},
				{Jump, -4},
				{Accumulate, +6},
			},
		},
		{
			7,
			[]Instruction{
				{NoOperation, +0},
				{Accumulate, +1},
				{Jump, +4},
				{Accumulate, +3},
				{Jump, -3},
				{Accumulate, -99},
				{Accumulate, +1},
				{NoOperation, -4},
				{Accumulate, +6},
			},
		},
	}

	for i, test := range tests {
		t.Run("CreateInstructionsCandidate() "+fmt.Sprint(i), func(t *testing.T) {
			want := test.want
			got := CreateInstructionsCandidate(test.input, instructions)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("CreateInstructionsCandidate(%v) should return %v but it returned %v", test.input, want, got)
			}
		})
	}
}

func Test_FindCorrectInstructionsAndProcess(t *testing.T) {
	instructions := []Instruction{
		{NoOperation, +0},
		{Accumulate, +1},
		{Jump, +4},
		{Accumulate, +3},
		{Jump, -3},
		{Accumulate, -99},
		{Accumulate, +1},
		{Jump, -4},
		{Accumulate, +6},
	}

	want := HandheldState{8, 9, []int{0, 1, 2, 6, 7, 8}}

	got, err := FindCorrectInstructionsAndProcess(instructions)
	if err != nil {
		t.Errorf("FindCorrectInstructionsAndProcess() threw an error %v ", err)
	} else if !reflect.DeepEqual(got, want) {
		t.Errorf("FindCorrectInstructionsAndProcess() should return %v but it returned %v", want, got)
	}
}
