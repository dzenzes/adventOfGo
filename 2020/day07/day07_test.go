package main

import (
	"fmt"
	"reflect"
	"testing"
)

var rules = map[string][]BagAndCount{
	"bright white": {{"shiny gold", 1}},
	"dark olive":   {{"faded blue", 3}, {"dotted black", 4}},
	"dark orange":  {{"bright white", 3}, {"muted yellow", 4}},
	"dotted black": {},
	"faded blue":   {},
	"light red":    {{"bright white", 1}, {"muted yellow", 2}},
	"muted yellow": {{"shiny gold", 2}, {"faded blue", 9}},
	"shiny gold":   {{"dark olive", 1}, {"vibrant plum", 2}},
	"vibrant plum": {{"faded blue", 5}, {"dotted black", 6}},
}

func Test_ParseAll(t *testing.T) {
	tests := []struct {
		rules []string
		want  map[string][]BagAndCount
	}{
		{
			[]string{
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
				"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
				"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
				"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
				"faded blue bags contain no other bags.",
				"dotted black bags contain no other bags.",
			},
			rules,
		},
	}

	for i, test := range tests {
		t.Run("ParseAll() "+fmt.Sprint(i), func(t *testing.T) {
			input := test.rules
			want := test.want
			got := ParseAll(input)

			for key, value := range got {
				if wantedValue, ok := want[key]; ok {
					if len(wantedValue) > 0 && len(wantedValue) == len(value) && !reflect.DeepEqual(value, wantedValue) {
						t.Errorf("ParseAll() for key %v returned %v but wanted %v", key, value, wantedValue)
					}
				} else {
					t.Errorf("ParseAll() missing key %v", key)
				}
			}

		})
	}
}

func Test_FindBag(t *testing.T) {
	tests := []struct {
		searched string
		key      string
		want     bool
	}{
		{"shiny gold", "bright white", true},
		{"shiny gold", "light red", true},
		{"dark orange", "bright white", false},
	}
	for _, test := range tests {
		t.Run("FindBag() searched: "+test.searched+" key:"+test.key, func(t *testing.T) {
			want := test.want
			got := FindBag(test.searched, rules, test.key)

			if got != want {
				t.Errorf("FindBag() returned %v but wanted %v", got, want)
			}
		})
	}
}

func Test_CountBagsThatContainColor(t *testing.T) {

	tests := []struct {
		searched string

		want int
	}{
		{"shiny gold", 4},
		{"light red", 0},
		{"dotted black", 7},
	}

	for _, test := range tests {
		t.Run("CountBagsThatContainColor() searched: "+test.searched, func(t *testing.T) {
			want := test.want
			got, err := CountBagsThatContainColor(test.searched, rules)
			if err != nil {
				t.Errorf("CountBagsThatContainColor() threw an error %v", err)
			}
			if got != want {
				t.Errorf("CountBagsThatContainColor() returned %v but wanted %v", got, want)
			}
		})
	}

}

func Test_GetNumberOfContainedBags(t *testing.T) {
	tests := []struct {
		color         string
		containedBags int
	}{
		{"faded blue", 0},
		{"dotted black", 0},
		{"vibrant plum", 11},
		{"dark olive", 7},
	}

	for _, test := range tests {
		t.Run("GetNumberOfContainedBags("+test.color+") => "+fmt.Sprint(test.containedBags), func(t *testing.T) {
			input := test.color
			want := test.containedBags
			got := GetNumberOfContainedBags(input, rules)

			if got != want {
				t.Errorf("GetNumberOfContainedBags() returned %v but wanted %v", got, want)
			}
		})
	}
}
