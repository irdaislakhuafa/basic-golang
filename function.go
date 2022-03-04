package main

import "fmt"

func main() {
	fmt.Println("\033\143")
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
func sayHello() {
	fmt.Println("Hello")
}
