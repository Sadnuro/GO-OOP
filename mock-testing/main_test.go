package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{Id: 1, Position: "CEO"}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{DNI: "1", Name: "Sadith", Age: 20}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					DNI:  "1",
					Name: "Sadith",
					Age:  20,
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGerPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()

		ft, err := GetFullTimeEmployeeById(test.id, test.dni)
		if err != nil {
			t.Errorf("Error when getting employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)
		}

	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGerPersonByDNI

}
