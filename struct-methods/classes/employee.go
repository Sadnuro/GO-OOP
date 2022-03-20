package classes

type Employee struct {
	Id     int
	Name   string
	Salary float64
}

func NewEmployee(id int, name string, salary float64) *Employee {
	return &Employee{
		Id:     id,
		Name:   name,
		Salary: salary,
	}
}

func (e *Employee) SetId(id int) {
	e.Id = id
}
func (e *Employee) GetId() int {
	return e.Id
}

func (e *Employee) SetName(name string) {
	e.Name = name
}
func (e *Employee) GetName() string {
	return e.Name
}

/*	Receiver functions
 */
