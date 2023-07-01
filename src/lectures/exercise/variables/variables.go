//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var colour = "Black"
	fmt.Println("My favourite colour is", colour)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	birthYear, ageinYears := 2000, 22
	fmt.Println("My birth year is", birthYear, "and my age is", ageinYears)

	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial = "E"
		lastInitial = "E"
	)
	fmt.Println("My first and last initials are", firstInitial, lastInitial)

	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	var ageInDays int
	ageInDays = 365 * ageinYears
	fmt.Println("My age in days is", ageInDays)
}
