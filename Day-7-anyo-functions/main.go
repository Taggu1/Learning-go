package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(i int) int {
		return i * 2
	})

	doubledNumbers := transformNumbers(&numbers, double)
	tribledNumbers := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubledNumbers)
	fmt.Println(tribledNumbers)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
