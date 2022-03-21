package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

// Crea Workers: Distribuye jobs a los workers
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

// Itera sobre maxWorkers
// Crea los Workers, de acuerdo al WorkerPool del dispatcher
// WorkerPool: Canales para los trabajadores
func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		// i: id del trabajador | workerpool: Grupo de trabajo al que pertenece
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Maneja toda la capacidad de procesamiento del servidor
func RequestHandler(res http.ResponseWriter, req *http.Request, jobQueue chan Job) {
	/*	Req format:
		name: Fib1
		value: 65
		delay: 3s
	*/

	if req.Method != "POST" { // Not allowed: GET, PUT, DELET, etc.
		res.Header().Set("Allow", "POST")
		res.WriteHeader(http.StatusMethodNotAllowed)
	}

	// FormValue: Permite acceder a todos los parametros del request
	// Obtiene parametro [Delay]
	delay, err := time.ParseDuration(req.FormValue("delay"))
	if err != nil {
		http.Error(res, "Invalid [delay]", http.StatusBadRequest)
		return
	}
	// Obtiene parametro [Number]
	value, err := strconv.Atoi(req.FormValue("value"))
	if err != nil {
		http.Error(res, "Invalid [value]", http.StatusBadRequest)
		return
	}
	// Obtiene parametro [name]
	name := req.FormValue("name")
	if name == "" {
		http.Error(res, "Invalid [name]", http.StatusBadRequest)
		return
	}

	job := Job{Name: name, Delay: delay, Number: value}
	jobQueue <- job
	res.WriteHeader(http.StatusCreated)
}

func main() {

	const (
		maxWorkers   = 4       // Max Quantity of workers
		maxQueueSize = 20      // Max Quantity of Jobs processed simultaneously
		port         = ":8081" // Server PORT
	)

	jobQueue := make(chan Job, maxQueueSize)          // Manejar todos los jobs recibidos en las peticiones
	dispatcher := NewDispatcher(jobQueue, maxWorkers) //

	dispatcher.Run()

	// http://localhost:8081/fib
	http.HandleFunc("/fib", func(res http.ResponseWriter, req *http.Request) {
		RequestHandler(res, req, jobQueue)
	})

	// Ejecuta si la funcion deja de funcionar
	log.Fatal(http.ListenAndServe(port, nil))
}
