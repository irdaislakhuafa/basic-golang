package main

import "fmt"

func main() {
	var name string

	name = "Irda Islakhu Afa"
	fmt.Println("My name is : ", name)

	name = "Ana Ardani"
	fmt.Println("My name is : ", name)

	// name = true -> error

	var friendName = "bambang"
	fmt.Println("My friend name is : ", friendName)

	var age = 20
	fmt.Println("Usia : ", age)

	country := "indonesia" // just for first declaration
	fmt.Println("Saya tinggal di : ", country, " haha")
	country = "palembang"
	fmt.Println("Saya tinggal di : ", country, " haha")

	// create two or more variable with go
	var (
		firstName = "irda"
		lastName  = "islakhu afa"
	)
	fmt.Println("Hello my name is ", firstName, lastName, "see you")
}
