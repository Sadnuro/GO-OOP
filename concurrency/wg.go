package main

import (
	"fmt"
	"sync"
	"time"
)

func main2() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // Add unit to waitgroup counter
		go doSomething1(i, &wg)
	}

	wg.Wait() // Espera hasta que el contador sea 0
}

func doSomething1(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrease by 1 waitgroup counter

	fmt.Println("Started:", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
}
