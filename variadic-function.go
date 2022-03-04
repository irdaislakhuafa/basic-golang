package main

import "fmt"

func main() {
	// variadic function parameter/varargs sama seperti RestParameter di Java dan js
	totalValue := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Total :", totalValue)

	// input slice to variadic function/RestParams
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Data :", sumAll(slice...))
}
func sumAll(numbers ...int) (total int) {
	total = 0
	for _, value := range numbers {
		total += value
	}
	return total
}
