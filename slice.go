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

}
