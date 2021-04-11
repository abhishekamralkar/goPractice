package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPass  = "Invalid password %q.\n"
	accessOK = "Access given to %q.\n"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	user1, user2 := "anay", "vinayaka"
	pass1, pass2 := "2016", "2020"

	u, p := os.Args[1], os.Args[2]

	if u != user1 && u != user2 {
		fmt.Printf(errUser, u)
	} else if u == user1 && p != pass1 {
		fmt.Printf(accessOK, u)

	} else if u == user2 && p == pass2 {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(errPass, p)
	}

}
