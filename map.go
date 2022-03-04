package main

import "fmt"

func main() {
	// map disini mirip/sama seperti HashMap di java atau asosiatif array di PHP
	// - cara membuat map
	// variable := map[keyType]valueType
	person := map[string]string{
		// add new data directly to map
		"name":    "irda islakhu afa",
		"address": "tuban",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// beberapa functionn di map
	// - len(map) -> untuk mendapatkan jumlah data di map
	// - map(key) -> mengambil data dari map dengan key
	// - map(key) = value -> merubah value pada key tertentu di map
	// - make(map[keyType]valueType) -> membuat map baru
	// - delete(map, key) -> menghapus data di map dengan key

	// mengambil jumlah data di map
	fmt.Println("\nMap length :", len(person))

	// membuat map tanpa data apapun dengan make(map[key]value)
	book := make(map[string]string)
	fmt.Println("Book :", book)

	book["title"] = "this is book 1"
	book["something"] = "wrong"
	fmt.Println("Book :", book)

	delete(book, "something")
	fmt.Println("Book :", book)

}
