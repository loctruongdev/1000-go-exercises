package section12

import (
	"bufio"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------

func List() {
	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	games := []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{
			item:  item{id: 2, name: "x-com 2", price: 30},
			genre: "strategy",
		},
		{
			item:  item{id: 3, name: "minecraft", price: 20},
			genre: "sandbox",
		},
	}

	// for _, game := range games {
	// 	fmt.Printf("id: %d name: %s price: %d genre: %s\n", game.id, game.name, game.price, game.genre)
	// }

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : lists all the games
  > quit   : quits
`)
		if !in.Scan() {
			break
		}

		fmt.Println()

		switch in.Text() {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}
		}
	}
}
