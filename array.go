package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "irda"
	names[1] = "islakhu"
	names[2] = "afa"

	fmt.Println(names)
	fmt.Println(len(names))

	sheNames := [2]string{
		"Ana", "Ardani"
	}
	fmt.Println(sheNames)
}
