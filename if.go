package main

import "fmt"

func main() {
	name := "ana   "

	if name == "irda" {
		fmt.Println("Hello", name)
	} else if name == "ana" {
		fmt.Println("Hallo", name)
	} else {
		fmt.Println("Anda bukan", name)
	}

	// if with sort statement
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	}
}
