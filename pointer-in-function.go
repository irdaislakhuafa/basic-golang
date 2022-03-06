package main

import (
	"fmt"
	"strings"
)

type DataCountry struct {
	Name string
}

func main() {
	defer func() {
		err := recover()
		switch {
		case err != nil:
			fmt.Println("terjadi error, Message :", err)

		}
	}()

	dataCountry := DataCountry{Name: "arab"}
	changeAddressToIndonesia(dataCountry) // call function to change name to indonesia
	fmt.Println("Data 1 :", dataCountry)  // tidak berubah

	println(strings.Repeat("-", 15))
	changeAddressToIndonesiaWithPointer(&dataCountry) // menggunakan pointer sebagai parameter / reference disarankan untuk menghemat memori
	fmt.Println("Data 1 :", dataCountry)              // data berubah
}
func changeAddressToIndonesia(data DataCountry) {
	data.Name = "indonesia"
}
func changeAddressToIndonesiaWithPointer(data *DataCountry) {
	data.Name = "indonesia"
}

// POINTER IN FUNTION :
/* sama seperti pointer di function di C/C++, jika pointer di function dirubah maka data asli juga berubah */
