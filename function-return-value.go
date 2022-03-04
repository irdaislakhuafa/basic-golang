package main

import "fmt"

func main() {
	fmt.Println(sayHello("irda ", "islakhu afa"))
}

func sayHello(firstName string, lastName string) string {
	return "Hello \"" + firstName + " " + lastName + "\""
}
