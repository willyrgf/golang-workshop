package main

import (
	"fmt"
)

func main() {
	ArrayOfNames := [6]string{"Joabe", "Phelipe", "Gabriel", "Igor", "Matheus"}
	fmt.Printf("1: ArrayOfNames: %v\n", ArrayOfNames)
	fmt.Printf("1: len(ArrayOfNames: %v\n", len(ArrayOfNames))

	fmt.Println()
	ArrayOfNames[0] = "Phelipe"
	ArrayOfNames[1] = "Joabe"
	fmt.Printf("2: ArrayOfNames: %v\n", ArrayOfNames)
	fmt.Printf("2: len(ArrayOfNames: %v\n", len(ArrayOfNames))

	fmt.Println()
	SliceOfNames := []string{"Joabe", "Phelipe", "Gabriel", "Igor", "Matheus"}
	fmt.Printf("3: SliceOfNames: %v\n", SliceOfNames)
	fmt.Printf("3: len(SliceOfNames: %v\n", len(SliceOfNames))

	fmt.Println()
	SliceOfNames = append(SliceOfNames, "Lucas")
	SliceOfNames = append(SliceOfNames, "Alex")
	fmt.Printf("4: SliceOfNames: %v\n", SliceOfNames)
	fmt.Printf("4: len(SliceOfNames: %v\n", len(SliceOfNames))
}
