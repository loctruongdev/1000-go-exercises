package section2

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Crunch the primes
//
//  1. Get numbers from the command-line.
//  2. `for range` over them.
//  4. Convert each one to an int.
//  5. If one of the numbers isn't an int, just skip it.
//  6. Print the ones that are only the prime numbers.
//
// RESTRICTION
//  The user can run the program with any number of
//  arguments.
//
// HINT
//  Find whether a number is prime using this algorithm:
//  https://stackoverflow.com/a/1801446/115363
//
// EXPECTED OUTPUT
//  go run main.go 2 4 6 7 a 9 c 11 x 12 13
//    2 7 11 13
//
//  go run main.go 1 2 3 5 7 A B C
//    2 3 5 7
// ---------------------------------------------------------

func CrunchPrimes() {
	args := os.Args[1:]

loop:
	for _, n := range args {

		n, err := strconv.Atoi(n)

		if err != nil {
			continue loop
		} else {
			if isprime(n) {
				fmt.Printf("%-2d", n)
			}
		}

	}
	fmt.Println()

}

func isprime(n int) bool {
	// Returns True if n is prime.

	if n < 2 || n%2 == 0 {
		return false
	}

	if n == 2 {
		return true
	}

	if n == 3 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	if n%3 == 0 {
		return false
	}

	i := 5
	w := 2

	for i*i <= n {
		if n%i == 0 {
			return false
		}

		i += w
		w = 6 - w
	}

	return true
}

func CrunchPrimesb() {
	// remember [1:] skips the first argument

main:
	for _, arg := range os.Args[1:] {
		n, err := strconv.Atoi(arg)
		if err != nil {
			// skip non-numerics
			continue
		}

		switch {
		// prime
		case n == 2 || n == 3:
			fmt.Print(n, " ")
			continue

		// not a prime
		case n <= 1 || n%2 == 0 || n%3 == 0:
			continue
		}

		for i, w := 5, 2; i*i <= n; {
			// not a prime
			if n%i == 0 {
				continue main
			}

			i += w
			w = 6 - w
		}

		// all checks ok: it's a prime
		fmt.Print(n, " ")
	}

	fmt.Println()
}
