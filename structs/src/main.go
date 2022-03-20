package main

import (
	"fmt"
	"structs/class"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	p := Person{Id: 0, Name: "Sadith"}

	fmt.Println(p)
	p.Age = 20
	fmt.Println(p)

	// Employee with zero values
	e := class.Employee{}
	fmt.Println(e)

	// Employee with initial values
	e2 := class.Employee{
		Id:       1,
		Name:     "Sadith",
		LastName: "Nunez",
		Position: "Senior Backend Go Developer",
		Salary:   12000,
		Currency: "USD",
		Company:  "Google INC",
	}
	fmt.Println(e2)
}
