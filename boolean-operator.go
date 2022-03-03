package main

func main() {
	ujian := 80
	absensi := 70

	lulusUjian := (ujian >= 80)
	lulusAbsensi := (absensi >= 80)

	println("Lulus ujian :", lulusUjian)
	println("Lulus absensi :", lulusAbsensi)
	println("Lulus :", (lulusUjian && lulusAbsensi))
}
