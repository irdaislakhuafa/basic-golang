package main

import (
	"errors"
	"fmt"
	"strings"
)

// interface error sama seperti Exception di java
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(strings.ToUpper("terjadi error!"))
		}
	}()

	error := errors.New("sesuatu sedang error!") // create error message/Exception message
	fmt.Println(error.Error())                   // call error message

	fmt.Println(strings.Repeat("-", 15))
	result1, err1 := pembagian(100, 10)

	if err1 == nil {
		fmt.Println("Hasil 1 :", result1)
	} else {
		fmt.Println("Error :", err1.Error())
	}
}

func pembagian(a int, b int) (int, error) { // mirip throws exception

	if b == 0 {
		fmt.Println("anaaa")
		return 0, errors.New(strings.ToUpper("pembagi error! tidak boleh 0"))
	} else {
		result := (a / b)
		return result, nil
	}
}
