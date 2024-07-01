package main

import "fmt"

func main() {
	age := 29
	agePointer := &age

	// adultYears := getAdultYears(agePointer)

	fmt.Println("Age: ", *agePointer)
	fmt.Println("Age Address: ", agePointer)
	// fmt.Println("Adult years: ", adultYears)

	aditAgeToAdultYears(agePointer)
	fmt.Println("Adult years: ", age)
}

/*
func getAdultYears(age *int) int {
	return *age - 18
}
*/

// with side effect instead of return
func aditAgeToAdultYears(age *int) {
	*age -= 18
}
