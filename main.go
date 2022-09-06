package main

import "fmt"

type employee struct {
	person
	employeeId int
}

type person struct {
	firstName              string
	lastName               string
	favoriteIceCreamFlavor []string
}

func main() {
	// exercise #1
	p1 := person{
		firstName:              "Hazhar",
		lastName:               "Aziz",
		favoriteIceCreamFlavor: []string{"chocolate", "vanilla"},
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for _, flavour := range p1.favoriteIceCreamFlavor {
		fmt.Println(flavour)
	}
}
