package main

import "fmt"

func main() {
	var value32 int32 = 100_000
	var value64 int64 = int64(value32)
	value16 := int16(value64)
	value8 := int8(value16)

	fmt.Println("value int32 :", value32)
	fmt.Println("value int64 :", value64)
	fmt.Println("value int16 :", value16)
	fmt.Println("value int8 :", value8)

	// try convert byte to string
	name := "irda islakhu afa"
	fmt.Printf(string(name[0]))
}
