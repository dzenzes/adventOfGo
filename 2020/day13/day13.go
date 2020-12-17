package main

import (
	"fmt"
	"github.com/dmies/adventOfGo/filehandler"
	"log"
	"strconv"
	"strings"
)

type Bus struct {
	ID     int
	Offset int
}

type Departure struct {
	Bus       Bus
	Departure int
}

func Parse(input []string) (timestamp int, busses []Bus, err error) {
	firstLine := input[0]
	timestamp, err = strconv.Atoi(firstLine)
	if err != nil {
		return -1, []Bus{}, err
	}
	ids := input[1]

	for i, candidate := range strings.Split(ids, ",") {
		if candidate != "x" {
			id, err := strconv.Atoi(candidate)
			if err != nil {
				return -1, []Bus{}, err
			}
			busses = append(busses, Bus{id, i})
		}
	}
	return timestamp, busses, nil
}

func GetNextDeparture(bus Bus, currentTime int) Departure {
	time := bus.ID
	for time < currentTime {
		time += bus.ID
	}
	return Departure{bus, time}
}

func FindBestBus(departures []Departure, currentTime int) Departure {
	result := departures[0]
	for _, departure := range departures[1:] {
		if currentTime < departure.Departure {
			if departure.Departure < result.Departure {
				result = departure
			}
		}
	}
	return result
}

func GetMinutesToWait(departure int, currentTime int) int {
	return departure - currentTime
}

func GetResultForPart1(input []string) (int, error) {
	currentTime, busList, err := Parse(input)
	if err != nil {
		return -1, err
	}
	departures := make([]Departure, 0)
	for _, bus := range busList {
		departure := GetNextDeparture(bus, currentTime)
		departures = append(departures, departure)
	}
	bestBus := FindBestBus(departures, currentTime)

	minutesToWait := GetMinutesToWait(bestBus.Departure, currentTime)

	return bestBus.Bus.ID * minutesToWait, nil

}
func GetResultForPart2(input []string) (int, error) {
	_, busList, err := Parse(input)
	if err != nil {
		return -1, err
	}

	result := 0
	mod := 1
	for _, bus := range busList {
		for (result+bus.Offset)%bus.ID != 0 {
			result += mod
		}
		mod *= bus.ID
	}
	return result, nil

}

func main() {
	input, err := filehandler.ImportStringList("./input.txt")
	if err != nil {
		log.Fatal("Couldn't read input")
	}

	solution1, err := GetResultForPart1(input)
	if err != nil {
		log.Fatalf("GetResultForPart1 returned an error %v", err)
	}

	fmt.Printf("day 13, part1 %v\n", solution1)

	solution2, err := GetResultForPart2(input)
	if err != nil {
		log.Fatalf("GetResultForPart2 returned an error %v", err)
	}

	fmt.Printf("day 13, part2 %v\n", solution2)
}
