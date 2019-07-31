package main

import "fmt"

func main() {
	var i float32 = 42.2357
	fmt.Println("i = ", i)

	j := 43
	fmt.Println("j = ", j)

	sum := i + float32(j)
	fmt.Println("sum = ", sum)
}
