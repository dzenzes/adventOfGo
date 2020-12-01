package main

import (
	"fmt"

	"github.com/dmies/adventOfGo/filehandler"
)

func main() {

	program := filehandler.ImportNumberList("./input.txt")

	fmt.Printf("program %v", program)
}
