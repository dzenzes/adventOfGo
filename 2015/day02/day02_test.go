package main

import (
	"fmt"
	"testing"
)

func Test_Parse(t *testing.T) {
	tests := []struct {
		input string
		want  PackageDimensions
	}{
		{"13x27x17", PackageDimensions{13, 27, 17}},
		{"9x21x1", PackageDimensions{9, 21, 1}},
		{"29x25x6", PackageDimensions{29, 25, 6}},
	}

	for _, test := range tests {
		t.Run("Parse("+test.input+")", func(t *testing.T) {
			input := test.input
			want := test.want
			got, err := Parse(input)
			if err != nil {
				t.Errorf("Parse() throws error %v", err)
			} else if got != want {
				t.Errorf("Parse() returned %v but %v was wanted", got, want)
			}
		})
	}

	t.Run("Parse() throws an error if the input is invalid", func(t *testing.T) {
		input := "12+13+14"
		_, err := Parse(input)
		if err == nil {
			t.Errorf("Parse() should throw an error for invalid input")
		}
	})
}

func Test_PackageDimensions_getSurface(t *testing.T) {
	tests := []struct {
		input PackageDimensions
		want  int
	}{
		{PackageDimensions{2, 3, 4}, 58},
		{PackageDimensions{9, 21, 1}, 447},
		{PackageDimensions{29, 25, 6}, 2248},
		{PackageDimensions{1, 1, 10}, 43},
	}

	for _, test := range tests {
		t.Run("PackageDimensions("+fmt.Sprint(test.input)+").getSurface", func(t *testing.T) {
			input := test.input
			want := test.want
			got := input.getSurface()

			if got != want {
				t.Errorf("getSurface() returned %v but %v was wanted", got, want)
			}
		})
	}
}

func Test_PackageDimensions_getWrap(t *testing.T) {
	tests := []struct {
		input PackageDimensions
		want  int
	}{
		{PackageDimensions{2, 3, 4}, 10},
		{PackageDimensions{1, 1, 10}, 4},
	}

	for _, test := range tests {
		t.Run("PackageDimensions("+fmt.Sprint(test.input)+").getWrap", func(t *testing.T) {
			input := test.input
			want := test.want
			got := input.getWrap()

			if got != want {
				t.Errorf("getWrap() returned %v but %v was wanted", got, want)
			}
		})
	}
}

func Test_PackageDimensions_getBow(t *testing.T) {
	tests := []struct {
		input PackageDimensions
		want  int
	}{
		{PackageDimensions{2, 3, 4}, 24},
		{PackageDimensions{1, 1, 10}, 10},
	}

	for _, test := range tests {
		t.Run("PackageDimensions("+fmt.Sprint(test.input)+").getBow", func(t *testing.T) {
			input := test.input
			want := test.want
			got := input.getBow()

			if got != want {
				t.Errorf("getBow() returned %v but %v was wanted", got, want)
			}
		})
	}
}

func Test_GetTotalSquareFeetOfWrappingPaper(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"13x27x17", "9x21x1", "29x25x6"}, 4978},
	}

	for _, test := range tests {
		t.Run("GetTotalSquareFeetOfWrappingPaper("+fmt.Sprint(test.input)+")", func(t *testing.T) {
			input := test.input
			want := test.want
			got, err := GetTotalSquareFeetOfWrappingPaper(input)
			if err != nil {
				t.Errorf("GetTotalSquareFeetOfWrappingPaper() throws error %v", err)
			} else if got != want {
				t.Errorf("GetTotalSquareFeetOfWrappingPaper() returned %v but %v was wanted", got, want)
			}
		})
	}

}

func Test_GetTotalFeetOfRibbon(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"13x27x17", "9x21x1", "29x25x6"}, 10648},
	}

	for _, test := range tests {
		t.Run("GetTotalFeetOfRibbon("+fmt.Sprint(test.input)+")", func(t *testing.T) {
			input := test.input
			want := test.want
			got, err := GetTotalFeetOfRibbon(input)
			if err != nil {
				t.Errorf("GetTotalFeetOfRibbon() throws error %v", err)
			} else if got != want {
				t.Errorf("GetTotalFeetOfRibbon() returned %v but %v was wanted", got, want)
			}
		})
	}

}
