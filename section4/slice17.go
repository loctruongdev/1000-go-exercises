package section4

import (
	"fmt"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort the backing array
//
//  1. Sort only the middle 3 items.
//
//  2. All the slices should see your changes.
//
//
// RESTRICTION
//
//  Do not sort manually. Sort by using the sort package.
//
//
// EXPECTED OUTPUT
//
//  Original: [pacman mario tetris doom galaga frogger asteroids simcity metroid defender rayman tempest ultima]
//
//  Sorted  : [pacman mario tetris doom galaga asteroids frogger simcity metroid defender rayman tempest ultima]
//
//
// HINT:
//
//   Middle items are         : [frogger asteroids simcity]
//
//   After sorting they become: [asteroids frogger simcity]
//
// ---------------------------------------------------------

func SortBackingArray() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)
	// ADD YOUR CODE HERE

	mid := items[5:8]
	sort.Strings(mid) // mid slice has the same backing array with items slice, so after sorting it also affect on items

	fmt.Println()
	fmt.Println("Sorted  :", items)
}
