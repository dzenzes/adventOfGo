package main

import "testing"

// get next number
// build input from it
// calculate hash
// check if hash starts with 5 leading zeros

func Test_AdventCoins(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, test := range tests {
		t.Run("AdventCoins("+test.input+")", func(t *testing.T) {
			input := test.input
			want := test.want
			got := AdventCoins(input, "00000")
			if got != want {
				t.Errorf("AdventCoins() returned %v but %v was wanted", got, want)
			}
		})
	}

}
