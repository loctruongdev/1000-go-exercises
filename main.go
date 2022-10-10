package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path: ", os.Args[0])
	fmt.Println("1st argument: ", os.Args[1])
	fmt.Println("2nd argument: ", os.Args[2])
	fmt.Println("3rd argument: ", os.Args[3])

	fmt.Println("Number of items indside os.Args: ", len(os.Args))
}
