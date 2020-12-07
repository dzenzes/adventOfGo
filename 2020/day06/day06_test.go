package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_GetGroupsFromInput(t *testing.T) {
	tests := []struct {
		input  []string
		wanted []GroupWithSize
	}{
		{[]string{
			"abc", "", "a", "b", "c", "", "ab", "ac", "", "a", "a", "a", "a", "", "b",
		}, []GroupWithSize{
			{"abc", 1}, {"abc", 3}, {"abac", 2}, {"aaaa", 4}, {"b", 1},
		}},
	}
	for i, test := range tests {
		t.Run("GetGroupsFromInput() "+fmt.Sprint(i), func(t *testing.T) {
			input := test.input
			want := test.wanted
			got := GetGroupsFromInput(input)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetGroupsFromInput() failed. Wanted %v but got %v", want, got)
			}
		})
	}
}

func Test_Unique(t *testing.T) {
	tests := []struct {
		input  string
		wanted string
	}{
		{"abc", "abc"},
		{"abac", "abc"},
		{"aaaa", "a"},
		{"b", "b"},
	}

	for _, test := range tests {
		t.Run("Unique("+test.input+")", func(t *testing.T) {
			input := test.input
			want := test.wanted
			got := Unique(input)
			if got != want {
				t.Errorf("Unique() failed. Wanted %v but got %v", want, got)
			}
		})
	}
}

func Test_CountCommonAnswers(t *testing.T) {
	tests := []struct {
		input  GroupWithSize
		wanted int
	}{
		{GroupWithSize{"abc", 1}, 3},
		{GroupWithSize{"abac", 2}, 1},
		{GroupWithSize{"aaaa", 4}, 1},
		{GroupWithSize{"b", 1}, 1},
	}

	for _, test := range tests {
		t.Run("CountCommonAnswers("+test.input.answers+", "+fmt.Sprint(test.input.size)+")", func(t *testing.T) {
			input := test.input
			want := test.wanted
			got := CountCommonAnswers(input)
			if got != want {
				t.Errorf("CountCommonAnswers() failed. Wanted %v but got %v", want, got)
			}
		})
	}
}

func Test_CountCharsPerLine(t *testing.T) {
	tests := []struct {
		input  []string
		wanted []int
	}{
		{[]string{
			"abc",
			"abc",
			"a",
			"b",
		}, []int{3, 3, 1, 1}},
	}

	for i, test := range tests {
		t.Run("CountCharsPerLine() "+fmt.Sprint(i), func(t *testing.T) {
			input := test.input
			want := test.wanted
			got := CountCharsPerLine(input)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("CountCharsPerLine() failed. Wanted %v but got %v", want, got)
			}
		})
	}
}

func Test_GetSumOfAnswers(t *testing.T) {
	tests := []struct {
		input  []string
		wanted int
	}{
		{[]string{"abc", "", "a", "b", "c", "", "ab", "ac", "", "a", "a", "a", "a", "", "b"}, 11},
	}

	for i, test := range tests {
		t.Run("GetSumOfAnswers() "+fmt.Sprint(i), func(t *testing.T) {
			input := test.input
			want := test.wanted
			got := GetSumOfAnswers(input)
			if got != want {
				t.Errorf("GetSumOfAnswers() failed. Wanted %v but got %v", want, got)
			}
		})
	}
}

func Test_GetSumOfCommonAnswers(t *testing.T) {
	tests := []struct {
		input  []string
		wanted int
	}{
		{[]string{"abc", "", "a", "b", "c", "", "ab", "ac", "", "a", "a", "a", "a", "", "b"}, 6},
	}

	for i, test := range tests {
		t.Run("GetSumOfCommonAnswers() "+fmt.Sprint(i), func(t *testing.T) {
			input := test.input
			want := test.wanted
			got := GetSumOfCommonAnswers(input)
			if got != want {
				t.Errorf("GetSumOfCommonAnswers() failed. Wanted %v but got %v", want, got)
			}
		})
	}
}
