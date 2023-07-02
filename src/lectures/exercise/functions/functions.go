//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greeting(name string) {
	fmt.Println("Welcome to Ghostdev-labs,", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func message() string{
	return "This is a good day to code"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func addThreeNum(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

//* Write a function that returns any number
func randNum() int {
	return 10
}

//* Write a function that returns any two numbers
func randNum2() (int, int) {
	return 20, 30
}

//* Call every function at least once

func main() {
	greeting("George")
	fmt.Println(message())

	//* Add three numbers together using any combination of the existing functions.
	num1, num2 := randNum2()
	addNumbers := addThreeNum(randNum(), num1, num2)
	//  * Print the result
	fmt.Println(addNumbers)

}
