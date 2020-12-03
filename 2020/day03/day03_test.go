package main

import (
	"fmt"
	"testing"
)

func Test_TreeOnMap(t *testing.T) {
	// some scenarios to test
	scenarios := []struct {
		mapData   []string
		x         int
		y         int
		treeOnMap bool
	}{
		{
			[]string{"...#..."},
			0, 0, false,
		},
		{
			[]string{"...#..."},
			3, 0, true,
		},
		{
			[]string{"...#...", "..#..#.", "...#...", "..###.."},
			10, 3, true,
		},
		// test from puzzle
		{
			[]string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			},
			2, 5, true,
		},
		{
			[]string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			},
			222, 5, true,
		},
	}

	t.Run("TreeOnMap returns an error if y > len(map)", func(t *testing.T) {
		_, err := TreeOnMap([]string{"single line"}, 0, 5)
		if err == nil {
			t.Errorf("TreeOnMap() didn't throw an error for too big y")
		}
	})

	for i, scenario := range scenarios {
		t.Run("TreeOnMap", func(t *testing.T) {
			want := scenario.treeOnMap
			got, err := TreeOnMap(scenario.mapData, scenario.x, scenario.y)
			if err != nil {
				t.Errorf("TreeOnMap() error = %v", err)
			}
			if got != want {
				t.Errorf("TreeOnMap Scenario %v: got '%v' want '%v'", i, got, want)
			}
		})
	}
}

func Test_CountTreesOnMapForSlope(t *testing.T) {
	mapData := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	slopes := []struct {
		right         int
		down          int
		numberOfTrees int
	}{
		{1, 1, 2},
		{3, 1, 7},
		{5, 1, 3},
		{7, 1, 4},
		{1, 2, 2},
	}
	for i, slope := range slopes {
		t.Run("CountTreesOnMapForSlope scenario "+fmt.Sprint(i), func(t *testing.T) {
			want := slope.numberOfTrees
			got, err := CountTreesOnMapForSlope(mapData, slope.right, slope.down)
			if err != nil {
				t.Errorf("CountTreesOnMapForSlope() error = %v", err)
			}
			if got != want {
				t.Errorf("CountTreesOnMapForSlope Scenario %v: got '%v' want '%v'", i, got, want)
			}
		})
	}
}
