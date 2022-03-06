package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("i([a-z])a") // membuat regex

	// MatchString mirip perbandingan
	fmt.Println(regex.MatchString("ira"))
	fmt.Println(regex.MatchString("irda"))
	fmt.Println(regex.MatchString("ana"))
	fmt.Println(regex.MatchString("ina"))
	fmt.Println(regex.MatchString("iDa"))

	// find all string if matching with regex
	valid := regex.FindAllString("ira ida anu enek ani nino nana nani iia", -1)
	fmt.Println("List valid regex :", valid)
}

/*
PACKAGE REGEXP :
- utilitas Go lang untuk melakukan pencarian regular expression
- regular expression di go lang menggunakan library C yang dibuat Gooogle bernama RE2
- syntax : https://github.com/google/re2/wiki/Syntax
- docs : https://golang.org/pkg/regexp

other desc : Ekspresi reguler adalah serangkaian karakter yang mendefinisikan sebuah pola pencarian. Pola tersebut biasanya digunakan oleh algoritme pencarian string untuk melakukan operasi "cari" atau "cari dan ganti" pada string, atau untuk memeriksa string masukan.
*/
