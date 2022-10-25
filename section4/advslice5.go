package section4

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Print daily requests
//
//  You've got request logs of a web server. The log data
//  contains 8-hourly totals per each day. It is stored
//  in the `reqs` slice.
//
//  Find and print the total requests per day, as well as
//  the grand total.
//
//  See the `reqs` slice and the steps in the code below.
//
//
// RESTRICTIONS
//
//   1. You need to produce the daily slice, don't just loop
//      and print the element totals directly. The goal is
//      gaining more experience in slice operations.
//
//   2. Your code should work even if you add to or remove the
//      existing elements from the `reqs` slice.
//
//      For example, after solving the exercise, try it with
//      this new data:
//
//      reqs := []int{
// 	      500, 600, 250,
// 	      200, 400, 50,
// 	      900, 800, 600,
// 	      750, 250, 100,
// 	      150, 654, 235,
// 	      320, 534, 765,
// 	      121, 876, 285,
// 	      543, 642,
// 	      // the last element is missing (your code should be able to handle this)
// 	      // that is why you shouldn't calculate the `size` below manually.
//      }
//
//      The grand total of the new data should be 10525.
//
//
// EXPECTED OUTPUT
//
//   Please run `solution/main.go` to see the expected
//   output.
//
//   go run solution/main.go
//
// ---------------------------------------------------------

func PrintDailyRequest() {
	// There are 3 request totals per day (8-hourly)
	const N = 3

	// DAILY REQUESTS DATA (8-HOURLY TOTALS PER DAY)
	reqs := []int{
		500, 600, 250, // 1st day: 1350 requests
		200, 400, 50, // 2nd day: 650 requests
		900, 800, 600, // 3rd day: 2300 requests
		750, 250, 100, // 4th day: 1100 requests
		999,
		// grand total: 5400 requests
	}

	// ================================================
	// #1: Make a new slice with the exact size needed.

	l := len(reqs)
	// fmt.Println(l)

	size := l / N // you need to find the size.
	if l%N != 0 {
		size++
	}
	// fmt.Println(size)

	daily := make([][]int, 0, size)

	// ================================================

	// ================================================
	// #2: Group the `reqs` per day into the slice: `daily`.
	//
	// So the daily will be:
	// [
	//  [500, 600, 250]
	//  [200, 400, 50]
	//  [900, 800, 600]
	//  [750, 250, 100]
	// ]

	for N <= len(reqs) {
		daily = append(daily, reqs[:N])
		reqs = reqs[N:]
	}
	// fmt.Println(daily)
	// fmt.Println(reqs)

	daily = append(daily, reqs)
	// fmt.Println(daily)
	// ================================================
	// #3: Print the results

	// Print a header
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	// Loop over the daily slice and its inner slices to find
	// the daily totals and the grand total.
	var total int
	for i, day := range daily {

		var sum int

		for _, v := range day {
			sum += v

		}
		fmt.Printf("%-15d %-15d\n", i+1, sum)
		total += sum

	}
	fmt.Printf("%-15s %-15d\n", "total", total)
	// ================================================
}
func PrintDailyRequestDA() {
	// There are 3 request totals per day
	const N = 3

	// DAILY REQUESTS DATA (PER 8-HOUR)
	reqs := []int{
		500, 600, 250, // 1st day: 1350 requests
		200, 400, 50, // 2nd day: 650 requests
		900, 800, 600, // 3rd day: 2300 requests
		750, 250, 100, // 4th day: 1100 requests
		// grand total: 5400 requests
	}

	// ALSO TRY IT WITH THIS DATA:
	// reqs = []int{
	// 	500, 600, 250,
	// 	200, 400, 50,
	// 	900, 800, 600,
	// 	750, 250, 100,
	// 	150, 654, 235,
	// 	320, 534, 765,
	// 	121, 876, 285,
	// 	543, 642,
	// }

	// ================================================
	// Allocate a slice efficiently with the exact size needed.
	//
	// Find the `size` automatically using the `reqs` slice.
	// Do not type it manually.
	//
	l := len(reqs)
	size := l / N
	if l%N != 0 {
		size++
	}
	daily := make([][]int, 0, size)
	//
	// ================================================

	// ================================================
	// Group the `reqs` per day into the slice: `daily`.
	//
	// So the daily will be:
	// [
	//  [500, 600, 250]
	//  [200, 400, 50]
	//  [900, 800, 600]
	//  [750, 250, 100]
	// ]
	//
	for N <= len(reqs) {
		daily = append(daily, reqs[:N]) // append the daily requests
		reqs = reqs[N:]                 // move the slice pointer for the next day
	}

	// add the residual data
	if len(reqs) > 0 {
		daily = append(daily, reqs)
	}

	// ================================================

	// ================================================
	// Print the header
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	// Print the data per day along with the totals
	var grand int

	for i, day := range daily {
		var sum int

		for _, req := range day {
			sum += req
			fmt.Printf("%-10d%-10d\n", i+1, req)
		}

		fmt.Printf("%9s %-10d\n\n", "TOTAL:", sum)

		grand += sum
	}

	fmt.Printf("%9s %-10d\n", "GRAND:", grand)
	// ================================================
}
