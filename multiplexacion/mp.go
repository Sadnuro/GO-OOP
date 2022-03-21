package main

import (
	"fmt"
	"time"
)

// i: Tiempo para dormir | c: enviar data | param: canal utilizado
func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}

func main() {
	// MULTIPLEXION
	/*
		Cuando una rutina se está comunicando con varios [channels]
		es muy util utilizar la palabra reservada [select]
		para poder intearactuar de manera más ordenada con todos
		los mensajes que están siendo recibidos
	*/

	c1 := make(chan int)
	c2 := make(chan int)

	// duraciones
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	// d1: duracion | c1: channel | channel num equivalent
	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)

	// C1: Demora más pero es el primero en imprimirse
	// ya que el programa espera hasta que haya un valor en c1
	// lo ideal es que se imprima el primero que termine
	// fmt.Println(<-c1)
	// fmt.Println(<-c2)

	// Lectura dinamica de canales finlaizads
	for i := 0; i < 2; i++ {
		// imprime el valor del canal que se valla comunicando
		select {
		case msgChannel1 := <-c1:
			fmt.Println(msgChannel1)
		case msgChannel2 := <-c2:
			fmt.Println(msgChannel2)
		}
	}
}
