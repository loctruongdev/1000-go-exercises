package section2

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game too easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
// ---------------------------------------------------------

const (
	usage6 = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.

The greater your number is, the harder it gets.

Wanna play?`
	maxTurn6 = 5
)

func DynamicDifficulty() {
	arg := os.Args[1:]

	if len(arg) < 1 {
		fmt.Printf(usage6, maxTurn6)
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

	var balancer int
	if guess > 10 {
		balancer = guess / 10
	}
	// fmt.Printf("Max Turns: %d \n", maxTurn6+balancer)

	rand.Seed(time.Now().UnixNano()) // random number source

	var verbose bool
	if arg[0] == "-v" {
		verbose = true
	}

	// Better solution: Avoid nested statements.

	for turn := 1; turn <= maxTurn6+balancer; turn++ {
		if guess < 10 {
			guess = 10
		}
		// fmt.Printf("%d ", guess)

		rn := rand.Intn(guess + 1)

		if verbose {
			fmt.Printf("%d ", rn)
		}

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
