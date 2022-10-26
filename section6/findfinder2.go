package section6

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file with their ordinals
//
//  Use the previous exercise.
//
//  This time, print the sorted items with ordinals
//  (see the expected output)
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     1. apple
//     2. banana
//     3. orange
//
//
// HINTS
//
//   ONLY READ THIS IF YOU GET STUCK
//
//   + You can use strconv.AppendInt function to append an int
//     to a byte slice. strconv contains a lot of functions for appending
//     other basic types to []byte slices as well.
//
//   + You can append individual characters to a byte slice using
//     rune literals (because: rune literal are typeless numerics):
//
//     var slice []byte
//     slice = append(slice, 'h', 'i', ' ', '!')
//     fmt.Printf("%s\n", slice)
//
//     Above code prints: hi !
// ---------------------------------------------------------

func SortWriteFfile2() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return // don't forget this !!!
	}
	// fmt.Println(args)

	sort.Strings(args)
	// fmt.Println(args)

	data := make([]byte, 0, len(args))
	for i, v := range args {
		// j := strconv.Itoa(i+1) + "." + " "
		// data = append(data, j...)
		data = strconv.AppendInt(data, int64(i+1), 10)
		data = append(data, '.', ' ')

		data = append(data, v...) //string == byte slice -> use v... = []byte...

		// data = append(data, "\n"...)
		data = append(data, '\n')

	}
	// fmt.Println(data)

	//chmod-calculator.com/ -> perms convert
	err := ioutil.WriteFile("sorted.txt", data, 0644)
	if err != nil {
		fmt.Println(err) // print the errors
		return
	}

}
