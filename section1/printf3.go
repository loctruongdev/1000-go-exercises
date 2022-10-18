package section1

import "fmt"

// ---------------------------------------------------------
// EXERCISE: False Claims
//
//  Use printf to print the expected output using a variable.
//
// EXPECTED OUTPUT
//  These are false claims.
// ---------------------------------------------------------

func FalseClaims() {
	tf := false
	fmt.Printf("These are %t claims.\n", tf)
}
