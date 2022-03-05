package main

import "fmt"

func main() {
	fmt.Println("Faktorial with loop :", factorialLoop(5))
	fmt.Println("Faktorial with recursive :", factorialRecursive(5))
}

var factorialLoop = func(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value <= 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
