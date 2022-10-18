package section3

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Assign the Arrays
//
//  1. Create an array named books
//
//  2. Add book titles to the array
//
//  3. Create two more copies of the array named: upper and lower
//
//  4. Change the book titles to uppercase in the upper array only
//
//  5. Change the book titles to lowercase in the lower array only
//
//  6. Print all the arrays
//
//  7. Observe that the arrays are not connected when they're copied.
//
// NOTE
//  Check out the strings package, it has functions to convert letters to
//  upper and lower cases.
//
// BONUS
//  Invent your own arrays with different types other than string,
//  and do some manipulations on them.
//
// EXPECTED OUTPUT
//   Note: Don't worry about the book titles here, you can use any title.
//
//   books: ["Kafka's Revenge" "Stay Golden" "Everythingship"]
//   upper: ["KAFKA'S REVENGE" "STAY GOLDEN" "EVERYTHINGSHIP"]
//   lower: ["kafka's revenge" "stay golden" "everythingship"]
// ---------------------------------------------------------

func AssignArrays() {

	books := [...]string{
		"My Book1",
		"My Book2",
		"My Book3",
	}

	// upper := books
	// lower := books
	upper, lower := books, books

	// won't effect
	// for _, v := range upper {
	// 	v = strings.ToUpper(v)
	// }

	// for i, _ := range upper {
	// 	upper[i] = strings.ToUpper(upper[i])
	// }

	// for i, _ := range lower {
	// 	lower[i] = strings.ToLower(lower[i])
	// }

	for i := range books {
		upper[i] = strings.ToUpper(upper[i])
		lower[i] = strings.ToLower(lower[i])
	}

	fmt.Printf("books: %q\n", books)
	fmt.Printf("upper: %q\n", upper)
	fmt.Printf("lower: %q\n", lower)

}
