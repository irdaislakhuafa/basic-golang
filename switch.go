package main

import "fmt"

func main() {
	// sama aja kyk swith di bahasa lain
	name := "ana"

	switch name {
	case "ana":
		fmt.Println("Hallo", name)
	case "irda":
		fmt.Println("Hallo", name)
	default:
		fmt.Println("anda siapa :v")
	}

	// switch with sort statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("Nama benar")
	}

	// switch with simple expression(like if else)
	switch {
	case len(name) > 10:
		fmt.Println("huruf terlalu panjang")
	case len(name) > 5:
		fmt.Println("huruf lumayan panjang")
	default:
		fmt.Println("huruf sudah benar")
	}
}
