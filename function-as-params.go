package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple Shell")
	fmt.Println(strings.Repeat("-", 30))
	// var shellInput, error string

	for true {
		fmt.Print("-> ")
		shellInput, _ := reader.ReadString('\n')
		shellInput = strings.Replace(shellInput, "\n", "", -1)

		switch {
		case strings.ToLower(shellInput) == "hi":
			fmt.Println("hi too :D")
		case strings.Contains(strings.ToLower(shellInput), "anjing"):
			// fmt.Println(sayHello(shellInput, keywordFilter))

			// or
			filter := keywordFilter
			fmt.Println(sayHello(shellInput, filter))
		default:
			continue
		}
	}
}

func sayHello(name string, filterKeyword func(string) string) string {
	return "Hello " + filterKeyword(name)
}

func keywordFilter(message string) string {
	lowerCaseMessage := strings.ToLower(message)
	if lowerCaseMessage == "anjing" || strings.Contains(lowerCaseMessage, "anjing") {
		return "..."
	} else {
		return message
	}
}
