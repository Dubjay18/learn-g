//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price int
}

func calculateTotalCost(list [4]Product) int {
	var sum int
	for i := 0; i < len(list); i++ {
		item := list[i]
		sum += item.price
	}
	return sum
}
func main() {
	shoppingList := [4]Product{}
	shoppingList[0] = Product{
		name:  "gun",
		price: 25,
	}
	shoppingList[1] = Product{
		name:  "tissue",
		price: 205,
	}
	shoppingList[2] = Product{
		name:  "katana",
		price: 2005,
	}

	fmt.Println(shoppingList[len(shoppingList)-1], len(shoppingList), calculateTotalCost(shoppingList))
	shoppingList[3] = Product{
		name:  "Watch",
		price: 20000005,
	}
	fmt.Println(shoppingList[3])
}
