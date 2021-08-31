package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func NewEmployee (id int, name string) *Employee {
	return &Employee{ //Se coloca con & para devolver la referencia, porque sino Go los trata como una copia y no es tan óptimo
		id: id, 
		name: name,
	}
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e4 := NewEmployee(1, "Name2")
	fmt.Println(e4.GetId())
	fmt.Println(e4.GetName())
}