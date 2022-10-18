package section1

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Find the Rectangle's Perimeter
//
//  1. Find the perimeter of a rectangle
//     Its width is  5
//     Its height is 6
//
//  2. Assign the result to the `perimeter` variable
//
//  3. Print the `perimeter` variable
//
// HINT
//  Formula = 2 times the width and height
//
// EXPECTED OUTPUT
//  22
//
// BONUS
//  Find more formulas here and calculate them in new programs
//  https://www.mathsisfun.com/area.html
// ---------------------------------------------------------

func RectanglePerimeter() {
	var (
		perimeter     int
		width, height = 5, 6
	)

	perimeter = (width + height) * 2

	fmt.Println(perimeter)
}

func RectanglePerimeter2(w, h int) {
	perimeter := (w + h) * 2

	fmt.Println(perimeter)
}
