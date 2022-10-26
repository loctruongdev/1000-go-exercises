package section6

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
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
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------

func SortWriteFfile() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return // don't forget this !!!
	}
	// fmt.Println(args)

	sort.Strings(args)
	// fmt.Println(args)

	data := make([]byte, 0, len(args))
	for _, v := range args {
		data = append(data, v...) //string == byte slice -> use v... = []byte...
		// data = append(data, "\n"...)
		data = append(data, '\n')
	}
	// fmt.Println(data)

	//https://chmod-calculator.com/ -> perms convert
	err := ioutil.WriteFile("sorted.txt", data, 0644)
	if err != nil {
		fmt.Println(err) // print the errors
		return
	}

}
