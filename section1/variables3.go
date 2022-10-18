package section1

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Assign With Expressions
//
//  1. Multiply 3.14 with 2 and assign it to `n` variable
//
//  2. Print the `n` variable
//
// HINT
//  Example: 3 * 2 = 6
//  * is the product operator (it multiplies numbers)
//
// EXPECTED OUTPUT
//  6.28
// ---------------------------------------------------------

func AssignWithExpressions() {
	n := 0.
	n = 3.14 * 2

	fmt.Println(n)
}
