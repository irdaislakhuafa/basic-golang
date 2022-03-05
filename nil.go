package main

import (
	"fmt"
	"strings"
)

// nil = null in java
func main() {
	fmt.Println("\033\143")
	var person map[string]string = nil

	fmt.Println("Person :", person)

	// call function createNewMap
	// data 1
	data1 := createNewMap("irda")
	fmt.Println("Data 1 :", data1)

	// data 2
	data2 := createNewMap("")
	if data2 == nil {
		fmt.Println("Data kosong!")
	} else {
		fmt.Println("Data 2 :", data2)
	}
}
func createNewMap(name string) map[interface{}]string {
	name = strings.Trim(name, " ")
	switch {
	case name == "":
		return nil
	default:
		return map[interface{}]string{
			"name": name,
		}
	}
}
