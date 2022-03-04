package main

import "fmt"

func main() {
	firstName, middleName, lastName := getFullName("irda", "islakhu", "afa")
	fmt.Println(firstName, middleName, lastName)
}
func getFullName(first, middle, last string) (firstName, middleName, lastName string) {
	firstName = first
	middleName = middle
	lastName = last
	return
}
