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
}

type ParcialtimeEmployee struct {
	// Esta notacion permite un acceso directo a los atributos
	Employee
	Person
}

func getMsg(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {

	ftEmployee := FulltimeEmployee{}
	ftEmployee.employee.id = 1
	ftEmployee.person.name = "Sadith"
	ftEmployee.person.age = 20

	fmt.Println(ftEmployee)
	getMsg(ftEmployee.person)

	ptEmployee := ParcialtimeEmployee{}
	ptEmployee.id = 2
	ptEmployee.name = "Vanessa"
	ptEmployee.age = 34
	fmt.Println(ptEmployee)

	getMsg(ptEmployee.Person)
}
