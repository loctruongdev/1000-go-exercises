package section1

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #2
//
//  Fix the code by using the conversion expression.
//
// EXPECTED OUTPUT
//  10.5
// ---------------------------------------------------------

func TypeConversion2() {
	a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)
}
