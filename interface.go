package main

import (
	"fmt"
	"log"
	"strings"
)

// create interface (same as interface in java)
type ThisIsInterface interface { // sama seperti interface di java, jadi harus ada class implementasinya
	GetName() (string, string)
	WhoAreYou() string
}

// cara implementasi interface di Golang tidak seperti java tapi secara otomatis, jika ada struct (atau class di java) yang memiliki struktur sama seperti interface maka itu adalah class/struct implementasinya

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(strings.ToUpper("terjadi error!"))
		}
	}()

	var interface1 ThisIsInterface
	sayHelloToInterface(interface1)
}

func sayHelloToInterface(example ThisIsInterface) {
	fmt.Println("Hello, your name is", example.WhoAreYou())
}
