package section2

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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

func PathSearcherb() {
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
