package section2

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

const (
	usage5 = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.

The greater your number is, the harder it gets.

Wanna play?`

	maxTurn5 = 5
)

func EnoughPicks() {
	arg := os.Args[1:]

	if len(arg) < 1 {
		fmt.Printf(usage5, maxTurn5)
		return
	}

	guess, err := strconv.Atoi(arg[len(arg)-1]) //1 2 3 -> arg: {1, 2, 3} -> len(arg): 3 -> arg[2] = guess
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

	// var verbose bool
	// if arg[0] == "-v" {
	// 	verbose = true
	// }

	// Better solution: Avoid nested statements.

	for turn := 1; turn <= maxTurn5; turn++ {
		if guess < 10 {
			guess = 10
		}
		// fmt.Printf("%d ", guess)

		rn := rand.Intn(guess + 1)

		// if verbose {
		// 	fmt.Printf("%d ", rn)
		// }

		if rn != guess {
			continue
		}

		if turn == 1 {
			fmt.Println("🥇 FIRST TIME WINNER!!!")
		} else {
			fmt.Println("🎉  YOU WON!")
		}
		return
	}

	fmt.Println("☠️ YOU LOST... Try again?")

}
