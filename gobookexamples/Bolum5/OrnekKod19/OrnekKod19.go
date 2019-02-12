package main

import (
	"fmt"
)

func main() {
	var slice = []int{1, 3, 5, 7, 9, 11, 0, 2, 4, 6, 8}

	var tek_sayilar []int = slice[0:6]
	var cift_sayilar []int = slice[6:]

	fmt.Println("Ana slice = ", slice)
	fmt.Println("Tek Sayılar = ", tek_sayilar)
	fmt.Println("Çift Sayılar = ", cift_sayilar)
}
