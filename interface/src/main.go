package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Employee struct {
	id int
}

type FulltimeEmployee struct {
	// Esta notacion obliga a acceder a cada struct
	// Para manipular los atributos
	employee Employee
	person   Person
	endDate  string
}

type ParcialtimeEmployee struct {
	// Esta notacion permite un acceso directo a los atributos
	Employee
	Person
	hours int
}

// Metodo para cada clase
// Equivalente a implementaciones de getMsg() en cada clase
// Implementan de manera implicita el metodo getMsg() de PrintInfo interface
func (fte FulltimeEmployee) getMsg() string {
	return "Full Time Employee"
}
func (fte ParcialtimeEmployee) getMsg() string {
	return "Parcial Time Employee"
}

// interface para implementar todos los metodos getMsg()
type PrintInfo interface {
	getMsg() string
}

func getMsg(p PrintInfo) {
	fmt.Println(p.getMsg())
}

func main() {

	ftEmployee := FulltimeEmployee{}
	ftEmployee.employee.id = 1
	ftEmployee.person.name = "Sadith"
	ftEmployee.person.age = 20

	fmt.Println(ftEmployee)

	ptEmployee := ParcialtimeEmployee{}
	ptEmployee.id = 2
	ptEmployee.name = "Vanessa"
	ptEmployee.age = 34
	fmt.Println(ptEmployee)

	getMsg(ftEmployee)
	getMsg(ptEmployee)
}
