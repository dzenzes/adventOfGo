package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
	want := 5

	got := CountOverlappingLines(input)

	assert.Equal(t, want, got)
}

func Test_Part2(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
	want := 12

	got := CountOverlappingLinesIncludingDiagonals(input)

	assert.Equal(t, want, got)
}
func Test_ParseLine(t *testing.T) {
	input := "0,9 -> 5,9"
	want1 := Point{0, 9}
	want2 := Point{5, 9}

	got1, got2, err := ParseLine(input)

	assert.NoError(t, err, "parsing should pass")
	assert.Equal(t, want1, got1)
	assert.Equal(t, want2, got2)
}

func Test_PointIsNoDiagonalLine(t *testing.T) {
	tests := []struct {
		p    Point
		q    Point
		want bool
	}{
		{Point{0, 9}, Point{1, 8}, false},
		{Point{0, 9}, Point{0, 8}, true},
		{Point{1, 9}, Point{0, 9}, true},
	}

	for _, test := range tests {
		t.Run("isNoDiagonalLine ("+fmt.Sprint(test.p)+", "+fmt.Sprint(test.q)+")", func(t *testing.T) {
			got := test.p.isNoDiagonalLine(test.q)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_PointsOnLine(t *testing.T) {
	tests := []struct {
		p    Point
		q    Point
		want []Point
	}{
		{Point{0, 8}, Point{3, 9}, []Point{}},
		{Point{1, 1}, Point{1, 3}, []Point{{1, 1}, {1, 2}, {1, 3}}},
		{Point{0, 9}, Point{3, 9}, []Point{{0, 9}, {1, 9}, {2, 9}, {3, 9}}},
		{Point{2, 3}, Point{4, 3}, []Point{{2, 3}, {3, 3}, {4, 3}}},
		{Point{4, 3}, Point{2, 3}, []Point{{4, 3}, {3, 3}, {2, 3}}},
	}

	for _, test := range tests {
		t.Run("PointsOnLine ("+fmt.Sprint(test.p)+", "+fmt.Sprint(test.q)+")", func(t *testing.T) {
			got := PointsOnLine(test.p, test.q, false)
			assert.Equal(t, test.want, got)
		})
	}
}
func Test_PointsOnLineInclDiagonals(t *testing.T) {
	tests := []struct {
		p    Point
		q    Point
		want []Point
	}{
		{Point{0, 0}, Point{3, 3}, []Point{{0, 0}, {1, 1}, {2, 2}, {3, 3}}},
	}

	for _, test := range tests {
		t.Run("PointsOnLine ("+fmt.Sprint(test.p)+", "+fmt.Sprint(test.q)+")", func(t *testing.T) {
			got := PointsOnLine(test.p, test.q, true)
			assert.Equal(t, test.want, got)
		})
	}
}
