package main

import (
	"fmt"
	"os"

	"example.org/bank/fileops"
)

var accountBalance float64 = 1000.0

const balanceFileName = "balance.txt"

func actOnChoice(choice int) {
	switch choice {
	case 1:
		checkBalance()
	case 2:
		depositMoney()
	case 3:
		withDrawMoney()
	case 4:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice...")
	}
	fileops.WriteFloatToFile(accountBalance, balanceFileName)
}

func checkBalance() {
	fmt.Println("Your balance is:", accountBalance)
}

func depositMoney() {
	var amount float64
	fmt.Print("Enter amount to deposit: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("You can only add positive numbers..")
	} else {
		accountBalance += amount
		fmt.Println("Money deposited successfully.")
	}
}

func withDrawMoney() {
	var amount float64
	fmt.Print("Enter amount to withdraw: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("You can only add positive numbers..")
	} else if (accountBalance - amount) < 0 {
		fmt.Println("You don't have enough balance..")
	} else {
		accountBalance -= amount
		fmt.Println("Money withdrawn successfully.")
	}
}
