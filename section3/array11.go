package section3

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Find the Average
//
//  Your goal is to fill an array with numbers and find the average.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Print the given numbers and their average.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// EXPECTED OUTPUT
//   go run main.go
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5 6
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5
//     Your numbers: [1 2 3 4 5]
//     Average: 3
//
//   go run main.go 1 a 2 b 3
//     Your numbers: [1 0 2 0 3]
//     Average: 2
// ---------------------------------------------------------

func FindAverage() {
	args := os.Args[1:]

	// fmt.Println(len(args))

	if len(args) < 1 || len(args) > 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}

	// numbers := [5]float64{}
	// var sum, l float64
	var (
		numbers [5]float64
		sum     float64
		length  float64
	)

	for i := 0; i < len(args); i++ {

		if n, err := strconv.ParseFloat(args[i], 64); err == nil {
			numbers[i] = n
			sum += n
			length += 1
		}
	}
	fmt.Printf("Your numbers: %v\nAverage: %v\n", numbers, sum/length)

}

func FindAverageb() {
	args := os.Args[1:]
	if l := len(args); l == 0 || l > 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}

	var (
		sum   float64
		nums  [5]float64
		total float64
	)

	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}

		total++
		nums[i] = n
		sum += n
	}

	fmt.Println("Your numbers:", nums)
	fmt.Println("Average:", sum/total)
}
