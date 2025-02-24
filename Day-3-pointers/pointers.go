package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("Age:", age)
	editAgeToAdultYears(agePointer)
	fmt.Scan()
	fmt.Println("Adult years:", age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
