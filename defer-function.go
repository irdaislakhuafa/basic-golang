package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\033\143")
	runApp(0)
}

func runApp(value int) {
	defer logging() // wajib di panggil di atas agas di eksekusi setelah function runApp()
	fmt.Println("RunApp function")

	result := 10 / value // disini error
	fmt.Println("value :", result)

	// logging() // => akan dijalankan, tapi dika terjadi error di atasnya maka tidak akan dijalankan
}

func logging() {
	fmt.Println(strings.Repeat("-", 15))
	fmt.Println("Selesai memanggil function...")
}

/*
DEFER FUNCTION :
function yang akan dieksekusi setelah sebuah function di eksekusi, walaupun function yang memanggil itu memiliki error, hampir mirip try catch.
*/
