package section4

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func HousingPricesAverages() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locas                                  []string
		sizes, beds, baths, prices             []int
		sumSizes, sumBeds, sumBaths, sumPrices int
		avgSizes, avgBeds, avgBaths, avgPrices float64
	)

	// Print Header
	headers := strings.Split(header, separator)

	for _, h := range headers {
		fmt.Printf("%-15s", h)

	}
	fmt.Println()

	// Print "==="
	fmt.Printf("%s\n", strings.Repeat("=", 80))

	rows := strings.Split(data, "\n")
	// fmt.Printf("%q\n", rows)

	for i := range rows {
		cols := strings.Split(rows[i], separator)

		locas = append(locas, cols[0])

		size, _ := strconv.Atoi(cols[1])
		sizes = append(sizes, size)
		sumSizes += size

		bed, _ := strconv.Atoi(cols[2])
		beds = append(beds, bed)
		sumBeds += bed

		bath, _ := strconv.Atoi(cols[3])
		baths = append(baths, bath)
		sumBaths += bath

		price, _ := strconv.Atoi(cols[4])
		prices = append(prices, price)
		sumPrices += price

	}

	// Print data
	for i := range rows {
		fmt.Printf("%-15s %-15d %-15d %-15d %-15d\n", locas[i], sizes[i], beds[i], baths[i], prices[i])
	}

	// Print "==="
	fmt.Printf("%s\n", strings.Repeat("=", 80))

	// Print average
	avgSizes = float64(sumSizes) / float64(len(rows))
	avgBeds = float64(sumBeds) / float64(len(rows))
	avgBaths = float64(sumBaths) / float64(len(rows))
	avgPrices = float64(sumPrices) / float64(len(rows))

	fmt.Printf("%-15s %-15.2f %-15.2f %-15.2f %-15.2f\n", "", avgSizes, avgBeds, avgBaths, avgPrices)

	// Solve this exercise by using your previous solution for
	// the "Housing Prices" exercise.
}

func HousingPricesAveragesb() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locs                       []string
		sizes, beds, baths, prices []int
	)

	rows := strings.Split(data, "\n")

	for _, row := range rows {
		cols := strings.Split(row, separator)

		locs = append(locs, cols[0])

		n, _ := strconv.Atoi(cols[1])
		sizes = append(sizes, n)

		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)

		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)

		n, _ = strconv.Atoi(cols[4])
		prices = append(prices, n)
	}

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range rows {
		fmt.Printf("%-15s", locs[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}

	// -------------------------------------------------------------------------
	// print the averages
	//
	// evil note:
	// later on, you will see how easy it would be to solve this
	// using the functions instead. there are a lot of repetitive code here.
	// -------------------------------------------------------------------------
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	// jump over the location column
	fmt.Printf("%-15s", "")

	var sum int

	for _, n := range sizes {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(sizes)))

	sum = 0
	for _, n := range beds {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(beds)))

	sum = 0
	for _, n := range baths {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(baths)))

	sum = 0
	for _, n := range prices {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(prices)))

	fmt.Println()
}
