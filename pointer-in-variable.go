package main

import (
	"fmt"
	"strings"
)

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Tuban", "Jawa Timur", "Indonesia"}
	address2 := address1 // ini past by value bukan past by reference (default)

	address2.City = "palembang"
	fmt.Println("Address 1 :", address1) // address1 tidak berubah jika address2 di ubah
	fmt.Println("Address 2 :", address2) // jadi dia past by value / duplicate
	fmt.Println(strings.Repeat("-", 15))

	// menggunakan past by reference
	// var address3 *Address = &address1
	address3 := &address1 // sama aja kek di atas

	// meassign ulang value address3
	// address3 = &Address{"palembang", "Jawa Timur", "Indonesia"} // dia berada di memori lain

	address3.Province = "Sumatra Selatan"
	fmt.Println("Address 3 :", address3)
	fmt.Println("Address 1 :", address1)

	// memaksa semua yang mereference ke reference address3 berubah alamat memorinya (skaligus datanya)
	*address3 = Address{"anu", "Jawa Timur", "Indonesia"} // ketika merubah data address3 maka semua variable yg mereference ke address1 juga akan berubah
	// coba buat address baru yg mereference ke addres2
	address4 := &address1
	address4.Country = "majapahit" // semua yang mereference ke address 1 juga berubah

	fmt.Println(strings.Repeat("-", 15))
	fmt.Println("Address 3 :", address3)
	fmt.Println("Address 1 :", address1)
	fmt.Println("Address 4 :", address4)

	// function new, sama seperti & yaitu membuat pointer dari data yg sudah ada kalau function new() mengembalikan pointer kosong
	address5 := new(Address)
	address5.City = "palembang" //
	fmt.Println("Address 5 :", address5)
}
func chngeCityWithoutPointer() {

}

/*
 pointer di Go sama seperti pointer di C/C++
*/
