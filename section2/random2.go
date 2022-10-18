package section2

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

const (
	usage2 = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.

The greater your number is, the harder it gets.

Wanna play?`

	maxTurn2 = 5
)

func RandomMessages() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		fmt.Printf(usage2, maxTurn2)
		return
	}

	guess, err := strconv.Atoi(arg[0])
	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if guess <= 0 { // can not be 0 becaue of always matching
		fmt.Println("Give me a positive number")
		return
	}
	// fmt.Printf("Your number: %d.\n", guess)

	rand.Seed(time.Now().UnixNano()) // random number source

	for turn := 1; turn <= maxTurn2; turn++ {
		n := rand.Intn(guess + 1)

		if n != guess {
			continue
		}

		c := rand.Intn(3)
		switch c {
		case 0:
			fmt.Println("YOU wIN! 111")
		case 1:
			fmt.Println("YOU wIN! 222")
		case 2:
			fmt.Println("YOU wIN! 333")
		}
		return
	}
	c := rand.Intn(3)
	switch c {
	case 0:
		fmt.Println("YOU LOST... Try again? 111")
	case 1:
		fmt.Println("YOU LOST... Try again? 222")
	case 2:
		fmt.Println("YOU LOST... Try again? 333")
	}

}
