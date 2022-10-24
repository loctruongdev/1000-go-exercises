package section4

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func ObserveLengthCapacity() {
	// --- #1 ---
	// 1. create a new slice named: games
	// var games []string

	// 2. print the length and capacity of the games slice
	// fmt.Printf("games's len: %d  games's cap: %d\n", len(games), cap(games))

	// 3. comment out the games slice
	//    then declare it as an empty slice
	// games := []string{}

	// 4. print the length and capacity of the games slice
	// fmt.Printf("games's len: %d  games's cap: %d\n", len(games), cap(games))

	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	// games = append(games, "pacman", "mario", "tetris", "doom")

	// 6. print the length and capacity of the games slice
	// fmt.Printf("games's len: %d  games's cap: %d\n", len(games), cap(games))

	// 7. comment out everything
	//
	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 5)
	games := []string{"pacman", "mario", "tetris", "doom"}
	fmt.Printf("games's len: %d  games's cap: %d\n", len(games), cap(games))
	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	// 2. print its length and capacity along the way (in the loop).

	fmt.Println()
	for i := 0; i <= len(games); i++ {
		s := games[:i]
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(s), cap(s))
	}

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	// 2. print the games and the new slice's len and cap
	//
	// 3. append a new element to the new slice
	//
	// 4. print the new slice's lens and caps
	//
	// 5. repeat the last two steps 5 times (use a loop)
	//
	// 6. notice the growth of the capacity after the 5th append
	//
	// Use this slice's elements to append to the new slice:
	// []string{"ultima", "dagger", "pong", "coldspot", "zetra"}
	fmt.Println()

	zero := games[:0]
	fmt.Printf("games's     len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	// zero = append(zero, "new elem")
	// fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	for _, v := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, v)
		fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	}

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	//
	// observe that, the range loop only loops for the length, not the cap.
	fmt.Println()

	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", n, len(s), cap(s))
	}

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	//
	// 2. print the elements of the zero slice in the loop.
	fmt.Println()

	zero = zero[:cap(zero)]
	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", n, len(s), cap(s), s)
	}

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	//
	// 2. change the same element of the games slice
	//
	// 3. print the games and the zero slices
	//
	// 4. observe that they don't have the same backing array
	fmt.Println()
	zero[2] = "myzero"
	games[2] = "mygame"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// try to slice the games slice beyond its capacity
	// uncomment and see the error.
	_ = games[:cap(games)+1]
	// or:
	// _ = games[:5]
}
