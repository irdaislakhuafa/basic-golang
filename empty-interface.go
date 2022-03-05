package main

import (
	"fmt"
	"strings"
)

//  empty interface{} sama seperti Object di java (class tertinggi)
func thisIsFunction(i interface{}) interface{} {
	fmt.Println(strings.Repeat("-", 15))
	switch {
	case i == 1:
		return "ini string"
	case i == 2:
		return true
	default:
		return (i.(int) + 2)
	}
}
func main() {
	fmt.Println("\033\143")
	fmt.Println("Data :", thisIsFunction(1)) // car function

	var data1 int = thisIsFunction(3).(int) // convert interface to int
	fmt.Println("Data1 :", data1)           // call function

}
