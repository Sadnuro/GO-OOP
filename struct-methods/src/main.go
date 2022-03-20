package main

import (
	"fmt"
	"struct-methods/classes"
)

func main() {
	e := classes.Employee{Id: 1, Name: "Default"}
	fmt.Println(e)

	e.SetId(5)
	e.SetName("Sadith")

	fmt.Println(e.GetId(), e.GetName())

	// Devuelve una referencia de memoria a la instancia creada
	e2 := new(classes.Employee)
	e2.Id = 100
	fmt.Println(*e2)

	// Devuelve una referencia de memoria a la instancia creada
	e3 := classes.NewEmployee(3, "Sadith", 12000)
	fmt.Println(e3)

	// p := classes.Person{Age: 20}
	// fmt.Println(p)

}
