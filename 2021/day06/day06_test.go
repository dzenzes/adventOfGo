package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	input := "3,4,3,1,2"

	tests := []struct {
		days int
		want int
	}{
		{80, 5934},
		{256, 26984457539},
	}
	for _, test := range tests {
		t.Run("HowManyLanternfishAfterDays("+fmt.Sprint(test.days)+")", func(t *testing.T) {
			got := HowManyLanternfishAfterDays(input, test.days)
			assert.Equal(t, test.want, got)
		})
	}
}
