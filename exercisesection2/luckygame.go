package exercisessection2

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

const (
	usage1 = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.

The greater your number is, the harder it gets.

Wanna play?`

	maxTurn1 = 5
)

func FirstTurnWinner() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		fmt.Printf(usage1, maxTurn1)
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

	for turn := 1; turn <= maxTurn1; turn++ {
		rn := rand.Intn(guess + 1)
		// fmt.Printf("#%d: Lucky number = %d\n", turn, rn)

		if rn == guess {
			if turn == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU WON!")
			}
			return
		}
	}
	/*
		// Better solution: Avoid nested statements.

		for turn := 1; turn <= maxTurns; turn++ {
			n := rand.Intn(guess) + 1

			// Better, why?
			//
			// Instead of nesting the if statement into
			//   another if statement; it simply continues.
			//
			// TLDR: Avoid nested statements.
			if n != guess {
				continue
			}

			if turn == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU WON!")
			}
			return
		}
	*/

	fmt.Println("☠️ YOU LOST... Try again?")

}

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

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

const (
	maxTurn3 = 5 // less is more difficult
	usage3   = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
)

func DoubleGuesses() {
	args := os.Args[1:]

	if len(args) != 2 && len(args) != 1 {
		fmt.Printf(usage3, maxTurn3)
	} else if len(args) == 1 {
		fmt.Println("We need two numbers")
		return
	}

	guess, err := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])

	if err != nil || err2 != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess <= 0 || guess2 <= 0 {
		fmt.Println("Please pick positive numbers.")
		return
	}

	max := guess
	if guess < guess2 {
		max = guess2
	}

	rand.Seed(time.Now().UnixNano())

	for turn := 0; turn < maxTurn3; turn++ {
		n := rand.Intn(max + 1)
		fmt.Printf("Lucky number: %d \n", n)

		if n == guess || n == guess2 {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const (
	usage4 = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.

The greater your number is, the harder it gets.

Wanna play?`

	maxTurn4 = 5
)

func VerboseMode() {
	arg := os.Args[1:]

	if len(arg) < 1 {
		fmt.Printf(usage4, maxTurn4)
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

	var verbose bool
	if arg[0] == "-v" {
		verbose = true
	}

	for turn := 1; turn <= maxTurn4; turn++ {
		rn := rand.Intn(guess + 1)
		// fmt.Printf("#%d: Lucky number = %d\n", turn, rn)

		if verbose {
			fmt.Printf("%d ", rn)
		}

		if rn == guess {
			if turn == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU WON!")
			}
			return
		}
	}
	/*
		// Better solution: Avoid nested statements.

		for turn := 1; turn <= maxTurns; turn++ {
			n := rand.Intn(guess) + 1

			// Better, why?
			//
			// Instead of nesting the if statement into
			//   another if statement; it simply continues.
			//
			// TLDR: Avoid nested statements.
			if n != guess {
				continue
			}

			if turn == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU WON!")
			}
			return
		}
	*/

	fmt.Println("☠️ YOU LOST... Try again?")

}

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

func DynamicDiff() {
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
