package main

import "fmt"

func main() {
	var myName = "Allen"
	fmt.Println("my name is", myName)

	// type anontation
	var name string = "Kathy"
	fmt.Println("name =", name)

	// create & assign
	userName := "admin"
	fmt.Println("username =", userName)

	var sum int
	fmt.Println("The sum is", sum)

	// compound assignment
	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "other is", other)

	// Ok comma
	part2, other := 2, 0
	fmt.Println("part 2 is", part2, "other is", other)

	// reassignment
	sum = part1 + part2
	fmt.Println("sum =", sum)

	// block assignment
	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("Lesson Name =", lessonName)
	fmt.Println("Lesson Type =", lessonType)

	// compound assignment with blank identifier
	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}