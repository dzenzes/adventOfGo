package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CountOneInBinaries(t *testing.T) {

	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	tests := []struct {
		position int
		want     int
	}{
		{0, 7},
		{1, 5},
		{2, 8},
		{3, 7},
		{4, 5},
	}

	for _, test := range tests {
		t.Run("CountOneInBinaries at position "+fmt.Sprint(test.position), func(t *testing.T) {
			got := CountOneInBinaries(input, test.position)
			assert.Equal(t, test.want, got)
		})

	}
}

func Test_CalculateMostCommonBit(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	tests := []struct {
		position int
		want     string
	}{
		{0, "1"},
		{1, "0"},
		{2, "1"},
		{3, "1"},
		{4, "0"},
	}

	for _, test := range tests {
		t.Run("GetMostCommonBit at position"+fmt.Sprint(test.position), func(t *testing.T) {
			got := GetMostCommonBit(input, test.position)
			assert.Equal(t, test.want, got)
		})

	}
	want := "1"
	got := GetMostCommonBit(input, 0)
	assert.Equal(t, want, got)
}

func Test_CalculateGammaAndEpsilonRate(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	want := 22
	want2 := 9
	gamma, epsilon, err := CalculateGammaAndEpsilonRate(input)
	if err != nil {
		t.Errorf("couldn't calculate gamma and epsilon rate")
	}
	assert.Equal(t, want, gamma)
	assert.Equal(t, want2, epsilon)
}
func Test_CalculatePowerConsumption(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	want := 198
	got, err := CalculatePowerConsumption(input)

	if err != nil {
		t.Errorf("couldn't calculate power consumption")
	}
	assert.Equal(t, want, got)
}

func Test_FilterByBit(t *testing.T) {

	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	tests := []struct {
		bit      string
		position int
		want     []string
	}{
		{
			"0",
			0,
			[]string{
				"00100",
				"01111",
				"00111",
				"00010",
				"01010",
			},
		},
		{
			"1",
			1,
			[]string{
				"11110",
				"01111",
				"11100",
				"11001",
				"01010",
			},
		},
	}

	for _, test := range tests {
		t.Run("FilterByBit("+fmt.Sprint(test.bit)+") at position "+fmt.Sprint(test.position), func(t *testing.T) {
			got := FilterByBit(input, test.bit, test.position)
			assert.Equal(t, test.want, got)
		})

	}
}

func Test_CalculateOxygenGeneratorRating(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	want := 23
	got, err := CalculateOxygenGeneratorRating(input)

	if err != nil {
		t.Errorf("couldn't calculate oxygen")
	}
	assert.Equal(t, want, got)
}
func Test_CalculateCO2ScrubberRating(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	want := 10
	got, err := CalculateCO2ScrubberRating(input)

	if err != nil {
		t.Errorf("couldn't calculate oxygen")
	}
	assert.Equal(t, want, got)
}
func Test_CalculateLifeSupport(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	want := 230
	got, err := CalculateLifeSupport(input)

	if err != nil {
		t.Errorf("couldn't calculate power consumption")
	}
	assert.Equal(t, want, got)
}
