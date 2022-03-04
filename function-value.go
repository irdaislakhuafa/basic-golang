package main

import "fmt"

func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("irda"))
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
