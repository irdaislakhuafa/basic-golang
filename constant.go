package main

import "fmt"

func main() {
	// create constant variable
	const firstName string = "irda"
	const lastName = "islakhu afa"
	const value = 1000

	fmt.Println("Nama :", firstName)

	// create multiple constant
	const (
		sheName = "Ana Ardani"
		myName  = "Irda Islakhu Afa"
	)
	fmt.Println(sheName)
	fmt.Println(myName)
}
