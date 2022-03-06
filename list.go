package main

import (
	"container/list"
	"fmt"
	"strings"
)

func main() {
	// create list
	linkedList := list.New()
	linkedList.PushBack("irda") // ke belakang/akhir
	linkedList.PushFront("ana") // ke depan/awal
	// fmt.Println("Linked List :", linkedList) // yg nampil alamat memorinya

	// print linked list
	// dari depan ke belakang
	for element := linkedList.Front(); element != nil; element = element.Next() {
		fmt.Println("Element :", element.Value)
	}

	fmt.Println(strings.Repeat("-", 20))
	// dari belakang ke depan
	for element := linkedList.Back(); element != nil; element = element.Prev() {
		fmt.Println("Element :", element.Value)
	}
}

/*
PACKAGE CONTAINER/LIST :
- package container/lit adalah implementasi struktur data double linked list di golang
- docs : http://golang.org/pkg/container/list
- ga punya index, data sebelumnya, saat ini, dan selanjutnya saling terhubung
*/
