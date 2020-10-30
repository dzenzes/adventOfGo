package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestIntComputer(t *testing.T) {
	tests := []struct {
		program []int
		result  []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
	}

	for _, test := range tests {
		t.Run(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(test.program)), ", "), "[]"), func(t *testing.T) {
			got := IntComputer(test.program)
			want := test.result
			if !reflect.DeepEqual(got, want) {
				t.Errorf("intcomputer: got %v want %v", got, want)
			}

		})
	}
}
