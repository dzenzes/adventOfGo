package main

import "testing"
func Test_VisitedHouses(t *testing.T) {
	tests := []struct {
		input string
		want int 
	}{
		{">", 2},
		{">>", 3},
		{"^>v<",4},
		{"^v^v^v^v^v",2},
	}

	for _, test := range tests {
		t.Run("VisitedHouses("+test.input+")", func(t *testing.T) {
			input := test.input
			want := test.want
			got, err := VisitedHouses(input)
			if err != nil {
				t.Errorf("Parse() throws error %v", err)
			} else if got != want {
				t.Errorf("Parse() returned %v but %v was wanted", got, want)
			}
		})
	}

}


func Test_VisitedHousesWithRobot(t *testing.T) {
	tests := []struct {
		input string
		want int 
	}{
		{">", 2},
		{"^v", 3},
		{">>", 2},
		{"^>v<",3},
		{"^v^v^v^v^v",11},
	}

	for _, test := range tests {
		t.Run("VisitedHousesWithRobot("+test.input+")", func(t *testing.T) {
			input := test.input
			want := test.want
			got, err := VisitedHousesWithRobot(input)
			if err != nil {
				t.Errorf("Parse() throws error %v", err)
			} else if got != want {
				t.Errorf("Parse() returned %v but %v was wanted", got, want)
			}
		})
	}

}