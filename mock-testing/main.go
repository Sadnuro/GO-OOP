package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}
type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(time.Second * 5)
	// SELECT * FROM PERSON WHERE...
	return Person{}, nil
}
var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(time.Second * 5)
	// SELECT * FROM PERSON WHERE...
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Employee = e

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = p

	return ftEmployee, nil
}
