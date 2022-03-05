package main

import "fmt"

// create struct Customer (or class Customer if in java)

type Customer struct {
	Name, Address string // cat use one line
	Age           int
}

func main() {
	// create struct like object style
	customer1 := Customer{}
	customer1.Name = "Irda Islakhu Afa"
	customer1.Age = 20
	customer1.Address = "indonesia"

	// print customer
	fmt.Printf("customer1: %v\n", customer1)

	// create struct variable like inner constructor in java
	customer2 := Customer{
		Name:    "Ana Ardani",
		Age:     21,
		Address: "palembang",
	}
	fmt.Printf("customer2: %v\n", customer2)

	// create struct variable like constructor in java
	customer3 := Customer{"Ardani", "Palembang", 21} // tidak disarankan karena jika struktur berubah ini error
	fmt.Printf("customer3: %v\n", customer3)
}

/*
STRUCT : like class in java :D
*/
