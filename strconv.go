package main

import (
	"fmt"
	"strconv"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Terjadi Error")
		}
	}()

	booleanValue, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("Boolean :", booleanValue)
	} else {
		fmt.Println("Boolean :", err.Error())
	}

	numberValue, err := strconv.ParseInt("1000", 10, 64)
	/*
		prefix following the sign (if present): 2 for "0b", 8 for "0" or "0o",
		16 for "0x", and 10 otherwise. Also, for argument base 0 only,
		underscore characters are permitted as defined by the Go syntax for
	*/
	if err != nil {
		fmt.Println("Number :", err.Error())
	} else {
		fmt.Println("Number :", numberValue)
	}

	stringNumber := strconv.FormatInt(10_000, 10)
	fmt.Println("String :", stringNumber)
}

/*
PACKAGE STRCONV :
- sebelumnya kita sudah belajar cara konversi tipe data, misal int8 ke int32, atau yang lainnya yang tipenya sama
- bagaimana jika kita butuh melakukan konversi yang tipe data yang berbeda. Misal dari int ke string atau sebaliknya
- hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
- docs : https://golang.org/pkg/strconv

BEBERAPA LIST METHOD STRCONV
- strconv.ParseInt(string) -> convert string ke int
- strconv.ParseBool(string) -> convert string ke bool
- strconv.ParseFloat(string) -> convert string ke float
- strconv.FormatInt(int) -> convert int ke string
- strconv.FormatFloat(float) -> convert float ke string
- strconv.FormatBool(bool) -> convert bool ke string
- strconv.Atoi(string) -> convert string ke int
- strconv.Itoa(int) -> convert string ke int
*/
