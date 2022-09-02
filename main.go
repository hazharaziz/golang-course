package main

import "fmt"

var y = "Hello"

type text string

var x text = "Hi"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// conversion

	y = string(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	x = text(y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

}
