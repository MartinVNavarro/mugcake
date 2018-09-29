package main

import (
	"fmt"
	"mugcake/components/ingredient"
	"os"
)

// main is a TEMPORARY main class and main package in order to
// quickly set up early stages of the project
func main() {
	mug, err := os.Open("examples/mug.json")

	if err != nil {
		os.Exit(1)
	}

	i, err := ingredient.NewFromReader(mug, "firstName")

	fmt.Println(i.Values)
}
