package main

import (
	"reflect"
	"testing"
)

func Test_Parse(t *testing.T) {
	tests := []struct {
		input string
		want  Instruction
	}{
		{"F10", Instruction{"F", 10}},
		{"N3", Instruction{"N", 3}},
		{"F7", Instruction{"F", 7}},
		{"R90", Instruction{"R", 90}},
		{"F11", Instruction{"F", 11}},
	}

	for _, test := range tests {
		t.Run("Parse("+test.input+")", func(t *testing.T) {
			got, err := Parse(test.input)
			if err != nil {
				t.Errorf("got an error %v", err)
			}
			if got != test.want {
				t.Errorf("got %v but wanted %v", got, test.want)
			}
		})
	}
}

func Test_ParseList(t *testing.T) {
	input := []string{"F10", "N3", "F7", "R90", "F11"}
	want := []Instruction{Instruction{"F", 10}, Instruction{"N", 3}, Instruction{"F", 7}, Instruction{"R", 90}, Instruction{"F", 11}}
	got, err := ParseList(input)
	if err != nil {
		t.Errorf("got an error %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

func Test_CalculateDirection_fails_if_value_Cannot_be_divided_by_90(t *testing.T) {
	oldDirection := "N"
	instruction := Instruction{"L", 91}

	_, err := CalculateDirection(oldDirection, instruction)

	if err == nil {
		t.Errorf("CalculateDirection should return an error if the value is not supported")
	} else if err.Error() != UnsupportedValueError {
		t.Errorf("expected '%v' but got '%v'", UnsupportedValueError, err)
	}
}

func Test_CalculateDirection(t *testing.T) {

	tests := []struct {
		oldDirection string
		instruction  Instruction
		want         string
	}{
		{"N", Instruction{"L", 90}, "W"},
		{"N", Instruction{"R", 270}, "W"},
		{"N", Instruction{"R", 180}, "S"},
		{"N", Instruction{"R", 450}, "E"},
	}

	for _, test := range tests {
		t.Run("CalculateDirection supports "+test.instruction.action, func(t *testing.T) {
			got, err := CalculateDirection(test.oldDirection, test.instruction)

			if err != nil {
				t.Errorf("CalculateDirection returned an error '%v'", err)
			} else if got != test.want {
				t.Errorf("got %v but wanted %v", got, test.want)
			}
		})
	}
}

func Test_ApplyInstruction_fails_on_unsupported_operations(t *testing.T) {
	instruction := Instruction{"Q", 123}
	ship := Ship{0, 0, "N"}

	_, err := ApplyInstruction(ship, instruction)
	if err == nil {
		t.Errorf("ApplyInstruction should return an error if the instruction is not supported")
	} else if err.Error() != IllegalInstructionError {
		t.Errorf("expected '%v' but got '%v'", IllegalInstructionError, err)
	}
}

func Test_ApplyInstruction(t *testing.T) {

	tests := []struct {
		instruction Instruction
		ship        Ship
		want        Ship
	}{
		{
			Instruction{"N", 123},
			Ship{0, 0, "N"},
			Ship{123, 0, "N"},
		}, {
			Instruction{"S", 123},
			Ship{100, 0, "N"},
			Ship{-23, 0, "N"},
		}, {
			Instruction{"E", 123},
			Ship{100, 0, "N"},
			Ship{100, 123, "N"},
		}, {
			Instruction{"W", 123},
			Ship{100, 10, "N"},
			Ship{100, -113, "N"},
		}, {
			Instruction{"L", 90},
			Ship{100, 10, "N"},
			Ship{100, 10, "W"},
		}, {
			Instruction{"R", 3150},
			Ship{100, 10, "N"},
			Ship{100, 10, "W"},
		}, {
			Instruction{"F", 120},
			Ship{100, 10, "N"},
			Ship{220, 10, "N"},
		},
	}

	for _, test := range tests {

		t.Run("ApplyInstruction supports '"+test.instruction.action+"' actions on Instructions", func(t *testing.T) {
			instruction := test.instruction
			ship := test.ship
			want := test.want

			got, err := ApplyInstruction(ship, instruction)
			if err != nil {
				t.Errorf("ApplyInstruction returned an error '%v'", err)
			} else if got != want {
				t.Errorf("got %v but wanted %v", got, want)
			}
		})
	}
}

func Test_ApplyInstructionWithWaypoint(t *testing.T) {

	tests := []struct {
		instruction    Instruction
		ship           Ship
		waypoint       Waypoint
		wantedShip     Ship
		wantedWaypoint Waypoint
	}{
		{
			Instruction{"F", 10},
			Ship{0, 0, "E"},
			Waypoint{1, 10},
			Ship{10, 100, "E"},
			Waypoint{1, 10},
		}, {
			Instruction{"N", 3},
			Ship{10, 100, "E"},
			Waypoint{1, 10},
			Ship{10, 100, "E"},
			Waypoint{4, 10},
		}, {
			Instruction{"E", 3},
			Ship{10, 100, "E"},
			Waypoint{1, 10},
			Ship{10, 100, "E"},
			Waypoint{1, 13},
		}, {
			Instruction{"S", 3},
			Ship{10, 100, "E"},
			Waypoint{1, 10},
			Ship{10, 100, "E"},
			Waypoint{-2, 10},
		}, {
			Instruction{"W", 3},
			Ship{10, 100, "E"},
			Waypoint{1, 10},
			Ship{10, 100, "E"},
			Waypoint{1, 7},
		}, {
			Instruction{"F", 7},
			Ship{10, 100, "E"},
			Waypoint{4, 10},
			Ship{38, 170, "E"},
			Waypoint{4, 10},
		}, {
			Instruction{"R", 90},
			Ship{38, 170, "E"},
			Waypoint{4, 10},
			Ship{38, 170, "E"},
			Waypoint{-10, 4},
		}, {
			Instruction{"R", 180},
			Ship{38, 170, "E"},
			Waypoint{4, 10},
			Ship{38, 170, "E"},
			Waypoint{-4, -10},
		}, {
			Instruction{"R", 270},
			Ship{38, 170, "E"},
			Waypoint{4, 10},
			Ship{38, 170, "E"},
			Waypoint{10, -4},
		}, {
			Instruction{"R", 360},
			Ship{38, 170, "E"},
			Waypoint{4, 10},
			Ship{38, 170, "E"},
			Waypoint{4, 10},
		}, {
			Instruction{"L", 270},
			Ship{38, 170, "E"},
			Waypoint{4, 10},
			Ship{38, 170, "E"},
			Waypoint{-10, 4},
		},
	}

	for _, test := range tests {

		t.Run("ApplyInstructionWithWaypoint supports '"+test.instruction.action+"' actions on Instructions", func(t *testing.T) {
			instruction := test.instruction
			ship := test.ship
			waypoint := test.waypoint
			wantedShip := test.wantedShip
			wantedWaypoint := test.wantedWaypoint

			gotShip, gotWaypoint, err := ApplyInstructionWithWaypoint(ship, waypoint, instruction)
			if err != nil {
				t.Errorf("ApplyInstructionWithWaypoint returned an error '%v'", err)
			}
			if gotShip != wantedShip {
				t.Errorf("got ship %v but wanted %v", gotShip, wantedShip)
			}
			if gotWaypoint != wantedWaypoint {
				t.Errorf("got waypoint %v but wanted %v", gotWaypoint, wantedWaypoint)
			}
		})
	}
}

func Test_Process(t *testing.T) {
	input := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	t.Run("Integration test for part 1", func(t *testing.T) {
		want := 25

		ship, err := Process(input)

		if err != nil {
			t.Errorf("Process returned an error '%v'", err)
		}
		got := ship.getManhattanDistance()
		if got != want {
			t.Errorf("got %v but wanted %v", got, want)
		}
	})
	t.Run("Integration test for part 2", func(t *testing.T) {
		want := 286

		ship, err := ProcessWithWaypoint(input)

		if err != nil {
			t.Errorf("ProcessWithWaypoint returned an error '%v'", err)
		}
		got := ship.getManhattanDistance()
		if got != want {
			t.Errorf("got %v but wanted %v", got, want)
		}
	})
}
