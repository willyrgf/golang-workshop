package main

import "fmt"

const heightTallPerson = 1.90

func main() {
	firstName := "Joabe"
	lastName := "Leão"
	heightInMeters := 1.72
	workInAVCorp := true
	monthsWork := 30

	if len(firstName) < 1 || firstName == "" {
		panic("O primeiro nome esta vazio.")
	}

	if firstName == "Joabe" {
		fmt.Println("Ele é o Joabe ", lastName)
	} else {
		panic("Não conheço essa pessoa.")
	}

	if heightInMeters > heightTallPerson {
		fmt.Println("Ele é alto, mede ", heightInMeters)
	} else {
		fmt.Println("Ele não é tão alto, mede ", heightInMeters)
	}

	if !workInAVCorp {
		panic("Essa pessoa não trabalha na AVCorp.")
	}

	if monthsWork > 12 {
		fmt.Printf("Ele já trabalha há um tempo na AVCorp, mais precisamente %d meses.\n", monthsWork)
	} else {
		fmt.Println("Ele trabalha a pouco tempo na AVCorp.")
	}
}
