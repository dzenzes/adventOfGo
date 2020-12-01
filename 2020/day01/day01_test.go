package main

import "testing"

func Test_FindNumbersThatSumTo2020(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	want1 := 1721
	want2 := 299

	got1, got2, err := FindNumbersThatSumTo2020(input)
	if err != nil {
		t.Errorf("FindNumbersThatSumTo2020() error = %v", err)
	}
	if want1 != got1 || want2 != got2 {
		t.Errorf("FindNumbersThatSumTo2020: got %v, %v want %v, %v", got1, got2, want1, want2)
	}
}

func Test_FindThreeNumbersThatSumTo2020(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	want1 := 979
	want2 := 366
	want3 := 675

	got1, got2, got3, err := FindThreeNumbersThatSumTo2020(input)
	if err != nil {
		t.Errorf("FindThreeNumbersThatSumTo2020() error = %v", err)
	}
	if want1 != got1 || want2 != got2 || want3 != got3 {
		t.Errorf("FindThreeNumbersThatSumTo2020: got %v, %v, %v want %v, %v,%v", got1, got2, got3, want1, want2, want3)
	}
}

func Test_Part1(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	want := 514579

	got, err := Part1(input)
	if err != nil {
		t.Errorf("Part1() error = %v", err)
	}
	if want != got {
		t.Errorf("Part1: got %v want %v", got, want)
	}
}

func Test_Part2(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	want := 241861950

	got, err := Part2(input)
	if err != nil {
		t.Errorf("Part2() error = %v", err)
	}
	if want != got {
		t.Errorf("Part2: got %v want %v", got, want)
	}
}
