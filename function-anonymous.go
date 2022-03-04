package main

import (
	"fmt"
	"strings"
)

type FunctionName func(string, string) bool

func registerUser(username, password string, blackList FunctionName) {
	if !blackList(username, password) {
		fmt.Println("you are blocked!", username)
	} else {
		fmt.Println("welcome", username)
	}
}

func main() {
	username := "irda"
	password := "anaardani"

	// cara 1
	blackList := func(name, pass string) bool {
		return (strings.ToLower(name) == "admin")
	}
	registerUser(username, password, blackList)
	fmt.Println(strings.Repeat("-", 15))

	// cara 2
	username = "admin"
	registerUser(username, password, func(username, password string) bool {
		return (strings.ToLower(username) == "admin")
	})
}
