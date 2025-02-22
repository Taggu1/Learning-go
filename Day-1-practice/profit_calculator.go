package main

import "fmt"

func main() {
	var revenue, expense, taxRate, ebt, profit, ratio float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expense: ")
	fmt.Scan(&expense)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt = revenue - expense
	profit = ebt - (1 - taxRate/100)
	ratio = ebt / profit

	fmt.Println("Eraning before tax: ", ebt)
	fmt.Println("Eraning after tax: ", profit)
	fmt.Println("Ratio (ebt/profit): ", ratio)

}
