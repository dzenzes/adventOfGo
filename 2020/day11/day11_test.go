package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_ParsePositionIllegalInput(t *testing.T) {
	tests := []string{
		"",
		"P",
		"##",
	}

	for _, test := range tests {
		t.Run("ParsePosition("+test+") should return an error", func(t *testing.T) {
			_, err := ParsePosition(test)
			if err == nil {

				t.Error("Expected ParsePosition to return an error")
			}

			if "couldn't parse input" != fmt.Sprint(err) {
				t.Error("wrong error message")
			}
		})
	}
}

func Test_ParsePosition(t *testing.T) {
	tests := []struct {
		input string
		want  Seat
	}{
		{"#", Seat{true, true}},
		{"L", Seat{true, false}},
		{".", Seat{false, false}},
	}

	for _, test := range tests {
		t.Run("ParsePosition("+test.input+")", func(t *testing.T) {
			got, err := ParsePosition(test.input)
			if err != nil {
				t.Errorf("method returned an error %v", err)
			}
			if got != test.want {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

func createSeat() Seat {
	return Seat{true, false}
}

func createFloor() Seat {
	return Seat{false, false}
}

func createOccupiedSeat() Seat {
	return Seat{true, true}
}

func Test_ParseMap(t *testing.T) {
	input := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
	}

	want := [][]Seat{
		{createSeat(), createFloor(), createSeat(), createSeat(), createFloor(), createSeat(), createSeat(), createFloor(), createSeat(), createSeat()},
		{createSeat(), createSeat(), createSeat(), createSeat(), createSeat(), createSeat(), createSeat(), createFloor(), createSeat(), createSeat()},
	}

	got, err := ParseMap(input)
	if err != nil {
		t.Errorf("method returned an error %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}

func Test_CountOccupiedNeighbors(t *testing.T) {
	seats := [][]Seat{
		{createSeat(), createFloor(), createSeat(), createOccupiedSeat(), createFloor(), createOccupiedSeat(), createOccupiedSeat(), createFloor(), createSeat(), createSeat()},
		{createSeat(), createSeat(), createOccupiedSeat(), createSeat(), createSeat(), createOccupiedSeat(), createOccupiedSeat(), createFloor(), createSeat(), createSeat()},
	}

	tests := []struct {
		x    int
		y    int
		want int
	}{
		{0, 0, 0},
		{3, 1, 2},
		{6, 0, 3},
	}

	for _, test := range tests {
		t.Run("CountOccupiedNeighbors("+fmt.Sprint(test.x)+", "+fmt.Sprint(test.y)+")", func(t *testing.T) {
			got := CountOccupiedNeighbors(test.x, test.y, seats)
			if got != test.want {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

func Test_UpdateSeat(t *testing.T) {
	seats := [][]Seat{
		{createSeat(), createFloor(), createOccupiedSeat(), createOccupiedSeat(), createOccupiedSeat(), createSeat(), createSeat(), createFloor(), createSeat(), createSeat()},
		{createSeat(), createSeat(), createSeat(), createOccupiedSeat(), createSeat(), createSeat(), createSeat(), createFloor(), createSeat(), createSeat()},
	}

	tests := []struct {
		x    int
		y    int
		want Seat
	}{
		{0, 0, createOccupiedSeat()},
		{3, 1, createSeat()},
	}

	for _, test := range tests {
		t.Run("UpdateSeat("+fmt.Sprint(test.x)+", "+fmt.Sprint(test.y)+")", func(t *testing.T) {
			got := UpdateSeat(test.x, test.y, seats, 4, CountOccupiedNeighbors)

			if got != test.want {
				t.Errorf("got %v but want %v", got, test.want)
			}
		})
	}
}

func Test_ProcessRound(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{[]string{
			"L.LL.LL.LL",
			"LLLLLLL.LL",
			"L.L.L..L..",
			"LLLL.LL.LL",
			"L.LL.LL.LL",
			"L.LLLLL.LL",
			"..L.L.....",
			"LLLLLLLLLL",
			"L.LLLLLL.L",
			"L.LLLLL.LL",
		},
			[]string{
				"#.##.##.##",
				"#######.##",
				"#.#.#..#..",
				"####.##.##",
				"#.##.##.##",
				"#.#####.##",
				"..#.#.....",
				"##########",
				"#.######.#",
				"#.#####.##",
			},
		},
		{[]string{
			"#.##.##.##",
			"#######.##",
			"#.#.#..#..",
			"####.##.##",
			"#.##.##.##",
			"#.#####.##",
			"..#.#.....",
			"##########",
			"#.######.#",
			"#.#####.##",
		}, []string{
			"#.LL.L#.##",
			"#LLLLLL.L#",
			"L.L.L..L..",
			"#LLL.LL.L#",
			"#.LL.LL.LL",
			"#.LLLL#.##",
			"..L.L.....",
			"#LLLLLLLL#",
			"#.LLLLLL.L",
			"#.#LLLL.##",
		},
		},
	}
	for i, test := range tests {
		t.Run("ProcessRound() "+fmt.Sprint(i), func(t *testing.T) {
			input, err := ParseMap(test.input)
			if err != nil {
				t.Errorf("method returned an error %v", err)
			}
			want, err := ParseMap(test.want)
			if err != nil {
				t.Errorf("method returned an error %v", err)
			}

			got := ProcessRound(input, 4, CountOccupiedNeighbors)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v but want %v", got, want)
			}
		})
	}
}

func Test_CountVisibleOccupiedSeats(t *testing.T) {
	tests := []struct {
		x     int
		y     int
		seats []string
		want  int
	}{{3, 4, []string{
		".......#.",
		"...#.....",
		".#.......",
		".........",
		"..#L....#",
		"....#....",
		".........",
		"#........",
		"...#.....",
	}, 8},
		{7, 1, []string{
			".............",
			".L.L.#.#.#.#.",
			".............",
		}, 2},
		{1, 2, []string{
			".##.##.",
			"#.#.#.#",
			"##...##",
			"...L...",
			"##...##",
			"#.#.#.#",
			".##.##.",
		}, 7},
		{3, 3, []string{
			".##.##.",
			"#.#.#.#",
			"##...##",
			"...L...",
			"##...##",
			"#.#.#.#",
			".##.##.",
		}, 0},
	}

	for i, test := range tests {
		seatMap, _ := ParseMap(test.seats)
		t.Run("CountVisibleOccupiedSeats() "+fmt.Sprint(i), func(t *testing.T) {
			got := CountVisibleOccupiedSeats(test.x, test.y, seatMap)

			if got != test.want {
				t.Errorf("got %v but want %v", got, test.want)
			}
		})
	}
}

func Test_ProcessUntilNoChange(t *testing.T) {
	seats := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}

	t.Run("Example for part 1", func(t *testing.T) {
		want := 37
		got, err := ProcessUntilNoChange(seats, 4, CountOccupiedNeighbors)
		if err != nil {
			t.Errorf("ProcessUntilNoChange returned an error %v", err)
		}
		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("Example for part 2", func(t *testing.T) {
		want := 26
		got, err := ProcessUntilNoChange(seats, 5, CountVisibleOccupiedSeats)
		if err != nil {
			t.Errorf("ProcessUntilNoChange returned an error %v", err)
		}
		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}
