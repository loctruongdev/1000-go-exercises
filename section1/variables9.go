package section1

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Swapper #2
//
//  1. Swap the values of `red` and `blue` variables
//
//  2. Print them
//
// EXPECTED OUTPUT
//  blue red
// ---------------------------------------------------------

func MySwapper2() {
	color, color2 := "red", "blue"
	color, color2 = color2, color
	fmt.Println(color, color2)
}
