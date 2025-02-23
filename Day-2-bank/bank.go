package main

import (
	"fmt"
	"os"
	"strconv"
)

var accountBalance float64 = 1000.0

const balanceFileName = "balance.txt"

func main() {
	accountBalance = readBalanceFromFile()
	fmt.Println("Welcome to Go bank!")

	for {
		userChoice := printOptions()
		actOnChoice(userChoice)
	}

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(balanceText), 0644)
}

func readBalanceFromFile() float64 {
	balanceText, err := os.ReadFile(balanceFileName)

	if err != nil {
		writeBalanceToFile(0)
		return 0
	}

	balance, err := strconv.ParseFloat(string(balanceText), 64)

	if err != nil {
		return 0
	}

	return balance
}

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
	writeBalanceToFile(accountBalance)
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
