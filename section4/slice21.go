package section4

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Correct the lyric
//
//  You have a slice that contains the words of Beatles'
//  legendary song: Yesterday. However, the order of the
//  words are incorrect.
//
// CURRENT OUTPUT
//
//  [all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
// EXPECTED OUTPUT
//
//  [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday]
//
//
// STEPS
//
//  INITIAL SLICE:
//    [all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
//
//  1. Prepend "yesterday" to the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
//
//  2. Put the words to the correct positions in the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday]
//
//
//  3. Print the `lyric` slice.
//
//
// BONUS
//
//   + Think about when does the append allocate a new backing array.
//
//   + Check whether your conclusions are correct.
//
//
// HINTS
//
//   If you get stuck, check out the hints.md file.
//
// ---------------------------------------------------------

func CorrectLyric() {
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)
	fmt.Printf("len1: %d cap1: %d\n", len(lyric), cap(lyric))

	//  1. Prepend "yesterday" to the `lyric` slice.
	lyric = append(lyric[:1], lyric...)
	fmt.Printf("len2: %d cap2: %d\n", len(lyric), cap(lyric))

	lyric = append(lyric[:0], "yesterday")[:23]
	fmt.Printf("len3: %d cap3: %d\n", len(lyric), cap(lyric))

	//  2. Put the words to the correct positions in the `lyric` slice.
	// [yesterday all my troubles seemed so far away | oh I believe in yesterday | now it looks as though they are here to stay]

	lyric = append(lyric, lyric[8:13]...)
	fmt.Printf("len4: %d cap4: %d\n", len(lyric), cap(lyric))

	lyric = append(lyric[:8], lyric[13:23]...)[:cap(lyric)]
	fmt.Printf("len5: %d cap5: %d\n", len(lyric), cap(lyric))

	lyric = append(lyric[:18], lyric[23:28]...)
	fmt.Printf("len6: %d cap6: %d\n", len(lyric), cap(lyric))

	fmt.Printf("%s\n", lyric)
}

func CorrectLyricDA() {
	// --- Correct Lyric ---
	// yesterday all my troubles seemed so far away
	// now it looks as though they are here to stay
	// oh I believe in yesterday
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)
	fmt.Printf("len1: %d cap1: %d\n", len(lyric), cap(lyric))
	// ------------------------------------------------------------------------
	// #1: Prepend "yesterday" to `lyric`
	// ------------------------------------------------------------------------

	//
	// --- BEFORE ---
	// all my troubles seemed so far away oh I believe in yesterday
	//
	// --- AFTER ---
	// yesterday all my troubles seemed so far away oh I believe in yesterday
	//
	// (warning: allocates a new backing array)
	//
	lyric = append([]string{"yesterday"}, lyric...)
	fmt.Printf("len2: %d cap2: %d\n", len(lyric), cap(lyric))
	// ------------------------------------------------------------------------
	// #2: Put the words to the correct positions in the `lyric` slice.
	// ------------------------------------------------------------------------

	//
	// yesterday all my troubles seemed so far away oh I believe in yesterday
	//                                              |               |
	//                                              v               v
	//                                       index: 8          pos: 13
	//
	const N, M = 8, 13

	// --- BEFORE ---
	//
	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay
	//
	// --- AFTER ---
	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay oh I believe in yesterday
	//                                              ^
	//
	//                                              |
	//                                              +--- duplicate
	//
	// (warning: allocates a new backing array)
	lyric = append(lyric, lyric[N:M]...)
	fmt.Printf("len3: %d cap3: %d\n", len(lyric), cap(lyric))
	//
	// --- BEFORE ---
	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay oh I believe in yesterday
	//
	// --- AFTER ---
	// yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday
	//
	// (does not allocate a new backing array because cap(lyric[:N]) > M)
	//
	lyric = append(lyric[:N], lyric[M:]...)
	fmt.Printf("len4: %d cap4: %d\n", len(lyric), cap(lyric))
	// ------------------------------------------------------------------------
	// #3: Print
	// ------------------------------------------------------------------------
	fmt.Printf("%s\n", lyric)
}
