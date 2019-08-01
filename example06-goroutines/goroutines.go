package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(threadNumber string) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%s, count=%d\n", threadNumber, i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 1; i <= 3; i++ {
		go f(fmt.Sprintf("thread_number=%d", i))
	}
	var input string
	fmt.Scanln(&input)
}
