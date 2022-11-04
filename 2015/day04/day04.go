package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
)

func isValidHash(input string, pattern string) bool {
	regexPattern := fmt.Sprintf("%s%s", "\\A", pattern)
	regex := regexp.MustCompile(regexPattern)
	return regex.MatchString(input)
}

func getMD5Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func AdventCoins(input string, pattern string) int {
	number := 0
	result := ""
	for !isValidHash(result, pattern) {
		number++
		hash := fmt.Sprintf("%s%d", input, number)
		result = getMD5Hash(hash)
	}
	return number
}

func main() {

	solution1 := AdventCoins("iwrupvqb", "00000")

	fmt.Printf("day 04, part1 %v\n", solution1)

	solution2 := AdventCoins("iwrupvqb", "000000")

	fmt.Printf("day 04, part2 %v\n", solution2)

}
