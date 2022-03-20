package main

import (
	"fmt"
	"time"
)

func main() {
	// NONONYMOUS FUNCTION : Asegurarse de no romper DRY

	// Funcion anonima definida
	anon := func(a float64) float64 {
		return a * 2
	} // (): Ejecuta e inserta valor de parametros

	fmt.Println("Func ref:", anon)           // Referencia a la funcion
	fmt.Println("Func ref exec:", anon(2.5)) // Ejecuta la funcion referenciada

	// Funcion anonima definida y  ejecutada
	x := 5
	y := func() int { // Aplican para uso unico
		return x * 2
	}()
	fmt.Println(y)

	// Funcion anonima en goroutines
	c := make(chan int, 1)
	go func() {
		fmt.Println("Starting process...")
		time.Sleep(time.Second * 5)
		fmt.Println("Process finished!")
		c <- 1
	}()
	<-c // Espera que el canal otenga un valor para continuar

}
