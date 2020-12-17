package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ParseMask(t *testing.T) {
	input := "mask = 10X0110X01100X00111XX00001X011101001"
	want := "10X0110X01100X00111XX00001X011101001"
	got := ParseMask(input)
	assert.Equal(t, got, want)
}

func Test_ParseMemory(t *testing.T) {
	input := "mem[45673] = 370803"
	wantAddress := 45673
	wantValue := 370803
	gotAddress, gotValue, err := ParseMemory(input)
	assert.Nil(t, err)
	assert.Equal(t, gotAddress, wantAddress)
	assert.Equal(t, gotValue, wantValue)
}

func Test_Parse(t *testing.T) {

	input := []string{

		"mask = 10X0110X01100X00111XX00001X011101001",
		"mem[45673] = 370803",
		"mem[32234] = 92667525",
		"mem[47600] = 955",
		"mem[6610] = 316949",
	}

	want := []Program{
		{"10X0110X01100X00111XX00001X011101001", 45673, 370803},
		{"10X0110X01100X00111XX00001X011101001", 32234, 92667525},
		{"10X0110X01100X00111XX00001X011101001", 47600, 955},
		{"10X0110X01100X00111XX00001X011101001", 6610, 316949},
	}

	got, err := Parse(input)
	assert.Nil(t, err)
	assert.EqualValues(t, got, want)
}

func Test_AddBitmasks(t *testing.T) {
	input := "101010"
	mask := "000000000000000000000000000000X1001X"
	want := "000000000000000000000000000000X1101X"

	got := AddBitmasks(input, mask)
	assert.Equal(t, got, want)
}

func Test_GetAddressesFromFloating(t *testing.T) {
	input := "000000000000000000000000000000X1101X"
	want := []string{
		"000000000000000000000000000000011010",
		"000000000000000000000000000000011011",
		"000000000000000000000000000000111010",
		"000000000000000000000000000000111011",
	}
	got := GetAddressesFromFloating(input)
	assert.EqualValues(t, got, want)
}

func Test_Program_calculate(t *testing.T) {
	tests := []struct {
		input Program
		want  int
	}{
		{Program{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 8, 11}, 73},
		{Program{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 7, 101}, 101},
	}

	for _, test := range tests {
		got, err := test.input.calculate()
		assert.Nil(t, err)
		assert.Equal(t, got, test.want)
	}

}

func Test_Program_calculateAddresses(t *testing.T) {
	tests := []struct {
		input Program
		want  []int
	}{
		{Program{"000000000000000000000000000000X1001X", 42, 100}, []int{26, 27, 58, 59}},
		{Program{"00000000000000000000000000000000X0XX", 26, 1}, []int{16, 17, 18, 19, 24, 25, 26, 27}},
	}

	for _, test := range tests {
		t.Run("calculateAddresses("+test.input.mask+")", func(t *testing.T) {
			got, err := test.input.calculateAddresses()
			assert.Nil(t, err)
			assert.EqualValues(t, got, test.want)
		})

	}

}

func Test_GetSumOfAllPrograms(t *testing.T) {
	input := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}
	want := 165

	got, err := GetSumOfAllPrograms(input)
	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func Test_GetSumOfAllProgramsV2(t *testing.T) {
	input := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}
	want := 208

	got, err := GetSumOfAllProgramsV2(input)
	assert.Nil(t, err)
	assert.Equal(t, got, want)
}
