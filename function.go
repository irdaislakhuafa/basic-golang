package main

import "fmt"

func main() {
	fmt.Println("\033\143")
	// for i := 0; i < 10; i++ {
	// 	sayHello(string(i))
	// }
	sayHello("irda", "islakhu afa")
}

// func sayHello() {
// 	fmt.Println("Hello")
// }
func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
