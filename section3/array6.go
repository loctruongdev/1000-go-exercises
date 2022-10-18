package section3

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Compare the Arrays
//
//  1. Uncomment the code
//
//  2. Fix the problems so that arrays become comparable
//
// EXPECTED OUTPUT
//  true
//  true
//  false
// ---------------------------------------------------------

func CompareArrays() {
	week := [...]string{"Monday", "Tuesday"}
	wend := [...]string{"Saturday", "Sunday"}

	fmt.Println(week != wend)

	evens := [...]int{2, 4, 6, 8, 10}
	evens2 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(evens == evens2)

	// Use     : uint8 for one of the arrays instead of byte here.
	// Remember: Aliased types are the same types.

	image := [5]uint8{'h', 'i'}
	var data [5]byte

	fmt.Println(data == image)
}
