package main

import (
	"basic-golang/helpers"
	"fmt"
)

// program ini mereverensi ke package helper
func main() {
	// mengakses yang public
	helpers.SayHelloFromHelper()
	fmt.Println("App name :", helpers.MyApplicationName)

	// mengakses yang public private
	// helpers.sayHelloFromHelper() // not exported ga bisa karena private
}

/*
ACCESS MODIFIER :
- sama seperti acces modifier di java hanya saja cara deklarasinya berbeda tergantung nama vaiavle/function nya
=> Variable = public
=> variable = private
*/
