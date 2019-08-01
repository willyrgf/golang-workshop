package main

import (
	"fmt"
	"time"
)

func main() {
	SliceOfNames := []string{"Joabe", "Phelipe", "Gabriel", "Igor", "Matheus", "Lucas", "Gabriel"}
	for i := 0; i < len(SliceOfNames); i++ {
		fmt.Println("1: Nome: ", SliceOfNames[i])
	}

	fmt.Println()
	for _, n := range SliceOfNames {
		fmt.Println("2: Nome: ", n)
	}

	fmt.Println()
	for i, n := range SliceOfNames {
		fmt.Printf("3: Posicao:%d Nome:%s\n", i, n)
	}

	fmt.Println()
	for {
		fmt.Println("Isso ta em loop?")
		time.Sleep(100 * time.Millisecond)
	}
}
