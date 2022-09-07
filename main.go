package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	oddsTotal := sumOfOddNumbers(sum, numbers)
	fmt.Println(oddsTotal)
}

func sum(numbers []int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func sumOfOddNumbers(sumCallback func(numbers []int) int, allNumbers []int) int {
	var oddNumbers []int = make([]int, 0)
	for _, number := range allNumbers {
		if number%2 == 1 {
			oddNumbers = append(oddNumbers, number)
		}
	}
	sum := sumCallback(oddNumbers)
	return sum
}
