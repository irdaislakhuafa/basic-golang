package main

import (
	"fmt"
	"reflect"
)

type Example struct {
	Name string
	Age  int
}

func (e Example) WhoAmI() string {
	return e.Name
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi error :D :", r)
		}
	}()

	example1 := Example{Name: "irda islakhu afa", Age: 20}
	PrintExampleStruct(&example1)
}
func PrintExampleStruct(example *Example) {

	exampleType := reflect.TypeOf(*example) // get relfect

	sliceField := make([]interface{}, exampleType.NumField(), exampleType.NumField()) // save field

	for i := 0; i < exampleType.NumField(); i++ {
		sliceField[i] = exampleType.Field(i).Name
	}

	sliceMethod := make([]string, 0, 1) // save method name
	for i := 0; i < exampleType.NumMethod(); i++ {
		sliceMethod = append(sliceMethod, exampleType.Method(i).Name)
	}

	fmt.Println("Example :", exampleType.Name())
	fmt.Println("Total field :", exampleType.NumField())
	fmt.Println("Field :", sliceField)
	fmt.Println("Total method :", exampleType.NumMethod())
	fmt.Println("Methods :", sliceMethod)

}

/*
PACKAGE REFLECT : -> mirip Class.getDeclaredFields() di java
- dalam bahasa pemrograman biasanya ada yang namanya reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan.
- hal ini bisa dilakukan di Go-lang dengan menggunakan package reflect.
- fitur ini mungkin tidak bisa dibahasa secara lengkap dalam satu video, anda bisa eksplorasi package reflect secara otodidak.
- reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan.
- docs : https://golang.org/pkg/reflect
*/
