package main

import "fmt"

func main() {
	anonymousFunc := func(number int) {
		fmt.Println("anonymous func")
		fmt.Println(number)
	}
	anonymousFunc(42)
}
