package main

import (
	"fmt"

	"github.com/dmies/adventOfGo/2019/filehandler"
)

func main() {

	program := filehandler.ImportNumberList("./input.txt")

	fmt.Printf("program %v", program)
}
