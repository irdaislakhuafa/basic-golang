package main

import (
	"fmt"
	"strings"
)

func main() {
	person1 := PersonData{Name: "irda"}
	fmt.Println("Person 1 :", person1)
	fmt.Println(strings.Repeat("-", 15))

	person1.maried()
	fmt.Println("Person 1 :", person1)
}

type PersonData struct {
	Name string
}

func (person *PersonData) maried() { // dengan referen data akan berubah
	person.Name = "Mr. " + person.Name
}

/*
POINTER IN METHOD :
- walaupun methoad akan menempel di struct tapi sebenarnya method mengambil parameter dengan pass by value
- sangat direkomendasikan menggunakan pointer saat menggunakan method/struct function, sehingga tidak boros memory karena tidak menduplikasi data struct di method nya
*/
