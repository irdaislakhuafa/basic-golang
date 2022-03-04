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
	line(30)

	// access slice with for
	slice := []string{"irda", "islakhu", "afa"}
	for i := 0; i < len(slice); i++ {
		fmt.Println("data ke :", i, "adalah = ", slice[i])
	}
	line(30)

	// access slice with for range
	for index, value := range slice {
		fmt.Println("index ke", index, "nilainya :", value)
	}
	line(30)

	// access map with for range
	mapExmaple := map[string]string{
		"firstName":  "irda",
		"middleName": "islakhu",
		"lastName":   "afa",
	}

	for key, value := range mapExmaple {
		fmt.Printf("key : %s, value : %s\n", key, value)
	}
}

func line(lineLong int) {
	for line := 0; line < lineLong; line++ {
		fmt.Print("=")
	}
	fmt.Println()
}
