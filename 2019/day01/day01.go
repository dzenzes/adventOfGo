package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) []int {
	lines := make([]int, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

// Fuel required to launch a given module is based on its mass.
// Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.
func Fuel(mass int) int {
	return mass/3 - 2
}

// TotalFuel calculates the needed fuel plus the fuel needed for that additional fuel (mass)
func TotalFuel(mass int) int {
	total := 0
	for fuel := Fuel(mass); fuel > 0; fuel = Fuel(fuel) {
		total += fuel
	}

	return total
}

func part1(nums []int) int {
	total := 0
	for _, num := range nums {
		total += Fuel(num)
	}
	return total
}

func part2(nums []int) int {
	total := 0
	for _, num := range nums {
		total += TotalFuel(num)
	}
	return total
}

func main() {
	nums := parse("./input.txt")

	fmt.Printf("Part1: %d \n", part1(nums))
	fmt.Printf("Part2: %d \n", part2(nums))
}
