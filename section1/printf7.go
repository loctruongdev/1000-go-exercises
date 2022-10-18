package section1

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the Type #2
//
//  Print the type and value of 3.14 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3.14 is float64
// ---------------------------------------------------------

func PrintType2() {
	fmt.Printf("Type of %.2f is %[1]T\n", 3.14)
}
