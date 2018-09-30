package main

import (
	"fmt"
	"mugcake/components/ingredient"
	"os"
)

// main is a TEMPORARY main class and main package in order to
// quickly set up early stages of the project
func main() {
	// mug, err := os.Open("examples/mug.json")

	// if err != nil {
	// 	os.Exit(1)
	// }

	// i, err := ingredient.ListFromReader(mug)

	// fmt.Println(i.Ingredients[0])
	// i := ingredient.NewList()

	ing1 := ingredient.Ingredient{Name: "Name", Type: "String", Size: 50, Precision: 0}
	ing1.Add("Bob")
	ing1.Add("Billy")
	ing1.Add("Jones")
	ing1.Add("Jimmy")

	ing2 := ingredient.Ingredient{Name: "Address", Type: "String", Size: 100}

	ing2.Add("1234 EHR Way")
	ing2.Add("777 Jesus Ave")
	ing2.Add("456 Hihi St")

	il := ingredient.NewList()

	il.Add(ing1)
	il.Add(ing2)

	fmt.Println(il)

	r, err := os.Create("new.json")

	if err != nil {
		os.Exit(1)
	}

	il.WriteToFile(r)
}
