package main

import "fmt"

func main() {
	// exercise #1
	// for i := 0; i <= 10000; i++ {
	// 	fmt.Println(i)
	// }

	// exercise #2
	// index := 1
	// for i := 65; i <= 90; i++ {
	// 	fmt.Println(index)
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Printf("\t\t%#U\n", i)
	// 	}
	// 	index++
	// }

	// exercise #3
	// year := 2000
	// for year <= 2022 {
	// 	if year > 2022 {
	// 		break
	// 	}
	// 	fmt.Println(year)
	// 	year++
	// }

	// exercise #4
	// year := 2000
	// for year <= 2022 {
	// 	fmt.Println(year)
	// 	year++
	// }

	// exercise #5
	// for i := 10; i <= 100; i++ {
	// 	fmt.Printf("%v\n", i%4)
	// }

	// exercise #6, #7
	// x := 33
	// fmt.Println("using if statement")
	// if x < 32 {
	// 	fmt.Println("x is less than 32")
	// } else if x == 32 {
	// 	fmt.Println("x is equal to 32")
	// } else {
	// 	fmt.Println("x is greater than 32")
	// }

	// exercise #8
	// y := 40
	// switch {
	// case y < 32:
	// 	fmt.Println("y is less than 32")
	// case y == 32:
	// 	fmt.Println("y is equal to 32")
	// case y > 32:
	// 	fmt.Println("y is greater than 32")
	// }

	// exercise #9
	favSport := "swimming"
	switch favSport {
	case "football":
		fmt.Println("1")
	case "volleyball":
		fmt.Println("2")
	case "basketball":
		fmt.Println("3")
	case "swimming":
		fmt.Println("4")
	}
}
