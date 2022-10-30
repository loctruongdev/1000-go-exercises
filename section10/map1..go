package section10

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func Warmup() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number

	phones1 := map[string]string{}
	fmt.Printf("%#v\n", phones1)

	var phones2 map[string]string
	fmt.Printf("%#v\n", phones2)

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable

	products := map[int]bool{}
	fmt.Printf("%#v\n", products)

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	multiPhones := map[string][]string{}
	fmt.Printf("%#v\n", multiPhones)

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity

	basket := map[int]map[int]int{}
	fmt.Printf("%#v\n", basket)
}
