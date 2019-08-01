package main

import (
	"fmt"
	"os"
)

func helloWorld(name string) string {
	hello := "Hello, " + name
	return hello
}

func main() {
	msg := helloWorld(`Will`)
	fmt.Println(msg)

	argName := os.Args[1:]
	if len(argName) >= 1 {
		fmt.Println(helloWorld(argName[0]))
	}
}
