package section4

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Observe the capacity growth
//
//  Write a program that appends elements to a slice
//  10 million times in a loop. Observe how the capacity of
//  the slice changes.
//
//
// STEPS
//
//  1. Create a nil slice
//
//  2. Loop 1e7 times
//
//  3. On each iteration: Append an element to the slice
//
//  4. Print the length and capacity of the slice "only"
//     when its capacity changes.
//
//  BONUS: Print also the growth rate of the capacity.
//
//
// EXPECTED OUTPUT
//
//  len:0               cap:0               growth:NaN
//  len:1               cap:1               growth:+Inf
//  len:2               cap:2               growth:2.00
//  ... and so on.
//
// ---------------------------------------------------------

func ObserveCapacityGrowth() {
	var (
		slice  []string
		oldCap float64
	)

	for len(slice) < 1e7 {
		c := float64(cap(slice))

		if c == 0 || c != oldCap {
			fmt.Printf("len: %-15dcap: %-15ggrowth: %-15.2f\n", len(slice), c, c/oldCap)
		}

		slice = append(slice, "elem")

		oldCap = c
	}
}
