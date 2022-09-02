package main

import "fmt"

func main() {

	// exercise #1
	// x := [5]int{23, 5, 21, 53, 66}
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("%T\n", x)

	// exercise #2
	// y := []int{23, 5, 21, 53, 66, 10, 11, 222, 25, 90}
	// for i, v := range y {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("%T\n", y)

	// exercise #3
	// z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// fmt.Println(z[:5])
	// fmt.Println(z[5:])
	// fmt.Println(z[2:7])
	// fmt.Println(z[1:6])

	// exercise #4
	// a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// a = append(a, 52)
	// fmt.Println(a)

	// a = append(a, 53, 54, 55)
	// fmt.Println(a)

	// b := []int{56, 57, 58, 59, 60}
	// a = append(a, b...)
	// fmt.Println(a)

	// exercise #5
	c := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	c = append(c[0:3], c[6:]...)
	fmt.Println(c)

}
