package main

import (
	"fmt"

	"example.org/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	accountBalance = fileops.GetFloatFromFile(balanceFileName)
	fmt.Println("Welcome to Go bank!")
	fmt.Println("Hello,", randomdata.SillyName())
	for {
		userChoice := printOptions()
		actOnChoice(userChoice)
	}

}
