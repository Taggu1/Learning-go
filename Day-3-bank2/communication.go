package main

import "fmt"

func printOptions() int {
	fmt.Println("\nWhat do you want to do")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Printf("4. Exit\n\n")

	fmt.Print("Enter choice: ")
	var choice int
	fmt.Scan(&choice)

	return choice

}
