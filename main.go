package main

import "fmt"

func main() {
	x := foo()
	y, z := bar()

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func foo() int {
	return 20
}

func bar() (int, string) {
	return 30, "Hello"
}
