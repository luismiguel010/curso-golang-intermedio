package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//Declaración de vbles
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	//Controlando error
	myValue, err := strconv.ParseInt("NaN", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	}else {
		fmt.Println(myValue)
	}
	//maps
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	//Slice de enteros
	s:= []int{1,2,3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	//Channel
	//c := make(chan int)
	//go doSomething(c)
	//<-c

	//Apuntadores
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

//Función para simular una goRoutine
func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}