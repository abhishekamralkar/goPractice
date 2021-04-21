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
	"time"
)

func main() {

	pToppings := []string{}
	dTimes := []time.Time{}
	gYears := []int{}
	sLight := []bool{}

	now := time.Now()
	pToppings = append(pToppings, "Panner", "Tomato", "Cheese")
	dTimes = append(dTimes,
		now,
		now.Add(time.Hour*24),
		now.Add(time.Hour*48))
	gYears = append(gYears, 2006, 2007, 2008)
	sLight = append(sLight, true, false, true)

	fmt.Printf("pizza : %s\n", pToppings)
	fmt.Printf("\ngraduations : %d\n", gYears)
	fmt.Printf("\ndepartures: %s\n", dTimes)
	fmt.Printf("\nlights : %t\n", sLight)

}
