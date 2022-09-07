package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi, I'm", p.first, p.last, "and I am", p.age, "years of age")
}

func main() {
	p1 := person{
		first: "Hazhar",
		last:  "Aziz",
		age:   22,
	}
	p1.speak()
}
