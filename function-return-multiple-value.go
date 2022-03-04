package main

import "fmt"

func main() {
	// firstName, lastName := getFullName()

	// if want to ignore result use _
	firstName, _ := getFullName()

	// fmt.Println(firstName, lastName)
	fmt.Println(firstName)
	fmt.Println(getFullName())
}

// function with return multiple value
func getFullName() (string, string) {
	return "irda", "islakhu afa"
}
