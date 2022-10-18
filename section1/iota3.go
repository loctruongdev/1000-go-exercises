package section1

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Iota Seasons
//
//  Use iota to initialize the season constants.
//
// HINT
//  You can change the order of the constants.
//
// EXPECTED OUTPUT
//  12 3 6 9
// ---------------------------------------------------------

func IotaSeasons() {
	// NOTE : You should remove all the initializers below
	//        first. Then use iota to fix it.
	const (
		Spring = 3 * (iota + 1) // 3*(0+1)
		Summer                  //3*(1+1)
		Fall                    //3*(2+1)
		Winter                  //3*(3+1)
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
