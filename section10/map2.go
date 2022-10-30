package section10

import (
	"fmt"
)

func PopulateLookup() {
	// ---------------------------------------------------------
	// EXERCISE: Populate and Lookup
	//
	//  Add elements to the maps that you've declared in the
	//  first exercise, and try them by looking up for the keys.
	//
	//  Either use the `make()` or `map literals`.
	//
	//  After completing the exercise, remove the data and check
	//  that your program still works.
	//
	//
	//  1. Phone numbers by last name
	//     --------------------------
	//     bowen  202-555-0179
	//     dulin  03.37.77.63.06
	//     greco  03489940240
	//
	//     Print the dulin's phone number.
	//
	//
	//  2. Product availability by Product ID
	//     ----------------
	//     617841573 true
	//     879401371 false
	//     576872813 true
	//
	//     Is Product ID 879401371 available?
	//
	//
	//  3. Multiple phone numbers by last name
	//     ------------------------------------------------------
	//     bowen  [202-555-0179]
	//     dulin  [03.37.77.63.06 03.37.70.50.05 02.20.40.10.04]
	//     greco  [03489940240 03489900120]
	//
	//     What is Greco's second phone number?
	//
	//
	//  4. Shopping basket by Customer ID
	//     -------------------------------
	//     100 [617841573:4 576872813:2]
	//     101 [576872813:5 657473833:20]
	//     102 []
	//
	//     How many of 576872813 the customer 101 is going to buy?
	//                (Product ID)  (Customer ID)
	//
	//
	// EXPECTED OUTPUT
	//
	//   1. Run the solution to see the output
	//   2. Here is the output with empty maps:
	//
	//      dulin's phone number: N/A
	//      Product ID #879401371 is not available
	//      greco's 2nd phone number: N/A
	//      Customer #101 is going to buy 5 from Product ID #576872813.
	//
	// ---------------------------------------------------------
	// phones := map[string]string{
	// 	"bowen": "202-555-0179",
	// 	"dulin": "03.37.77.63.06",
	// 	"greco": "03489940240",
	// }

	// args := os.Args[1:]

	// who, phone := args[0], "N/A"
	// if v, ok := phones[who]; ok {
	// 	phone = v
	// }
	// fmt.Printf("%s's phone number: %s\n", who, phone)

	product := map[int]bool{
		617841573: true,
		879401371: false,
		576872813: true,
	}

	id, status := 617841573, "available"

	if !product[id] {
		status = "not " + status
	}
	fmt.Printf("Product ID %d is %s\n", id, status)

	multiPhones := map[string][]string{
		"bowen": {"202-555-0179"},
		"dulin": {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": {"03489940240", "03489900120"},
	}

	// greco := multiPhones["greco"][1]
	// fmt.Println(greco)

	who, phone := "greco", "N/A"
	if phones := multiPhones[who]; len(phones) >= 2 {
		phone = phones[1]
	}
	fmt.Printf("%s's phone number is: %s\n", who, phone)

	//     How many of 576872813 the customer 101 is going to buy?
	//                (Product ID)  (Customer ID)

	basket := map[int]map[int]int{
		100: {617841573: 4, 576872813: 2},
		101: {576872813: 5, 657473833: 20},
		102: {},
	}

	result := basket[101][576872813]
	fmt.Println(result)

	cid, pid := 101, 576872813
	solds := basket[cid][pid]
	fmt.Printf("Customer #%d is going to buy %d from Product ID #%d\n", cid, solds, pid)
}
