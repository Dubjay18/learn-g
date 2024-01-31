//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type rectangle struct {
	width  int
	height int
}

func (rect rectangle) area() int {
	return rect.height * rect.width
}
func perimeter(rect rectangle) int {
	return 2 * (rect.height + rect.width)
}
func main() {
	r := rectangle{
		width:  10,
		height: 20,
	}
	fmt.Println(r.area(), perimeter(r))
}
