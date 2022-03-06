package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("Round 1.7 :", math.Round(1.7))
	fmt.Println("Round 1.3 :", math.Round(1.3))
	fmt.Println(strings.Repeat("-", 20))

	fmt.Println("FLoor 1.7 :", math.Floor(1.7))
	fmt.Println("Ceil 1.3 :", math.Ceil(1.3))
	fmt.Println(strings.Repeat("-", 20))

	fmt.Println("Max :", math.Max(10, 20))
	fmt.Println("Min :", math.Min(10, 20))
}

/*
PACKAGE MATH :
- merupakan package yang berisikan constant dan fungsi2 matematika
- docs : https://golang.org/pkg/math

BEBERAPA LIST METHOD DI MATH:
- math.Round(float64) -> membulatkan float64 ke atas atau kebawah sesuai nilai yang paling dekat
- math.Floor(float64) -> membulatkan float64 kebawah
- math.Ceil(float64) -> membulatkan float64 keatas
- math.Min(float64, float64) -> mengembalikan nilai yang paling kecil
- math.Max(float64, float64) -> mengembalikan nilai yang paling besar
*/
