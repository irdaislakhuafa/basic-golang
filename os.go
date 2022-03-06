package main

import (
	"fmt"
	"os" // untuk selebihnya bisa liat di dokumentasinya
)

func main() {
	// os.Args => sama seperti main(String.. args) di java
	fmt.Println("Args nya : ", os.Args)

	// mos.Hostname => sama seperti System.getProperty("user.name")
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", hostname)
	}

	fmt.Println("Enviroment : ", os.Getenv("GOPATH")) // mengambil enviroment variable
}

/*
PACKAGE OS : mirip package System di java
- Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi kita secara independen (bisa digunakan di sistem operasi apapun)
*/
