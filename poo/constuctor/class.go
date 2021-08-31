package main

import "fmt"

type Employee struct {
	id   int
	name string
	vacation bool
}

func NewEmployee (id int, name string, vacation bool) *Employee {
	return &Employee{ //Se coloca con & para devolver la referencia, porque sino Go los trata como una copia y no es tan óptimo
		id: id, 
		name: name,
		vacation: vacation,
	}
}

func main() {
	// 1
	e := Employee{}
	fmt.Printf("%v\n", e)
	e.id = 1
	e.name = "Name"
	e.vacation = true

	// 2
	e2 := Employee{
		id: 1,
		name: "Name",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)

	//3
	e3 := new(Employee)
	fmt.Printf("%v", *e3)
	e3.id = 1
	e3.name = "Name"
	e3.vacation = true

	//4
	e4 := NewEmployee(1, "Name2", true)
	fmt.Printf("%v", *e4)
}