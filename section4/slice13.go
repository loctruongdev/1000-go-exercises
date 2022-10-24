package section4

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

func SliceNumbers() {

	data := "2 4 6 1 3 5"

	// dataSlice := strings.Split(data, " ")
	dataSlice := strings.Fields(data)

	var (
		dataInt, dataEvens, dataOdds []int
	)

	for _, v := range dataSlice {
		n, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		dataInt = append(dataInt, n)

	}

	fmt.Printf("dataInt: %v\n", dataInt)

	for _, v := range dataInt {
		if v%2 == 0 {
			dataEvens = append(dataEvens, v)
		} else {
			dataOdds = append(dataOdds, v)
		}
	}

	middle := dataInt[2:4]
	first := dataInt[0:2]
	last := dataInt[len(dataInt)-2:]
	lastEvens := dataEvens[len(dataEvens)-1:]
	lastOdds := dataOdds[len(dataOdds)-2:]

	fmt.Printf("dataEvens: %v\n", dataEvens)
	fmt.Printf("dataOdds: %v\n", dataOdds)
	fmt.Printf("middle: %v\n", middle)
	fmt.Printf("first: %v\n", first)
	fmt.Printf("last: %v\n", last)
	fmt.Printf("lastEven: %v\n", lastEvens)
	fmt.Printf("lastOdds: %v\n", lastOdds)

}
