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
	"math"
)

func main() {
	var (
		radius = 10.
		area   float64
	)

	// ADD YOUR CODE HERE
	area = math.Pi * radius * radius

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)
}
