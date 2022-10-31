package section11

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words
//
//  Create a program that prints the total and unique words
//  from an input stream.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Configure the scanner to scan for the words.
//
//  4. Count the unique words using a map.
//
//  5. Print the total and unique words.
//
//
// EXPECTED OUTPUT
//
//  There are 99 words, 70 of them are unique.
//
// ---------------------------------------------------------

func UniqueWords() {
	in := bufio.NewScanner(os.Stdin)

	words := map[string]string{}

	var n, u int

	for in.Scan() {
		content := strings.Fields(in.Text())

		n += len(content)

		for _, v := range content {
			words[v] = ""
		}

	}

	u = len(words)

	fmt.Printf("There are %d words, %d of them are unique.\n", n, u)

}

func UniqueWordsDA() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	total, words := 0, make(map[string]int)
	for in.Scan() {
		total++
		words[in.Text()]++
	}
	fmt.Printf("There are %d words, %d of them are unique.\n",
		total, len(words))
}
