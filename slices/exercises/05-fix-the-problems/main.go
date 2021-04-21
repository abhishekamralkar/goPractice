// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"sort"
)

func main() {
	var names []string
	names = []string{}
	names = []string{
		"Einstein",
		"Shepard",
		"Tesla",
	}

	// -----------------------------------
	var books []string = []string{
		"Stay Golden",
		"Fire",
		"Kafka's Revenge",
	}

	sort.Strings(books)

	// -----------------------------------
	// // this time, do not change the nums array to a slice
	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}

	// // use the slicing expression to change the nums array to a slice below
	sort.Ints(nums[:])

	// -----------------------------------
	// Here: Use the strings.Join function to join the names
	//       (see the expected output)
	fmt.Printf("%q\n", names)

	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}
