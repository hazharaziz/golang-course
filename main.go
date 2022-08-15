package main

import "fmt"

// declare and assign
var y = 20

// just declare and set zero value of 0 for type int
var v int

func main() {
	// short declaration (inside func body)
	x := 100
	fmt.Println(x)

	// var for both func body and outside
	var z = 30
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(v)
}

