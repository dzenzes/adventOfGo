package intcomputer_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/dmies/adventOfGo/2019/intcomputer"
)

func TestIntComputer(t *testing.T) {
	tests := []struct {
		program []int
		result  []int
	}{
		{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
	}

	for _, test := range tests {
		t.Run(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(test.program)), ", "), "[]"), func(t *testing.T) {
			computer := intcomputer.Create(test.program)
			got := computer.Process().Memory
			want := test.result
			if !reflect.DeepEqual(got, want) {
				t.Errorf("intcomputer: got %v want %v", got, want)
			}

		})
	}
}
