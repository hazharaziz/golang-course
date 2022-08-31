package main

import "fmt"

func main() {
	var arr [4]byte = [4]byte{10, 20, 30, 40}
	fmt.Println("Before")
	for i, v := range arr {
		fmt.Printf("%d: %d\n", i, v)
	}

	fmt.Println("")

	zero(&arr)
	fmt.Println("After")
	for i, v := range arr {
		fmt.Printf("%d: %d\n", i, v)
	}
}

func zero(ptr *[4]byte) {
	*ptr = [4]byte{}
} 