package main

import "fmt"

func main() {
	// type declarations for create new alias for data type with existsing data type like string, int8, int16 etc
	type NoKTP string
	type isMarried bool

	// create new variable with new alias of data type above
	var noKtpIrda NoKTP = "156276575763"
	fmt.Println("No KTP Irda :", noKtpIrda)

	var irdaIsMaried isMarried = false
	fmt.Println("Sudah menikah :", irdaIsMaried)
}
