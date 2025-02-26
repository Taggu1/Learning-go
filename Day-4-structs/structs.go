package main

import (
	"fmt"

	"example.org/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	userStruct, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println("Some user details are missing....")
		return
	}

	userStruct.OutputDetails()
	userStruct.ClearUserName()
	userStruct.OutputDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
