package main

import "fmt"

func main() {
	var myName = "Jayson"
	fmt.Println("my name is", myName)

	var name string = "Dorthy"
	fmt.Println("name =", name)

	userName := "admin"
	fmt.Println("username =", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "other is", other)
	part2 := 9
	sum = part1 + part2
	fmt.Println("sum =", sum)

	var (
		lessonName = "variables"
		lessonType = "Demo"
	)

	fmt.Println("lessonName", lessonName)
	fmt.Println("lessonType", lessonType)

}
