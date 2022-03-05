package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	println("\033\143")
	runApp(false)
	runApp(true)
}
func runApp(err bool) {
	defer endApp() // call function with defer
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	println("\nStarting...")

	if err {
		panic(strings.ToUpper("Terjadi error!"))
	}
	// messageError := reco

	fmt.Println("Aplikasi berjalan:D")
}
func endApp() {
	fmt.Println("Ending start app")
}

/*
PANIC AND RECOVER LIKE TRY AND CATCH
====================================
PANIC FUNCTION :
- adalah function yang bisa digunakan untuk menghentikan program
- biasanya dipanggil ketika terjadi error disaat program berjalan
- saat panic function dipanggil program akan terhenti, tapi panic function tetap akan di eksekusi (secara default jika terjadi eror program akan memanggil panic function untuk berhenti)

RECOVER FUNCTION :
- adalah function yang bisa digunakan untuk menangkap data panic
- dengan recover proses panic akan berhenti jadi function main tidak akan terhendi oleh panic sehingga program tetap akan berjalan
*/
