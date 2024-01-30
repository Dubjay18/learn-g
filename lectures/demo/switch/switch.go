package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("cheap")

	case p < 10:
		fmt.Println("Moderate")
	default:
		fmt.Println("Expensive item")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Busines seating")
	case FirstClass:
		fmt.Println("first class seating")
	default:
		fmt.Println("Other seating")
	}

}
