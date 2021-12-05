package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandleTestInput(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2"}
	want := Position{
		x: 15,
		y: 10,
	}

	submarine := NewSubmarine()
	got, err := submarine.calculatePosition(input)

	if err != nil {
		t.Errorf("couldn't calculate position %v", err)
	}
	assert.Equal(t, want, got)
}

func Test_ParseInput(t *testing.T) {
	tests := []struct {
		input    string
		command  string
		distance int
	}{
		{"forward 5", "forward", 5},
		{"forward 15", "forward", 15},
	}
	for _, test := range tests {
		t.Run("ParseInput("+fmt.Sprint(test.input)+")", func(t *testing.T) {
			command, distance, err := ParseInput(test.input)
			if err != nil {
				t.Errorf("got an error %v", err)
			}
			assert.Equal(t, test.command, command)
			assert.Equal(t, test.distance, distance)
		})

	}
}

func Test_ParseInputThrowsErrorOnIllegalInput(t *testing.T) {
	input := "forward one"
	_, _, err := ParseInput(input)
	assert.True(t, err != nil)
}

func Test_CalculatePosition(t *testing.T) {
	tests := []struct {
		commands []string
		position Position
	}{
		{[]string{}, Position{x: 0, y: 0}},
		{[]string{"forward 5"}, Position{x: 5, y: 0}},
		{[]string{"down 5"}, Position{x: 0, y: 5}},
		{[]string{"up 5"}, Position{x: 0, y: -5}},
	}

	for _, test := range tests {
		t.Run("CalculatePosition("+fmt.Sprint(test.commands)+")", func(t *testing.T) {

			submarine := NewSubmarine()
			got, err := submarine.calculatePosition(test.commands)

			if err != nil {
				t.Errorf("couldn't calculate position %v", err)
			}
			assert.Equal(t, test.position, got)
		})

	}
}

func Test_HandleTestInputPart2(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2"}
	want := Position{
		x: 15,
		y: 60,
	}

	submarine := NewSubmarineWithAim()
	got, err := submarine.calculatePosition(input)

	if err != nil {
		t.Errorf("couldn't calculate position %v", err)
	}
	assert.Equal(t, want, got)
}
func Test_CalculatePositionWithAim(t *testing.T) {
	tests := []struct {
		commands []string
		position Position
		aimValue int
	}{
		{[]string{}, Position{x: 0, y: 0}, 0},
		{[]string{"forward 5"}, Position{x: 5, y: 0}, 0},
		{[]string{"down 5"}, Position{x: 0, y: 0}, 5},
		{[]string{"up 5"}, Position{x: 0, y: 0}, -5},
		{[]string{"forward 5", "down 5", "forward 8"}, Position{x: 13, y: 40}, 5},
	}

	for _, test := range tests {
		t.Run("CalculatePosition("+fmt.Sprint(test.commands)+")", func(t *testing.T) {

			submarine := NewSubmarineWithAim()
			got, err := submarine.calculatePosition(test.commands)

			if err != nil {
				t.Errorf("couldn't calculate position %v", err)
			}
			assert.Equal(t, test.aimValue, submarine.aimValue)
			assert.Equal(t, test.position, got)
		})

	}
}
