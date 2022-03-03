package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var result = a + b

	fmt.Println("10 + 10 =", result)
	fmt.Println("10 * 10 =", (a * b))

	// augmented assignment
	// a = a + 1 ---> a+= 1
	a += 12
	println("value A :", a)
	a++
	println("value A :", a)
}
