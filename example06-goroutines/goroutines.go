package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f(threadNumber string, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%s, count=%d\n", threadNumber, i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go f(fmt.Sprintf("thread_number=%d", i), &wg)
	}
	wg.Wait()
	fmt.Println("finish")
}
