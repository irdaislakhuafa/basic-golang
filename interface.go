package main

import (
	"fmt"
	"log"
	"strings"
)

// create interface (same as interface in java)
type ThisIsInterface interface { // sama seperti interface di java, jadi harus ada class implementasinya
	GetName() (string, string)
	WhoAreYou() string
	TestName(variable string) string
}

// cara implementasi interface di Golang tidak seperti java tapi secara otomatis, jika ada struct (atau class di java) yang memiliki struktur sama seperti interface maka itu adalah class/struct implementasinya

// one struct must implement all method in interface

// create struct first to implement
type Person struct {
	FirstName, LastName string
}

// create struct animal
type Animal struct {
	Name string
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(strings.ToUpper("terjadi error!"))
		}
	}()

	// var person ThisIsInterface // error, karena harus struc implementasinya
	person := Person{FirstName: "Irda", LastName: "Islakhu Afa"}
	sayHelloToInterface(person)

	// create variable with type struct Animal
	animal := Animal{Name: "Anjing"}
	animalInfo(animal)
}

func sayHelloToInterface(example ThisIsInterface) {
	fmt.Println("Hello, your name is", example.WhoAreYou())
}

func (p Person) GetName() (string, string) {
	return p.FirstName, p.LastName
}

func (p Person) WhoAreYou() string {
	return (p.FirstName + p.LastName)
}
func (p Person) TestName(text string) string {
	firstName, lastName := p.GetName()
	return firstName + " " + lastName + " <- " + text
}

// animal
func (animal Animal) TestName(text string) string {
	return animal.Name + " <- " + text
}
func (animal Animal) GetName() (string, string) {
	return animal.Name, ""
}
func (animal Animal) WhoAreYou() string {
	return animal.Name
}

func animalInfo(interfaceName ThisIsInterface) {
	fmt.Println("Animal name is : ", interfaceName.TestName("ini namanya"))
}
