package main

import (
	"errors"
	"fmt"
)

func isEnable(enable bool) (bool, error) {
	if enable {
		return false, errors.New("Erro ao habilitar.")
	}

	return true, nil
}

func isDisable(disable bool) (bool, error) {
	if disable {
		return false, fmt.Errorf("Erro ao desabilitar.")
	}

	return true, nil
}

func main() {
	b, err := isEnable(true)
	if err != nil {
		fmt.Println("Error isEnable(true):", err)
	}

	b, err = isDisable(true)
	if err != nil {
		fmt.Println("Error isDisabled(true):", err)
	}

	b, err = isDisable(false)
	if err != nil {
		fmt.Println("Error isDisabled(false):", err)
	} else {
		fmt.Printf("Retorno isDisable(false): b=%v, err=%v\n", b, err)
	}
}
