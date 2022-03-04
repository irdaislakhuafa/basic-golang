package main

import (
	"fmt"
	"strings"
)

// closure mirip scope di java
func main() {
	counter := 0
	name := "irda"

	increment := func() {
		fmt.Println("Increment...")
		counter++

		// misal lupa ini
		// name := "ana"
		name = "ana"
		fmt.Println("She name :", name)
	}
	increment()
	increment()

	println()
	fmt.Println(strings.ToTitle("data counter :"), counter)
	fmt.Println(strings.ToTitle("data name :"), name)
}
