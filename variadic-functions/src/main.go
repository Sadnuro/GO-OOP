package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func variadicSumFunc(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (double int, triple int, quad int) {
	// Go retorna todas las variables de retorno declaradas
	double = x * 2
	triple = x * 3
	quad = x * 4
	return
}

func main() {
	/** VARIADIC FUNCTIONS
	Si se tienes una funcion de suma esta define la cantidad
	de parámetros que debe recibir. Pero no siempre se quiere
	sumar 2 numeros, pueden ser 3, 4, 5 o [n] cantidad de numeros

	Lo típico sería crear una funcion que reciba la cantidad de
	parámetros que se requiera; Esto sería impráctico.

	La funcion variádica permite realizar operaciones para
	cualquier cantidad de parámetros del mismo tipo de datos
	*/

	// fmt.Println(sum(4, 5))
	fmt.Println(variadicSumFunc(4, 5, 7, 9))

	printNames("Sadith", "Jorge", "Zenaida", "Luises")

	fmt.Println(getValues(3))
	double, triple, quad := getValues(5)
	fmt.Println(double, triple, quad)
}
