package section12

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

func Warmup() {

	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	// item1 := item{id: 1, name: "god of war", price: 50}
	// fmt.Println(item1)

	// game1 := game{
	// 	item:  item{id: 1, name: "god of war", price: 50},
	// 	genre: "action adventure",
	// }
	// fmt.Println(game1)

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

	fmt.Println(games)
	for _, game := range games {
		fmt.Printf("id: %d name: %s price: %d genre: %s\n", game.id, game.name, game.price, game.genre)
	}

}
