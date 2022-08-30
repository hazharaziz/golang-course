package main

import "fmt"

func main() {
	x := 1
	y := &x
	fmt.Println(*y)
	fmt.Println("address:", y)
	*y = 2
	fmt.Println(x)
	fmt.Println("address:", &x)

	p := 1
	incrementByPointer(&p)
	fmt.Println(p)

	p = 1
	incrementByValue(p)
	fmt.Println(p)
}


func incrementByValue(p int) int {
	p++
	return p
}

func incrementByPointer(p *int) int {
	*p++
	return *p
}

