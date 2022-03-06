package database

import "fmt"

var connection string

// function init akan dieksekusi pertama kali ketika package ini dipanggil
func init() { // private
	fmt.Println("Init dipanggil ")
	connection = "MySQL"
}
func GetConnection() string { // public
	return connection
}

/*
PACKAGE INITIALIZATION :
mirip Constructor di java
- adalah function yang akan dieksekusi pertama kali saat package tersebut di panggil
- ini sangat cocok ketika package berisi function-function untuk berkomunikasi dengan database misal
- untuk membuat package initialization (constructor kalau di java) kita cukup membuat function dengan nama inti()
*/
