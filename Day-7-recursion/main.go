package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return factorial(n-1) * n
}
