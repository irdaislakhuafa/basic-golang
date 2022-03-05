package main

import (
	"fmt"
	"strings"
)

func main() {
	println("\033\143")
	runApp(false)
	runApp(true)
}
func runApp(err bool) {
	defer endApp() // call function with defer

	println("\nStarting...")

	if err {
		panic(strings.ToUpper("Terjadi error!"))
	}
	fmt.Println("Aplikasi berjalan:D")
}
func endApp() {
	fmt.Println("Ending start app")
}
