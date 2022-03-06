package helpers

import (
	"fmt"
	"strings"
)

var appVersion = "1.2.3"
var MyApplicationName = "Belajar Golang"

func SayHelloFromHelper() {
	fmt.Println("Hello :D")
}

func sayHelloFromHelper() {
	fmt.Println("ini say hello yang private")
}
func Line(longLine int, character ...*string) {
	var result string
	if len(character) > 0 {
		result = *character[0]
	} else {
		result = "-"
	}
	fmt.Println(strings.Repeat(result, longLine))
}

/*
PACKAGE : mirip seperti package di java, dan harus membuat folder yang sama sengan stuctur package
- adalah tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di Go-Lang
- di 1 package tidak boleh ada function/struct yang sama
*/
