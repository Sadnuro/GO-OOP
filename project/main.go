package main

import (
	"fmt"
	"time"
)

/*	Task to execute: Numero del cual calcular la serie de Fibonacci
 */
type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

/*
 */
type Worker struct {
	Id         int
	JobQueue   chan Job      // Leerá los jobs que van llegando
	WorkerPool chan chan Job // Canal de canales
	QuitChan   chan bool     // Canal en caso de peticion de cierre de workers
}

/* Envia todo a los workers
 */
type Dispatcher struct {
	WorkerPool chan chan Job // Canal de canales que enviará Jobs
	MaxWorkers int           // Cant. Max. de Concurrencia
	JobQueue   chan Job
}

func (w Worker) Start() {
	go func() {
		for {
			// Lee JobQueue y asigna a WorkerPool
			w.WorkerPool <- w.JobQueue

			select {
			// JobQueue: Contiene todos los trabajos a ejecutar por el worker
			case job := <-w.JobQueue: // Leyendo job de la cola
				fmt.Printf("Started: Worker id: %d\n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Finished: Worker id: %d | Job.Number: %d | Fibonacci result: %d\n", w.Id, job.Number, fib)
			case <-w.QuitChan:
				fmt.Printf("Stopped: Worker id: %d\n", w.Id)
			}
		}
	}()
}

// Worker constructor
func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job), // asigna canal nuevo
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

// Enviar marca de detención al QuitChan del worker especificado
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	// maxWorkers: Cantidad maxima de trabajadores para el canal
	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

// Despacha los jobs (trabajos/tareas) que van llegando
// A la jobQueue de los Workers
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()

		}
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
