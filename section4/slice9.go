package section4

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Append #3 — Fix the problems
//
//  Fix the problems in the code below.
//
// BONUS
//
//  Simplify the code.
//
// EXPECTED OUTPUT
//
//  toppings: [black olives green peppers onions extra cheese]
//
// ---------------------------------------------------------

func Append3() {
	toppings := []string{"black olives", "green peppers"}

	var pizza []string

	pizza = append(pizza, toppings...)

	pizza = append(pizza, "onions", "extra", "cheese")

	fmt.Printf("pizza       : %s\n", pizza)

	// toppings := []int{"black olives", "green peppers"}

	// var pizza [3]string
	// append(pizza, ...toppings)
	// pizza = append(toppings, "onions")
	// toppings = append(pizza, extra cheese)

	// fmt.Printf("pizza       : %s\n", pizza)
}
