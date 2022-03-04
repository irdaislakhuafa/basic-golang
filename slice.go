package main

import "fmt"

func main() {
	// slice mereference ke array
	// jika array di buah maka data didalam slice juga akan berubah begitupun sebaliknya
	// slice mirip sub array / sub list kalau dijava
	month := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	slice1 := month[3:7]
	fmt.Println("Slice1 :", slice1)

	// beberapa function di slice
	// - len(slice) -> untuk medapatkan panjang/length slice
	// - cap(slice) -> untuk mendapatkan total kapasitas slice (cap for capacity)
	// - append(slice, data) -> membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh maka akan dibuat slice baru
	// - make([]TypeData, length, capacity) -> untuk membuat slice baru
	// - copy(destination, source) -> mengcopy/menyalin slice ke slice lain
	fmt.Println("Length slice1 :", len(slice1))
	fmt.Println("Capacity slice1 :", cap(slice1))

	fmt.Println("Merubah slice ke 5 :")
	slice1[1] = "(ana ardani)"
	fmt.Println(slice1)
	month[5] = "(ini di ubah)"
	fmt.Println(slice1)

	// create slice2
	slice2 := month[10:]
	fmt.Println("Slice2 :", slice2)
	slice2 = append(slice2, "bubu")
	fmt.Println("Slice2 :", slice2)

	slice3 := append(slice2, "bubur")
	fmt.Println("Slice3 :", slice3)
	fmt.Println("Months :", month)

	// create new slice without create own array
	newSlice := make(
		// data type
		[]string,
		// length
		3,
		// capacity
		5,
	)

	fmt.Println("\ncreate new slice")
	printSlice(newSlice)

	newSlice[0] = "nama_1"
	newSlice[1] = "nama_2"
	newSlice[2] = "nama_3"
	newSlice = append(newSlice, "nama_4")
	newSlice = append(newSlice, "nama_5")
	newSlice = append(newSlice, "nama_6")
	// newSlice[3] = "nama_4" -> index out of bound
	printSlice(newSlice)

	// copy slice
	line(30)
	copySlice := make([]string, len(newSlice), cap(newSlice))
	printSlice(copySlice)

	// copy slice with function slice
	copy(copySlice, newSlice) // kapasitas harus sama
	printSlice(copySlice)

	// array vs slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("Ini Array :", iniArray)
	fmt.Println("Ini Slice :", iniSlice)
}

func line(longValue int) string {
	var line string
	for i := 0; i < int(longValue); i++ {
		line += "="
	}
	return line
}

func printSlice(slice []string) {
	fmt.Println(line(30))
	fmt.Println("Slice :", slice)
	fmt.Println("Length :", len(slice))
	fmt.Println("Capacity :", cap(slice))
}
