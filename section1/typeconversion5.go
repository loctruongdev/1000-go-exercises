package section1

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #5
//
//  Fix the code.
//
// HINTS
//   maximum of int8  can be 127
//   maximum of int16 can be 32767
//
// EXPECTED OUTPUT
//  1127
// ---------------------------------------------------------

func TypeConversion5() {
	// DO NOT TOUCH THESE VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(max + int16(min))
}

/* EXPLANATION
`int8(max)` destroys the information of max
It reduces it to 127
Which is the maximum value of int8

Correct conversion is int16(min)
Because, int16 > int8
When you do so, min doesn't lose information
*/
