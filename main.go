package main

import "fmt"

type myType int

var x myType
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println("x: ", x)

	y = int(x)
	fmt.Println("y: ", y)
}
