package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\033\143")
	myName := "-> Irda Islakhu Afa "
	fmt.Println(strings.Contains(myName, "irda"))
	fmt.Println(strings.Contains(myName, "ana"))
	CreateLine()

	fmt.Println(strings.Split(myName, " ")[0]) // return array of string
	CreateLine()

	fmt.Println(strings.ToLower(myName))
	CreateLine()

	fmt.Println(strings.ToUpper(myName))
	CreateLine()

	fmt.Println(strings.ReplaceAll(strings.ToLower(myName), "i", "a"))
	CreateLine()

	fmt.Println(strings.Replace(strings.ToLower(myName), "i", "a", 1 /*  berapa jumlah yang di ubah */))
	CreateLine()

	fmt.Println(strings.Trim(myName, "-> "))
	CreateLine()
}
func CreateLine() {
	fmt.Println(strings.Repeat("-", 20))
}

/*
PACKAGE STRINGS	 :
- package yang berisikan function function untuk memanipulasi tipe data string
- docs : https://golang.org/pkg/strings

BEBERAPA METHOD DI STRINGS
- strings.Trim(string, cutset) -> memotong cutset diawal dan diakhir string
- strings.ToLower(string) -> convert string ke lowercase
- strings.ToUpper(string) -> convert string ke uppercase
- strings.Split(string, separator) -> memotong string berdasarkan separator
- strings.Contains(string, search) -> mengecek apakah string mengandung string lain
- strings.ReplaceAll(string, from, to) -> mengubah semua string dari from ke to
*/
