package main

import (
	"basic-golang/helpers"
	"fmt"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error :D :", r)
		}
	}()

	now := time.Now()
	fmt.Println("Sekarang :", now) // mendapatkan waktu saat ini
	fmt.Println("tahun ", now.Year())
	fmt.Println("bulan :", now.Month())
	fmt.Println("hari :", now.Day())
	fmt.Println("jam :", now.Hour())
	fmt.Println("menit :", now.Minute())
	fmt.Println("detik :", now.Second())
	fmt.Println("mili detik :", now.Nanosecond())
	helpers.Line(50)

	createdTime := time.Date(2002, time.January, 1, 23, 00, 00, 00, time.UTC) // membuat time/Date sendiri
	fmt.Println("Waktu dibuat :", createdTime)
	helpers.Line(50)

	// parse time layout (mirip SimpleDateFormater di java )
	// time.RFC1123
	// layoutdate := time.RFC822Z
	layoutdate := "2006-01-02"
	parse := "2021-10-01"
	myLayoutDate, err := time.Parse(layoutdate, parse) // parse string to date

	if err != nil {
		fmt.Println("Layout date Error :", err)
	} else {
		fmt.Println("Layout date :", myLayoutDate)
	}
}

/*
PACKAGE TIME :
- Time adalah package yang berisikan fungsionalitas waktu di golang (mirip java.util.Date di java)
- docs : https://golang.org/pkg/time/
*/
