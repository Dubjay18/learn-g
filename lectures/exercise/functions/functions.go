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

func greet(name string) string {
	return "yo" + name
}
func rando() string {
	return "i alone i'm the honoured one"
}
func sum(num1, num2, num3 int) int {
	return num1 + num2 + num3
}
func anyNum() int {
	return 909
}
func two() (int, int) {
	return 4, 4
}
func main() {
	fmt.Println(greet("Jay"))
	fmt.Println(rando())
	fmt.Println(sum(1, 2, 3))
	fmt.Println(anyNum())

}
