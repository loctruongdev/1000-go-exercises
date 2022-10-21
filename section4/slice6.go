package section4

import (
	"fmt"
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//
// EXPECTED OUTPUT
//
//   They are equal.
//
//
// HINTS
//
//   1. strings.Split function splits a string and
//      returns a string slice
//
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.
//
// ---------------------------------------------------------

func CompareSlices() {
	namesA := "Da Vinci, Wozniak, Carmack"

	namesA2 := strings.Split(namesA, ", ")
	sort.Strings(namesA2)
	// fmt.Printf("%q\n", namesA2)

	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	sort.Strings(namesB)
	// fmt.Printf("%q\n", namesB)

	if len(namesA2) == len(namesB) {
		for i := range namesA2 {
			if namesA2[i] != namesB[i] {
				fmt.Println("NOT equal.")
				return
			}
		}
		fmt.Println("They are equal.")
	}

}
