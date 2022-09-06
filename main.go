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
	p1 := employee{
		person: person{
			firstName: "Hazhar",
			lastName:  "Aziz",
		},
		employeeId: 10,
	}

	fmt.Println(p1)
	fmt.Println(p1.firstName, p1.lastName)
	fmt.Println(p1.person.firstName, p1.person.lastName)
}
