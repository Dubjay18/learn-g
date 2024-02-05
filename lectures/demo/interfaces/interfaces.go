package main

import "fmt"

type Preparer interface {
	PrepareDish()
}
type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}
func (s Salad) PrepareDish() {
	fmt.Println("toss salad")
}

func prepareDishes(dishes []Preparer) {
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}
func main() {
	dishes := []Preparer{Chicken("chicken wings"), Salad("garden salad")}
	prepareDishes(dishes)

	dishes = append(dishes, Chicken("chicken parm"))
	prepareDishes(dishes)

	dishes = append(dishes, Salad("caesar salad"))

	prepareDishes(dishes)

}
