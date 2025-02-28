package main

import "fmt"

func main() {
	numbers := []float64{34, 5, 6, 7, 8, 2}

	dNums := transformNumbers(&numbers, triple)
	fmt.Println(dNums)
}

func transformNumbers[T int | float64](numbers *[]T, transform func(T) T) []T {
	dNumbers := []T{}
	for _, number := range *numbers {
		dNumbers = append(dNumbers, transform(number))
	}
	return dNumbers
}

func getGetTransformerFunction() func() func(int) int {
	return getGetTransformerFunction()
}

func getTransformerFunction() func(int) int {
	return double
}

func double[T int | float64](number T) T {
	return number * 2
}

func triple[T int | float64](number T) T {
	return number * 3
}
