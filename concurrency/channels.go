package main

import "fmt"

func main() {
	// Unbuffered channels y buffered channels

	// Canal sin buffer genera bloqueo
	// requiere un programa que espere el resultado
	// c := make(chan int)
	// c <- 1
	// fmt.Println(<-c)

	// Canal con buffers
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
