package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	e := Employee{}
	e.id = 1
	e.name = "Name"
	fmt.Printf("%v", e)
}