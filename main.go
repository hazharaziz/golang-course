package main

import "fmt"

type employee struct {
	person
	employeeId int
}

type person struct {
	firstName string
	lastName  string
}

func main() {
	p1 := struct {
		firstName string
		lastName  string
	}{
		firstName: "Hazhar",
		lastName:  "Aziz",
	}

	fmt.Println(p1)
	fmt.Println(p1.firstName, p1.lastName)
}
