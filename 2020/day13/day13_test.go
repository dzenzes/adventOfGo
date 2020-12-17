package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Parse(t *testing.T) {
	input := []string{
		"123",
		"1,x,2,3,x,x,x,4,5,6",
	}
	wantedTimestamp := 123
	wantedBusList := []Bus{
		Bus{1, 0},
		Bus{2, 2},
		Bus{3, 3},
		Bus{4, 7},
		Bus{5, 8},
		Bus{6, 9},
	}

	gotTime, gotBusList, err := Parse(input)
	if err != nil {
		t.Errorf("got an error %v", err)
	}
	if gotTime != wantedTimestamp {
		t.Errorf("got %v but wanted %v", gotTime, wantedTimestamp)
	}
	if !reflect.DeepEqual(gotBusList, wantedBusList) {
		t.Errorf("got %v but wanted %v", gotBusList, wantedBusList)
	}
}

func Test_GetNextDeparture(t *testing.T) {

	tests := []struct {
		currentTime int
		bus         Bus
		want        Departure
	}{
		{939, Bus{7, 0}, Departure{Bus{7, 0}, 945}},
		{939, Bus{13, 0}, Departure{Bus{13, 0}, 949}},
		{939, Bus{59, 0}, Departure{Bus{59, 0}, 944}},
		{939, Bus{31, 0}, Departure{Bus{31, 0}, 961}},
		{939, Bus{19, 0}, Departure{Bus{19, 0}, 950}},
	}

	for _, test := range tests {
		t.Run("GetNextDeparture("+fmt.Sprint(test.currentTime)+", "+fmt.Sprint(test.bus)+")", func(t *testing.T) {
			got := GetNextDeparture(test.bus, test.currentTime)

			if got != test.want {
				t.Errorf("got %v but wanted %v", got, test.want)
			}
		})
	}

}

func Test_FindBestBus(t *testing.T) {
	time := 939
	departures := []Departure{
		Departure{Bus{7, 0}, 945},
		Departure{Bus{13, 0}, 949},
		Departure{Bus{59, 0}, 944},
		Departure{Bus{31, 0}, 961},
		Departure{Bus{19, 0}, 950},
	}
	want := Departure{Bus{59, 0}, 944}
	got := FindBestBus(departures, time)

	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

func Test_GetMinutesToWait(t *testing.T) {
	input := 944
	currentTime := 939
	want := 5
	got := GetMinutesToWait(input, currentTime)
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}

}

func Test_GetResultForPart1(t *testing.T) {
	input := []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}
	want := 295

	got, err := GetResultForPart1(input)
	if err != nil {
		t.Errorf("got an error %v", err)
	}
	if want != got {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

func Test_GetResultForPart2(t *testing.T) {
	input := []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}
	want := 1068781

	got, err := GetResultForPart2(input)
	if err != nil {
		t.Errorf("got an error %v", err)
	}
	if want != got {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
