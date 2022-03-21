package main

import "fmt"

/* jobs: Recibe todos los trabajos asignados al workers
results: Guarda valores calculados en la serie
*/
func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d,  started job with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d\n", id, job, fib)

		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40} // Numeros a calcular la serie
	nWorkers := 3                             // Cantidad de workers que calcularán la serie para cada numero en tasks

	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	// Se lanzan 3 workers que estarán cada uno escuchando en su canal
	// Al momento de recibir datos por el canal, procesarán la serie
	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	// Asigna cada valor a calcular a un canal especifico
	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}
