package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dmies/adventOfGo/filehandler"
)

type KeyValue struct {
	key   string
	value string
}

// PassportKeys are the needed keys in a passport
var PassportKeys = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	//"cid",
}

var EyeColors = []string{
	"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
}

// ExtractSingleLinePassport takes the splitted passports and joins them
func ExtractSingleLinePassport(input []string) ([]string, error) {
	result := make([]string, 0)
	passport := ""

	for _, line := range input {
		if len(line) == 0 {
			result = append(result, strings.Trim(passport, " "))
			passport = ""
		} else {
			passport = passport + " " + line
		}
	}
	result = append(result, strings.Trim(passport, " "))

	return result, nil
}

// GetKeysFromPassport takes a passport (string) and returns the keys in that passport
func GetKeysFromPassport(passport string) ([]string, error) {
	result := make([]string, 0)
	keyValues := strings.Split(passport, " ")
	for _, keyValuePair := range keyValues {
		key := strings.Split(keyValuePair, ":")[0]
		result = append(result, key)
	}

	return result, nil
}

// GetKeyValuesFromPassport takes a passport (string) and returns the key value pairs in that passport
func GetKeyValuesFromPassport(passport string) ([]KeyValue, error) {
	result := make([]KeyValue, 0)
	keyValues := strings.Split(passport, " ")
	for _, keyValuePair := range keyValues {
		keyValue := strings.Split(keyValuePair, ":")
		key := keyValue[0]
		value := keyValue[1]
		result = append(result, KeyValue{key, value})
	}

	return result, nil
}

func checkYear(value string, min int, max int) (bool, error) {
	yearToCheck, err := strconv.Atoi(value)
	if err != nil {
		return false, fmt.Errorf("checkYear: couldn't cast year to int (%v)", value)
	}
	return yearToCheck >= min && yearToCheck <= max, nil
}

func validateHeight(value string, unit string, min int, max int) (bool, error) {
	valueAsString := strings.Split(value, unit)[0]
	valueAsNumber, err := strconv.Atoi(valueAsString)
	if err != nil {
		return false, err
	}
	return valueAsNumber >= min && valueAsNumber <= max, nil
}

// ValidateKeyValue checks if the given key value pair from the passport is valid
func ValidateKeyValue(keyValue KeyValue) (bool, error) {

	switch keyValue.key {
	case "byr":
		return checkYear(keyValue.value, 1920, 2002)
	case "iyr":
		return checkYear(keyValue.value, 2010, 2020)
	case "eyr":
		return checkYear(keyValue.value, 2020, 2030)
	case "hgt":
		if strings.Contains(keyValue.value, "cm") {
			result, err := validateHeight(keyValue.value, "cm", 150, 193)

			if err != nil {
				return false, err
			}
			return result, nil
		}
		if strings.Contains(keyValue.value, "in") {
			result, err := validateHeight(keyValue.value, "in", 59, 76)
			if err != nil {
				return false, err
			}
			return result, nil
		}
		return false, nil
	case "hcl":
		if len(keyValue.value) == 7 && strings.Index(keyValue.value, "#") == 0 {
			validChars := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
			code := keyValue.value[1:]
			for _, char := range code {
				if !contains(validChars, string(char)) {
					return false, nil
				}
			}
			return true, nil
		}
		return false, nil
	case "ecl":
		return contains(EyeColors, keyValue.value), nil
	case "pid":
		_, err := strconv.Atoi(keyValue.value)
		if err != nil {
			return false, nil
		}
		return len(keyValue.value) == 9, nil
	case "cid":
		return true, nil
	default:
		return false, fmt.Errorf("ValidateKeyValue: couldn't handle key (%v)", keyValue.key)
	}

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// CheckKeysForCompleteness checks that are all required keys are present
func CheckKeysForCompleteness(keys []string) bool {
	for _, key := range PassportKeys {
		if !contains(keys, key) {
			return false
		}
	}
	return true
}

// Part1 counts the valid passports
func Part1(passports []string) (int, error) {
	result := 0
	passportList, err := ExtractSingleLinePassport(passports)
	if err != nil {
		return -1, err
	}

	for _, passport := range passportList {
		keyList, err := GetKeysFromPassport(passport)
		if err != nil {
			return -1, err
		}

		if CheckKeysForCompleteness(keyList) {
			result++
		}
	}
	return result, nil
}

// Part2 counts the valid passports
func Part2(passports []string) (int, error) {
	result := 0
	passportList, err := ExtractSingleLinePassport(passports)
	if err != nil {
		return -1, err
	}

	for _, passport := range passportList {
		keyValueList, err := GetKeyValuesFromPassport(passport)
		if err != nil {
			return -1, err
		}
		valid := true

		keyList := make([]string, 0)
		for _, keyValue := range keyValueList {
			keyList = append(keyList, keyValue.key)
		}

		if CheckKeysForCompleteness(keyList) {
			for _, keyValue := range keyValueList {

				keyValueOkay, err := ValidateKeyValue(keyValue)
				if err != nil {
					return -1, err
				}

				valid = valid && keyValueOkay

			}
			if valid {
				result++
			}

		}

	}
	return result, nil
}

func main() {
	passports, err := filehandler.ImportSringList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution1, err := Part1(passports)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day 04, part1 %v\n", solution1)

	solution2, err := Part2(passports)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day 04, part2 %v\n", solution2)

}
