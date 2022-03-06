package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	ringData := ring.New(5)

	for i := 0; i < ringData.Len(); i++ {
		ringData.Value = "this is value " + strconv.Itoa(i) // insert value
		ringData = ringData.Next()                          // go to next ring
	}
	fmt.Println("Ring : ", ringData) // yang muncul alamat memori nya

	// iteration ring/jangan pke for loop karena doi gada ujung
	ringData.Do(func(value interface{}) {
		fmt.Println("Data :", value)
	})
}

/*
PACKAGE CONTAINER RING : (cincin)
- adalah implementasi struktur data curcular list (cincin ga ada ujungnya)
- ciscular list adalah struktur data ring dimana akhir elemen akan kembali ke awal (head)
- docs : https://golang.org/pkg/container/ring
*/
