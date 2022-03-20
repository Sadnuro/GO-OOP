package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// Execution as normal function
	doSomething()

	// Esta ejecución no se evidenciará en consola
	// Ya que toma otro hilo y el actual finaliza
	// go goRoutine()

	c := make(chan int, 1) // Se crea el canal para ejecutar la go routine
	go goRoutine(c)        // Se pasa el canal como parametro al routine
	value := <-c           // Espera hasta que c obtenga un valor dentro de la routine

	fmt.Println("value obtained from goRoutine():", value)
}

func doSomething() {
	time.Sleep(time.Second * 3)
	fmt.Println("doSomething finished!")
}

func goRoutine(c chan<- int) {
	time.Sleep(time.Second * 3)
	fmt.Println("goRoutine finished!")

	// Asigna un valor al terminar la rutina.
	// El canal entiende que finalizó y vuelve al hilo principal
	c <- 1
}
