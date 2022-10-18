package section1

import "fmt"

////// TYPE-CONVERSION EXERCISES /////
// ---------------------------------------------------------
// EXERCISE: Convert and Fix
//
//  Fix the code by using the conversion expression.
//
// EXPECTED OUTPUT
//  15.5
// ---------------------------------------------------------

func TypeConversion() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)
}
