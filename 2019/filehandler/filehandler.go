package filehandler

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// ImportNumberList opens the given file and parses its contents (comma separated numbers) to an int array
func ImportNumberList(filename string) []int {
	program := make([]int, 0)

	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fileContentAsString := string(fileContent)
	stringList := strings.Split(strings.TrimSpace(fileContentAsString), ",")
	for _, numberFromList := range stringList {
		number, err := strconv.Atoi(numberFromList)
		if err != nil {
			log.Fatal(err)
		}
		program = append(program, number)
	}
	return program
}
