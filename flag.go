package main

import (
	"flag"
	"fmt"
)

var TestVar string

func main() {
	// flag.String(field, default value, description)
	flag.StringVar(&TestVar, "test-var", "ini test var", "put test var here!") // pake ini kalau variable udh di definisikan di tempat lain

	host := flag.String("host", "localhost", "Pust your host here!")
	username := flag.String("username", "root", "Pust your username here!")
	password := flag.String("password", "root", "Pust your password here!")

	// apply flag
	flag.Parse()

	fmt.Println("Host :", *host)
	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
	fmt.Println("Test Var :", TestVar)
}

/*
PACKAGE FLAG :
- contoh flag : command --help
- package flag berisikan fungsionalitas untuk memparsing command line agrs
- docs : https://golang.org/pkg/flag
*/
