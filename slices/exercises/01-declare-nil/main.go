// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------

func main() {

	var names []string
	var distances []int
	var data []uint8
	var ratios []float64
	var alives []bool

	fmt.Printf("names : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances : %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives : %T %d %t\n", alives, len(alives), alives == nil)
}
