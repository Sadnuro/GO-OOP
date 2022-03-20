package main

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}
type FullTimEmployee struct {
	Employee
	Person
}
