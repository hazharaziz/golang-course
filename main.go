package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	total := foo(numbers...)
	fmt.Println(total)

	otherTotal := bar(numbers)
	fmt.Println(otherTotal)
}

func foo(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func bar(numbers []int) int {
	return foo(numbers...)
}
