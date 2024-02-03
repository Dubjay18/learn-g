package main

import "fmt"

func printslice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {
	route := []string{"Grocery", "Department Store", "Salon"}
	printslice("Route 1", route)

	route = append(route, "Home")
	printslice("Route 2", route)
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")

	route = route[2:]
	printslice("Remaining locations", route)

}
