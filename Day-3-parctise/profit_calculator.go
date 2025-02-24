package main

import (
	"errors"
	"fmt"
	"os"
)

const permessionsNumber = 0644

func main() {

	revenue, err1 := input("Enter revenue: ")
	expense, err2 := input("Enter expense: ")
	taxRate, err3 := input("Enter tax rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	ebt, profit, ratio := calculator(revenue, expense, taxRate)

	fmt.Println("Eraning before tax: ", ebt)
	fmt.Println("Eraning after tax: ", profit)
	fmt.Printf("Ratio (ebt/prfit): %.2f\n", ratio)

	storeIntoFile(ebt, profit, ratio)
}

func storeIntoFile(ebt, profit, ratio float64) {
	resaultsString := fmt.Sprintf("%.2f\n%.2f\n%.2f\n", ebt, profit, ratio)

	os.WriteFile("database.txt", []byte(resaultsString), permessionsNumber)
}

func input(text string) (float64, error) {
	var userNumber float64
	fmt.Print(text)
	_, err := fmt.Scan(&userNumber)

	if err != nil {
		return 0, errors.New("number should be greater than zero")
	}

	return userNumber, nil
}

func calculator(revenue, expense, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expense
	profit := ebt - (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
