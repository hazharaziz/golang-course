package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	p1 := person{
		firstName: "Hazhar",
		lastName:  "Aziz",
	}

	fmt.Println(p1)
	fmt.Println(p1.firstName, p1.lastName)

}
