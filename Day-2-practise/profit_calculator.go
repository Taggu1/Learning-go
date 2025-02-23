package main

import "fmt"

func main() {

	revenue := input("Enter revenue: ")
	expense := input("Enter expense: ")
	taxRate := input("Enter tax rate: ")

	ebt, profit, ratio := calculator(revenue, expense, taxRate)

	fmt.Println("Eraning before tax: ", ebt)
	fmt.Println("Eraning after tax: ", profit)
	fmt.Printf("Ratio (ebt/prfit): %.2f\n", ratio)
}

func input(text string) float64 {
	var userText float64
	fmt.Print(text)
	fmt.Scan(&userText)
	return userText
}

func calculator(revenue, expense, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expense
	profit := ebt - (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
