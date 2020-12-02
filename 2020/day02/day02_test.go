package main

import "testing"

var tests = []struct {
	policyAndPassword string
	min               int
	max               int
	letter            string
	password          string
	valid             bool
	validInNewShop    bool
}{
	{"15-19 k: kkkkkkkkkkkkzkkkkkkk", 15, 19, "k", "kkkkkkkkkkkkzkkkkkkk", true, false},
	{"1-11 s: sbssswsqsssssrlss", 1, 11, "s", "sbssswsqsssssrlss", false, false},
	{"8-9 b: pbbbbbbkbz", 8, 9, "b", "pbbbbbbkbz", false, true},
	{"4-10 w: wwccwcqwdmbktjrxhw", 4, 10, "w", "wwccwcqwdmbktjrxhw", true, false},
	{"1-6 x: jvscgqsnt", 1, 6, "x", "jvscgqsnt", false, false},
	{"1-7 x: xxxxxxcx", 1, 7, "x", "xxxxxxcx", true, true},
	{"1-3 a: abcde", 1, 3, "a", "abcde", true, true},
	{"1-3 b: cdefg", 1, 3, "b", "cdefg", false, false},
	{"2-9 c: ccccccccc", 2, 9, "c", "ccccccccc", true, false},
}

func Test_Parse(t *testing.T) {
	for _, test := range tests {
		t.Run("Parse: "+test.policyAndPassword, func(t *testing.T) {
			gotMin, gotMax, gotLetter, gotPassword, err := Parse(test.policyAndPassword)
			if err != nil {
				t.Errorf("Parse() error = %v", err)
			}
			if gotMin != test.min {
				t.Errorf("Parse: got '%v' want '%v'", gotMin, test.min)
			}
			if gotMax != test.max {
				t.Errorf("Parse: got '%v' want '%v'", gotMax, test.max)
			}
			if gotLetter != test.letter {
				t.Errorf("Parse: got '%v' want '%v'", gotLetter, test.letter)
			}
			if gotPassword != test.password {
				t.Errorf("Parse: got '%v' want '%v'", gotPassword, test.password)
			}
		})
	}
}
func Test_ValidatePassword(t *testing.T) {

	for _, test := range tests {
		t.Run("ValidatePassword: "+test.policyAndPassword, func(t *testing.T) {
			got := ValidatePassword(test.min, test.max, test.letter, test.password)

			if got != test.valid {
				t.Errorf("ValidatePassword: got %v want %v", got, test.valid)
			}

		})
	}
}

func Test_ValidatePasswordForNewShop(t *testing.T) {

	for _, test := range tests {
		t.Run("ValidatePasswordForNewShop: "+test.policyAndPassword, func(t *testing.T) {
			got := ValidatePasswordForNewShop(test.min, test.max, test.letter, test.password)

			if got != test.validInNewShop {
				t.Errorf("ValidatePasswordForNewShop: got %v want %v", got, test.validInNewShop)
			}

		})
	}
}
