package main

import "fmt"

func main() {
	// map disini mirip/sama seperti HashMap di java atau asosiatif array di PHP
	// - cara membuat map
	// variable := map[keyType]valueType
	person := map[string]string{
		// add new data directly to map
		"name":    "irda islakhu afa",
		"address": "tuban",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}
