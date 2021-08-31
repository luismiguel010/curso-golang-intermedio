package main

import "fmt"

type Persona struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Persona
	Employee
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 25
	ftEmployee.id = 5
	fmt.Printf("%v", ftEmployee)
}