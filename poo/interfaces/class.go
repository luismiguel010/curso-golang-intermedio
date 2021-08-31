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
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time Employee"
}

type TemporaryEmployee struct {
	Persona
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary time Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}