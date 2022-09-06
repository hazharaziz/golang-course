package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	// exercise #3
	myTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}
	fmt.Println("------- truck --------")
	fmt.Println("doors:", myTruck.doors)
	fmt.Println("color:", myTruck.color)
	fmt.Println("fourWheel:", myTruck.fourWheel)

	mySedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}
	fmt.Println("------- sedan --------")
	fmt.Println("doors:", mySedan.doors)
	fmt.Println("color:", mySedan.color)
	fmt.Println("luxury:", mySedan.luxury)
}
