package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(threadNumber string) {
	for i := 0; i < 10; i++ {
		fmt.Println(threadNumber, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 3; i++ {
		go f(fmt.Sprintf("thread number: %d", i))
	}
	var input string
	fmt.Scanln(&input)
}
