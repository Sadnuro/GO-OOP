package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Limita la cantidad de rutinas ejecutadas al tiempo
	c := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		/*
			Agrega una rutina al canal
			junto a la cantidad de canal del buffer genera un bloqueo
			para liminar la cantida de routinas a ejecutar
			c := [][]
			c := [goRoutine_1] []
			c := [goRoutine_1] [goRoutine_2]
			c := [] [goRoutine_2]
			c := [goRoutine_3] [goRoutine_2]
			...
		*/
		c <- 1 // Ocupa un canal disponible
		wg.Add(1)
		go doSomething2(i, &wg, c)
	}

	wg.Wait()
}

func doSomething2(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()

	fmt.Printf("Id %d started\n", i)
	// fmt.Printf("channel", c)
	time.Sleep(time.Second * 5)
	fmt.Printf("Id %d finished\n", i)
	<-c // decosupa el canal usado
}
