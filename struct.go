package main

import (
	"fmt"
	"strings"
)

// create struct Customer (or class Customer if in java)

type Customer struct {
	Name, Address string // cat use one line
	Age           int
}

func main() {
	println("\033\143")
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

	fmt.Sprintln(strings.Repeat("-", 15))
	customer1.sayHello("hahahaha")
	customer1.saySomething()
}

func (customer Customer) sayHello(name ...string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Hello, %s name is %s\n", "kawan", customer.Name)
		}
	}()

	fmt.Printf("Hello, %s name is %s\n", name[0], customer.Name)
}
func (customer Customer) saySomething() {
	fmt.Printf("Hello, i want to say SOMETHING :D\n") // cringe :'
}

/*
STRUCT : like class in java :D
STRUCT FUNCTION : mirip method di class java hanya saja ini sebenernya bukan milik struct tapi seakan akan milik struct
*/
