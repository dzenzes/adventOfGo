package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_ExtractSingleLinePassport(t *testing.T) {

	tests := []struct {
		input []string
		want  []string
	}{
		{[]string{
			"abc", "def", "", "ghi", "", "jkl",
		}, []string{
			"abc def", "ghi", "jkl",
		}},
		{[]string{

			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
			"byr:1937 iyr:2017 cid:147 hgt:183cm",
			"",
			"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
			"hcl:#cfa07d byr:1929",
			"",
			"hcl:#ae17e1 iyr:2013",
			"eyr:2024",
			"ecl:brn pid:760753108 byr:1931",
			"hgt:179cm",
			"",
			"hcl:#cfa07d eyr:2025 pid:166559648",
			"iyr:2011 ecl:brn hgt:59in",
		}, []string{
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
			"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
			"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
			"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in",
		}},
	}

	for i, test := range tests {
		t.Run("ExtractSingleLinePassport "+fmt.Sprint(i), func(t *testing.T) {
			input := test.input
			want := test.want
			got, err := ExtractSingleLinePassport(input)
			if err != nil {
				t.Errorf("ExtractSingleLinePassport() error = %v", err)
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("ExtractSingleLinePassport() got %v but wanted %v", got, want)
			}
		})
	}

}

func Test_GetKeysFromPassport(t *testing.T) {
	tests := []struct {
		passport string
		keys     []string
	}{
		{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "cid", "hgt"}},
		{"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929", []string{"iyr", "ecl", "cid", "eyr", "pid", "hcl", "byr"}},
		{"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm", []string{"hcl", "iyr", "eyr", "ecl", "pid", "byr", "hgt"}},
		{"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in", []string{"hcl", "eyr", "pid", "iyr", "ecl", "hgt"}},
	}

	for i, test := range tests {
		t.Run("GetKeysFromPassport "+fmt.Sprint(i), func(t *testing.T) {
			input := test.passport
			want := test.keys
			got, err := GetKeysFromPassport(input)
			if err != nil {
				t.Errorf("GetKeysFromPassport() error = %v", err)
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetKeysFromPassport() got %v but wanted %v", got, want)
			}
		})
	}

}

func Test_CheckKeysForCompleteness(t *testing.T) {

	tests := []struct {
		keys     []string
		complete bool
	}{
		{[]string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "cid", "hgt"}, true},
		{[]string{"iyr", "ecl", "cid", "eyr", "pid", "hcl", "byr"}, false},
		{[]string{"hcl", "iyr", "eyr", "ecl", "pid", "byr", "hgt"}, true},
		{[]string{"hcl", "eyr", "pid", "iyr", "ecl", "hgt"}, false},
	}

	for i, test := range tests {
		t.Run("CheckKeysForCompleteness "+fmt.Sprint(i), func(t *testing.T) {
			input := test.keys
			want := test.complete
			got := CheckKeysForCompleteness(input)

			if got != want {
				t.Errorf("CheckKeysForCompleteness() got %v but wanted %v", got, want)
			}
		})
	}

}

func Test_GetKeyValuesFromPassport(t *testing.T) {
	tests := []struct {
		passport  string
		keyValues []KeyValue
	}{
		{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
			[]KeyValue{{"ecl", "gry"}, {"pid", "860033327"}, {"eyr", "2020"}, {"hcl", "#fffffd"}, {"byr", "1937"}, {"iyr", "2017"}, {"cid", "147"}, {"hgt", "183cm"}}},
		{"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
			[]KeyValue{{"iyr", "2013"}, {"ecl", "amb"}, {"cid", "350"}, {"eyr", "2023"}, {"pid", "028048884"}, {"hcl", "#cfa07d"}, {"byr", "1929"}}},
		{"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
			[]KeyValue{{"hcl", "#ae17e1"}, {"iyr", "2013"}, {"eyr", "2024"}, {"ecl", "brn"}, {"pid", "760753108"}, {"byr", "1931"}, {"hgt", "179cm"}}},
		{"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in",
			[]KeyValue{{"hcl", "#cfa07d"}, {"eyr", "2025"}, {"pid", "166559648"}, {"iyr", "2011"}, {"ecl", "brn"}, {"hgt", "59in"}}},
	}

	for i, test := range tests {
		t.Run("GetKeyValuesFromPassport "+fmt.Sprint(i), func(t *testing.T) {
			input := test.passport
			want := test.keyValues
			got, err := GetKeyValuesFromPassport(input)
			if err != nil {
				t.Errorf("GetKeyValuesFromPassport() error = %v", err)
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetKeyValuesFromPassport() got %v but wanted %v", got, want)
			}
		})
	}

}

func Test_ValidateKeyValue(t *testing.T) {
	tests := []struct {
		comment string
		input   KeyValue
		valid   bool
	}{
		// byr
		{"max case", KeyValue{"byr", "2002"}, true},
		{"lowest supported year", KeyValue{"byr", "1920"}, true},
		{"normal case", KeyValue{"byr", "1981"}, true},
		{"year too big", KeyValue{"byr", "2003"}, false},
		{"year too small", KeyValue{"byr", "1910"}, false},
		// iyr
		{"normal case", KeyValue{"iyr", "2012"}, true},
		{"lowest supported year", KeyValue{"iyr", "2010"}, true},
		{"max supported year", KeyValue{"iyr", "2020"}, true},
		{"year too big", KeyValue{"iyr", "2021"}, false},
		{"year too small", KeyValue{"iyr", "2009"}, false},
		{"year too small", KeyValue{"iyr", "1910"}, false},
		// eyr
		{"normal case", KeyValue{"eyr", "2022"}, true},
		{"year too big", KeyValue{"eyr", "2031"}, false},
		{"lowest supported year", KeyValue{"eyr", "2020"}, true},
		{"max supported year", KeyValue{"eyr", "2030"}, true},
		{"year too small", KeyValue{"eyr", "1910"}, false},
		// hgt
		{"min valid height (cm)", KeyValue{"hgt", "150cm"}, true},
		{"max valid height(cm)", KeyValue{"hgt", "193cm"}, true},
		{"normal height (cm)", KeyValue{"hgt", "175cm"}, true},
		{"too small (cm)", KeyValue{"hgt", "149"}, false},
		{"too big (cm)", KeyValue{"hgt", "194cm"}, false},
		{"valid height (in)", KeyValue{"hgt", "60in"}, true},
		{"lowest height (in)", KeyValue{"hgt", "59in"}, true},
		{"max height (in)", KeyValue{"hgt", "76in"}, true},
		{"height too big (in)", KeyValue{"hgt", "190in"}, false},
		{"height too small (in)", KeyValue{"hgt", "58in"}, false},
		{"height without unit", KeyValue{"hgt", "190"}, false},
		// hcl
		{"valid case", KeyValue{"hcl", "#123abc"}, true},
		{"valid case", KeyValue{"hcl", "#ffffff"}, true},
		{"valid case", KeyValue{"hcl", "#012345"}, true},
		{"valid case", KeyValue{"hcl", "#6789ab"}, true},
		{"valid case", KeyValue{"hcl", "#cdef12"}, true},
		{"illegal hex code", KeyValue{"hcl", "#123abz"}, false},
		{"code without #", KeyValue{"hcl", "123abc"}, false},
		{"code too long", KeyValue{"hcl", "#123abcx"}, false},
		// ecl
		{"valid", KeyValue{"ecl", "amb"}, true},
		{"valid", KeyValue{"ecl", "blu"}, true},
		{"valid", KeyValue{"ecl", "brn"}, true},
		{"valid", KeyValue{"ecl", "gry"}, true},
		{"valid", KeyValue{"ecl", "grn"}, true},
		{"valid", KeyValue{"ecl", "hzl"}, true},
		{"valid", KeyValue{"ecl", "oth"}, true},
		{"invalid", KeyValue{"ecl", "wat"}, false},
		// pid
		{"valid", KeyValue{"pid", "000000001"}, true},
		{"too long", KeyValue{"pid", "0123456789"}, false},
		{"only digits allowed", KeyValue{"pid", "0000o0001"}, false},
		// cid
		{"can be ignored", KeyValue{"cid", "whatever"}, true},
	}

	for _, test := range tests {
		testName := "ValidateKeyValue() for key: " + test.input.key + " with value: " + test.input.value + " (" + test.comment + ")"
		t.Run(testName, func(t *testing.T) {
			input := test.input
			want := test.valid

			got, err := ValidateKeyValue(input)
			if err != nil {
				t.Errorf("ValidateKeyValue() error = %v", err)
			}
			if got != want {
				t.Errorf("ValidateKeyValue() got %v but wanted %v", got, want)
			}
		})
	}
}

func Test_ValidateKeyValueThrowsErrors(t *testing.T) {
	tests := []struct {
		keyValue KeyValue
		reason   string
	}{
		{KeyValue{"byr", "abc"}, "value should be number"},
		{KeyValue{"def", "2003"}, "unsupported key"},
	}

	for i, test := range tests {
		t.Run("ValidateKeyValue "+fmt.Sprint(i), func(t *testing.T) {
			input := test.keyValue

			_, err := ValidateKeyValue(input)
			if err == nil {
				t.Errorf("ValidateKeyValue() should have thrown error. %v", test.reason)
			}

		})
	}
}

func Test_Part1(t *testing.T) {
	input := []string{

		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	}
	want := 2
	got, err := Part1(input)
	if err != nil {
		t.Errorf("Part1() error = %v", err)
	}

	if got != want {
		t.Errorf("Part1() got %v but wanted %v", got, want)
	}
}

func Test_Part2(t *testing.T) {
	tests := []struct {
		comment                string
		passports              []string
		numberOfValidPassports int
	}{
		{"all invalid",
			[]string{
				"eyr:1972 cid:100",
				"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
				"",
				"iyr:2019",
				"hcl:#602927 eyr:1967 hgt:170cm",
				"ecl:grn pid:012533040 byr:1946",
				"",
				"hcl:dab227 iyr:2012",
				"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
				"",
				"hgt:59cm ecl:zzz",
				"eyr:2038 hcl:74454a iyr:2023",
				"pid:3556412378 byr:2007",
			}, 0},
		{"all valid",
			[]string{
				"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
				"hcl:#623a2f",
				"",
				"eyr:2029 ecl:blu cid:129 byr:1989",
				"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
				"",
				"hcl:#888785",
				"hgt:164cm byr:2001 iyr:2015 cid:88",
				"pid:545766238 ecl:hzl",
				"eyr:2022",
				"",
				"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
			}, 4,
		},
	}

	for _, test := range tests {
		testName := "Part2() for key: " + test.comment
		t.Run(testName, func(t *testing.T) {
			input := test.passports
			want := test.numberOfValidPassports

			got, err := Part2(input)
			if err != nil {
				t.Errorf("Part2() error = %v", err)
			}
			if got != want {
				t.Errorf("Part2() got %v but wanted %v", got, want)
			}
		})
	}
}
