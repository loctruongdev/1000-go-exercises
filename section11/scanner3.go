package section11

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words 2
//
//  Use your solution from the previous "Unique Words"
//  exercise.
//
//  Before adding the words to your map, remove the
//  punctuation characters and numbers from them.
//
//
// BE CAREFUL
//
//  Now the shakespeare.txt contains upper and lower
//  case letters too.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < shakespeare.txt
//
//   There are 100 words, 69 of them are unique.
//
// ---------------------------------------------------------

func UniqueWords2() {
	// This is the regular expression pattern you need to use:
	// [^A-Za-z]+
	//
	// Matches to any character but upper case and lower case letters

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	rx := regexp.MustCompile(`[^A-Za-z]+`)

	total, words := 0, make(map[string]int)

	for in.Scan() {
		total++

		word := rx.ReplaceAllString(in.Text(), "")

		word = strings.ToLower(word)

		words[word]++
	}

	fmt.Printf("There are %d words, %d of them are unique.\n",
		total, len(words))

}
