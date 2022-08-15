package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	foo()

	for i := 0; i < 20; i++ {
		if i % 2 == 0 {
			fmt.Println("Even Number:", i)
		}
	}

	bar()

	fmt.Println("Exit the program!")
}

func foo() {
	fmt.Println("In foo func!")
}

func bar()  {
	fmt.Println("Last func")
}