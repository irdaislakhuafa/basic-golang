package main

import (
	"fmt"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi error :")
		}
	}()

	// type assertion atau konversi tipe data
	randomInterface := random()
	resultString := randomInterface.(string)
	fmt.Println("DATA 1 :", resultString)

	// jangan salah ketika menkonversi "string" ke int atau ke tipe data yg lainnya
	// fmt.Println("Contoh error :", randomInterface.(int)) //terjadierror1

	// menggunakan switch saat menggunakan type assertions
	switch val := random().(type) {
	case string:
		fmt.Printf("variable val tipenya adalah string %s", val)
	case int:
		fmt.Printf("variable val tipenya adalah int %d", val)
	default:
		fmt.Printf("kamu bukan tipeku :D")
	}
}
func random() interface{} {
	return "ini random"
}

/*
TYPE ASSERTIONS :
adalah kemampuan merubah tipe data menjadi tipe data yang di inginkan.
fitur ini biasanya digunakan untuk merubah interface kosong menjadi tip data tertentu, ini mirip Integer.valueOf(object) atau String.valueOf(object) di java

TYPE ASSERTIONS WITH SWITCH :
- jika terjadi error saat menggunakan assertions maka bisa berakibat terjadi panic, jika panic dan tidak ter recover maka program bisa berhenti
- agar lebih aman sebaiknya menggunakan switch untuk melakukan assertions/convertion
*/
