package exercisesection2b

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const corpus = "" + "lazy cat jumps again and again and again"

func WordFinder() {
	words := strings.Fields(corpus)

	query := os.Args[1:]

	for _, q := range query {
		// fmt.Println(q)
		for i, w := range words {
			if w == strings.ToLower(q) {
				fmt.Printf("#%d: %q\n", i+1, w)

				break
			}
		}
	}

}

// ---------------------------------------------------------
// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.
// ---------------------------------------------------------

const corpus2 = "lazy cat jumps again and again and again"

func CaseInsensitiveSearch() {
	words := strings.Fields(corpus2)
	query := os.Args[1:]

queries:
	for _, q := range query {
		// case insensitive search
		q = strings.ToLower(q)

	search:
		for i, w := range words {
			switch q {
			case "and", "or", "the": // filter
				break search
			}

			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				continue queries
			}
		}
	}
}

// ---------------------------------------------------------
// EXERCISE: Path Searcher
//
//  Your program should search inside the path environment
//  variable.
//
//  Remove the corpus constant then get the corpus from the
//  environment variable "Path" or "PATH".
//
// HINTS
//  1. Search the web to find out what is an environment
//     variable and how to use it (if you don't know
//     what it is already).
//
//  2. Look up for the necessary functions for getting
//     an environment variable. It's in the "os" package.
//
//     Search for it on the Go online documentation.
//
//  3. Look up for the necessary function for splitting
//     the path variable into directory names.
//
//     It's in the "path/filepath" package.
//
// EXAMPLE
//  For example, on my Mac, my PATH environment variable
//  looks like this:
//
//    "/usr/local/bin:/sbin:/Users/inanc/go/bin"
//
//  So, if the user runs the program like this:
//
//    go run main.go /sbin
//
//  It should print this:
//
//    #2 : "/sbin"
// ---------------------------------------------------------

// ---------------------------------------------------------
// BONUS EXERCISE
//  Make your program cross platform. So, it can search
//  the path environment variable when you run it on
//  a Windows or on a Mac (OS X) or on a Linux.
//
// BONUS EXERCISE #2
//  Also find the directories for the directories that
//  contains the searched query.
//
//  And let it match in a case-insensitive manner. See the examples.
//
//  EXAMPLE:
//    Let's say:
//      PATH="/usr/local/bin:/sbin:/Users/inanc/go/bin"
//
//  So, if the user runs the program like this:
//    go run main.go bin
//
//  It should print this:
//    #1 : "/usr/local/bin"
//    #2 : "/sbin"
//    #3 : "/Users/inanc/go/bin"
//
//  Or like this (case insensitive):
//    go run main.go Us
//
//  Output:
//    #1 : "/usr/local/bin"
//    #2 : "/Users/inanc/go/bin"
// ---------------------------------------------------------

func PathSearcher() {
	// Get and split the PATH environment variable

	// SplitList function automatically finds the
	// separator for the path env variable

	words := filepath.SplitList(os.Getenv("PATH"))

	// words2 := strings.Split(os.Getenv("PATH"), ":") // Alternative way

	// Alternative way, but above one is better:
	// words := strings.Split(
	// 	os.Getenv("PATH"),
	// 	string(os.PathListSeparator))

	query := os.Args[1:]

	for _, q := range query {
		for i, w := range words {
			q, w = strings.ToLower(q), strings.ToLower(w)
			// fmt.Printf("q: %q | w: %q\n", q, w)

			if strings.Contains(w, q) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
			}
		}
	}
}

func PathSearcher2() {
	path := os.Getenv("PATH")
	// fmt.Printf("Type: %T\nValue: %v\n", path, path)

	sPath := filepath.SplitList(path)
	// fmt.Printf("Type: %T\nValue: %v\n", sPath, sPath)

	query := os.Args[1:]
	// fmt.Printf("Type: %T\nValue: %v\n", query, query)

	for _, q := range query {
		for i, p := range sPath {
			q, p := strings.ToLower(q), strings.ToLower(p)
			// fmt.Printf("q: %q, p: %q\n", q, p)

			// if q == p {
			// 	fmt.Printf("#%-2d: %q\n", i+1, q)
			// }

			if strings.Contains(p, q) {
				fmt.Printf("#%-2d: %q\n", i+1, q)
			}
		}
	}

}

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

func CrunchPrimes2() {
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
