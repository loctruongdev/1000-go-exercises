package section10

import (
	"fmt"
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Students
//
//  Create a program that returns the students by the given
//  Hogwarts house name (see the data below).
//
//  Print the students sorted by name.
//
//  "bobo" doesn't belong to Hogwarts, remove it by using
//  the delete function.
//
//
// RESTRICTIONS
//
//  + Add the following data to your map as is.
//    Do not sort it manually and do not modify it.
//
//  + Slices in the map shouldn't be sorted (changed).
//    HINT: Copy them.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//  Please type a Hogwarts house name.
//
//
//  go run main.go bobo
//
//  Sorry. I don't know anything about "bobo".
//
//
//  go run main.go hufflepuf
//
//  ~~~ hufflepuf students ~~~
//
//        + diggory
//        + helga
//        + scamander
//        + wenlock
//
// ---------------------------------------------------------

func Students() {
	// House        Student Name
	// ---------------------------

	students := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	houses := os.Args[1:]
	if len(houses) != 1 {
		fmt.Println("Please type a Hogwarts house name.")
		return
	}

	delete(students, "bobo")

	switch houses[0] {
	case "gryffindor", "hufflepuf", "ravenclaw", "slytherin", "bobo":
		// result := students[houses[0]] // -> // will sort students
		result := append([]string(nil), students[houses[0]]...) // -> new backing-array will not sort students

		sort.Strings(result)

		fmt.Printf("~~~ %s students ~~~\n", houses[0])
		for _, v := range result {
			fmt.Printf("\t+ %s\n", v)
		}
	default:
		fmt.Printf("Sorry. I don't know anything about %q.\n", houses[0])
		return
	}
	fmt.Println(students)

}

func StudentsDA() {
	houses := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory", "bobo"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "bobo", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	// remove the "bobo" house
	delete(houses, "bobo")

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please type a Hogwarts house name.")
		return
	}

	house, students := args[0], houses[args[0]]
	if students == nil {
		fmt.Printf("Sorry. I don't know anything about %q.\n", house)
		return
	}

	// only sort the clone
	clone := append([]string(nil), students...)
	sort.Strings(clone)

	fmt.Printf("~~~ %s students ~~~\n\n", house)
	for _, student := range clone {
		fmt.Printf("\t+ %s\n", student)
	}
}
