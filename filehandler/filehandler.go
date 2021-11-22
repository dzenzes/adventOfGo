package filehandler

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// ImportNumberList opens the given file and parses its contents (comma separated numbers) to an int array
func ImportNumberList(file string) ([]int, error) {
	handle, err := os.Open(file)

	if err != nil {
		return nil, err
	}
	defer handle.Close()
	return ToNumberList(handle)
}

// ToNumberList takes a reader and transform the results of it to a []int
func ToNumberList(handle io.Reader) ([]int, error) {
	program := make([]int, 0)
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		stringList := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		for _, numberFromList := range stringList {
			number, err := strconv.Atoi(numberFromList)
			if err != nil {
				return nil, err
			}
			program = append(program, number)
		}
	}
	return program, nil
}

// ImportNumberPerLineList opens the given file and parses its contents (one number per line) to an []int 
func ImportNumberPerLineList(file string) ([]int, error) {
	handle, err := os.Open(file)

	if err != nil {
		return nil, err
	}
	defer handle.Close()
	return ToNumberPerLineList(handle)
}

// ToNumberPerLineList takes a reader and transform the results of it to a []int
func ToNumberPerLineList(handle io.Reader) ([]int, error) {
	program := make([]int, 0)
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {

		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		program = append(program, number)

	}
	return program, nil
}

// ImportStringList opens the given file and parses its contents (one String per line) to an String array
func ImportStringList(file string) ([]string, error) {
	handle, err := os.Open(file)

	if err != nil {
		return nil, err
	}
	defer handle.Close()
	return ToSringList(handle)
}

// ToSringList takes a reader and transform the results of it to a []string
func ToSringList(handle io.Reader) ([]string, error) {
	program := make([]string, 0)
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		text := scanner.Text()
		program = append(program, text)
	}
	return program, nil
}
