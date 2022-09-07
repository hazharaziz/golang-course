package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("in main func")
	fmt.Println("func main exits")
}

func foo() {
	defer func() {
		fmt.Println("in foo defer")
	}()
	fmt.Println("func foo executing")
}
