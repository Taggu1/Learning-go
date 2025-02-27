package main

import (
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	hobbies := [3]string{"Programming", "Working out", "Poetry"}
	fmt.Printf("hobbies: %v\n", hobbies)

	fmt.Printf("First hobby: %v\n", hobbies[0])
	fmt.Printf("Rest of hobbies: %v\n", hobbies[1:])

	firstAndSecondHobbies := hobbies[:2]
	fmt.Printf("firstAndSecondHobbies: %v\n", firstAndSecondHobbies)
	firstAndSecondHobbies2 := hobbies[0:2]
	fmt.Printf("firstAndSecondHobbies2: %v\n", firstAndSecondHobbies2)

	firstAndSecondHobbies = hobbies[1:]
	fmt.Printf("firstAndSecondHobbies2: %v\n", firstAndSecondHobbies)

	goals := []string{"Goal 1", "Goal 2"}

	goals[1] = "Whaat"

	goals = append(goals, "Goal 3")
	fmt.Printf("goals: %v\n", goals)

	products := []Product{
		{
			title: "Product 1",
			id:    "a34234234",
			price: 20,
		},
		{
			title: "Product 1",
			id:    "a34234234",
			price: 20,
		},
	}
	fmt.Printf("products: %v\n", products)

	dynamicTypesArray := []any{1, 1.3, "Nuce"}

	fmt.Printf("dynamicTypesArray: %v\n", dynamicTypesArray)

}
