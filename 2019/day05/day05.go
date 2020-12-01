package main

import (
	"fmt"
	"log"

	"github.com/dmies/adventOfGo/filehandler"
)

func main() {

	program, err := filehandler.ImportNumberList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("program %v", program)
}
