package main

import "fmt"

func main() {
	counter := 1

	// for in golang
	for counter <= 10 {
		fmt.Println("perulangan ke :", counter)
		counter++
	}
	line(30)

	// for with statement
	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke :", (i + 1))
	}
}

func line(lineLong int) {
	for line := 0; line < lineLong; line++ {
		fmt.Print("=")
	}
	fmt.Println()
}
