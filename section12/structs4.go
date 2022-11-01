package section12

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
// ---------------------------------------------------------

func Encode() {

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

	// Query by ID
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}
	fmt.Println(byID)

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : lists by id (ex: id 2)
  > save   : save game to json and quits
  > quit   : quits
`)
		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())

		if len(cmd) == 0 {
			fmt.Println("Please choose your options")
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}
		case "id":
			if len(cmd) != 2 {
				fmt.Println("Wrong ID")
				continue
			}

			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("Wrong ID")
				continue
			}
			// id : game -> byID := make(map[int]game
			g, ok := byID[id]
			if !ok {
				fmt.Println("sorry. I don't have the game")
				continue
			}
			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.id, g.name, "("+g.genre+")", g.price)
		case "save":
			type jsonGame struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Genre string `json:"genre"`
				Price int    `json:"price"`
			}

			var encodable []jsonGame
			for _, g := range games {
				encodable = append(encodable, jsonGame{g.id, g.name, g.genre, g.price})
			}

			out, err := json.MarshalIndent(encodable, "", "\t")
			if err != nil {
				fmt.Println("Sorry: ", err)
			}

			fmt.Println(string(out))
			return

		default:
			fmt.Println("Wrong ID")
			continue
		}
	}
}
