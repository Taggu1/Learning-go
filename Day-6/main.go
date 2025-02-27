package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 4)
	userNames = append(userNames, "Hisham")
	userNames = append(userNames, "Max")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["Go"] = 4.5
	courseRatings["React"] = 2.2
	courseRatings["Flutter"] = 5.0
	courseRatings.output()

	for index, userName := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", userName)
	}
}
