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
	"strings"
	"unicode/utf8"
)

func main() {
	// currently it prints 17
	// it should print 5

	name := "inanç           "
	fmt.Println(utf8.RuneCountInString(strings.TrimSpace(name)))
}
