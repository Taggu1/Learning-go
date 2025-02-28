package main

import "fmt"

func main() {

	sum := sumUp(2, 56, 7, 7, 4, 2)
	fmt.Println(sum)
}

func sumUp(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
